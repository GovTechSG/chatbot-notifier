package yamlcustom

import (
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// ConfigAWS struct for yaml AWS config
type ConfigAWS struct {
	Arn string `yaml:"arn"`
}

// ConfigTelegram struct for yaml telegram config
type ConfigTelegram struct {
	Token  string `yaml:"token"`
	ChatID int64  `yaml:"chatid"`
}

//EncryptConfigTelegram encrypt ConfigTelegram struct value
type EncryptConfigTelegram struct {
	Token  string `yaml:"token"`
	ChatID string `yaml:"chatid"`
}

// Config struct for overall yaml config
type Config struct {
	Aws      []ConfigAWS      `yaml:"aws"`
	Telegram []ConfigTelegram `yaml:"telegram"`
}

// EncryptConfig struct for saving encrypted yaml
type EncryptConfig struct {
	Aws      []ConfigAWS             `yaml:"aws"`
	Telegram []EncryptConfigTelegram `yaml:"telegram"`
}

// ParseYAML parse yaml config file
func ParseYAML(fileName string) Config {
	filename, _ := filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config Config

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

// ParseEncyptYAML parse ecnrypted yaml config file
func ParseEncyptYAML(fileName string) EncryptConfig {
	filename, _ := filepath.Abs(fileName)
	yamlFile, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	var config EncryptConfig

	err = yaml.Unmarshal(yamlFile, &config)
	if err != nil {
		panic(err)
	}

	return config
}

// EditYAML Modify value to config to encryption
func EditYAML(a string, t string, c string) EncryptConfig {
	ec := EncryptConfig{
		Aws: []ConfigAWS{
			ConfigAWS{
				Arn: a,
			},
		},
		Telegram: []EncryptConfigTelegram{
			EncryptConfigTelegram{
				Token:  t,
				ChatID: c,
			},
		},
	}
	return ec
}