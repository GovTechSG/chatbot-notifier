package main

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"chatbot-notifier/internal/awskms"
	"chatbot-notifier/internal/telegram"
	"chatbot-notifier/internal/yamlcustom"
	"gopkg.in/yaml.v2"
)

// Command struct for flag
type Command struct {
	fs *flag.FlagSet

	name    string
	subName string
}

// EncryptCommand Contain flag command for encrypt
func EncryptCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("encrypt", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "f", "", "File to be encrypt")

	return gc
}

// SendCommand Contain flag command for decrypt
func SendCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("send", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "f", "", "File to be decrypt and get credential")
	gc.fs.StringVar(&gc.subName, "m", "hello", "file message that need to be send")

	return gc
}

// Name to return subcommand name
func (g *Command) Name() string {
	return g.fs.Name()
}

// Init get argument after subcommand
func (g *Command) Init(args []string) error {
	return g.fs.Parse(args)
}

// Crypto to decide encryption or decryption
func (g *Command) Crypto() {

	if g.Name() == "encrypt" {
		Encrypt(g)
	}

	if g.Name() == "send" {
		Decrypt(g)
	}
}

// SendMessage send message to telegram chat
func SendMessage(t string, c int64, m string) {
	telegram.SendMsg(t, c, m)
}

// Decrypt will trigger decryption of yaml file
func Decrypt(g *Command) {
	output := yamlcustom.ParseEncyptYAML(g.name)
	t := output.Telegram[0].Token
	c := output.Telegram[0].ChatID
	a := output.Aws[0].Arn

	tk := awskms.DecryptAwsCred(a, t)
	ci := awskms.DecryptAwsCred(a, c)

	sci, err := strconv.ParseInt(ci, 10, 64)
	if err != nil {
		panic(err)
	}

	content, err := ioutil.ReadFile(g.subName)
	if err != nil {
		log.Fatal(err)
	}

	// Convert []byte to string and print to screen
	text := string(content)

	SendMessage(tk, sci, text)
}

// Encrypt will trigger encryption of yaml file
func Encrypt(g *Command) {
	output := yamlcustom.ParseYAML(g.name)
	t := output.Telegram[0].Token
	c := output.Telegram[0].ChatID
	a := output.Aws[0].Arn
	s := strconv.FormatInt(c, 10)
	encToken := awskms.ReadAwsCred(a, t)
	encChatID := awskms.ReadAwsCred(a, s)

	convertYAML := yamlcustom.EditYAML(a, encToken, encChatID)

	d, err := yaml.Marshal(&convertYAML)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	err = ioutil.WriteFile(g.name, d, 0644)
	if err != nil {
		log.Fatal(err)
	}
}

// Runner to run function with same method
type Runner interface {
	Init([]string) error
	Name() string
	Crypto()
}

func help() string {
	helpMessage := "Usage: \nnotifier send -f < credential.yml credential file> -m < message.txt text message file > >\nnotifier encrypt -f <credential.yml credential to encrypt>\n"
	return helpMessage
}

func root(args []string) error {
	if len(args) < 1 {
		h := help()
		fmt.Println(h)
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		EncryptCommand(),
		SendCommand(),
	}

	subcommand := os.Args[1]

	for _, cmd := range cmds {
		if cmd.Name() == subcommand {
			cmd.Init(os.Args[2:])
			cmd.Crypto()
			return nil
		}
	}

	return fmt.Errorf("Unknown subcommand: %s", subcommand)
}

func main() {
	if err := root(os.Args[1:]); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
