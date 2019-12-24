package option

import (
	"chatbot-notifier/internal/awskms"
	"chatbot-notifier/internal/telegram"
	"chatbot-notifier/internal/yamlcustom"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"

	"gopkg.in/yaml.v2"
)

// Command struct for sub command name
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
	gc.fs.StringVar(&gc.subName, "m", "", "file message that need to be send")

	return gc
}

// TextCommand Contain flag command for sending text
func TextCommand() *Command {
	gc := &Command{
		fs: flag.NewFlagSet("text", flag.ContinueOnError),
	}

	gc.fs.StringVar(&gc.name, "f", "", "File to be decrypt and get credential")
	gc.fs.StringVar(&gc.subName, "m", "", "text message that need to be send")

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
	} else if g.Name() == "send" {
		t := g.ReadFile()
		Decrypt(g, t)
	} else if g.Name() == "text" {
		Decrypt(g, g.subName)
	}
}

// SendMessage send message to telegram chat
func SendMessage(t string, c int64, m string) {
	telegram.SendMsg(t, c, m)
}

// Decrypt will trigger decryption of yaml file
func Decrypt(g *Command, tx string) {
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

	SendMessage(tk, sci, tx)
}

//ReadFile to read content of file message
func (g *Command) ReadFile() string {
	// Read message file content
	content, err := ioutil.ReadFile(g.subName)
	if err != nil {
		log.Fatal(err)
	}

	return string(content)
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

// Root Manage command line
func Root(args []string) error {
	if len(args) < 1 {
		h := help()
		fmt.Println(h)
		return errors.New("You must pass a sub-command")
	}

	cmds := []Runner{
		EncryptCommand(),
		SendCommand(),
		TextCommand(),
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
