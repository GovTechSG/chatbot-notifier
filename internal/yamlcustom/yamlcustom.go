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
		// log.Fatal("Yaml already encrypted")
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
		// log.Fatal("Yaml already encrypted")
		panic(err)
	}

	return config
}

// EditYAML Modify value to config to encryption
func EditYAML(a string, t string, c string) EncryptConfig {
	tff := EncryptConfig{
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
	// fmt.Printf("%+v", tff)
	return tff
}

// // EditYAML Modify value to config to encryption
// func EditYAML(a string) Config {
// 	tff := Config{
// 		Aws: []ConfigAWS{
// 			ConfigAWS{
// 				Arn: a,
// 			},
// 		},
// 		Telegram: []ConfigTelegram{
// 			ConfigTelegram{
// 				Token:  "1058949406:AAEdGEVVitmvD9KWJhp1Iz1Z7rGrO3c9fR4",
// 				ChatID: -219357966,
// 			},
// 		},
// 	}

// 	log.SetFlags(log.Lshortfile)

// 	// data, err := ioutil.ReadFile("regexes.yaml")
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// var v interface{}

// 	// d, err := yaml.Marshal(&tff)
// 	// if err != nil {
// 	// 	log.Fatalf("error: %v", err)
// 	// }

// 	// err = yaml.Unmarshal(d, &v)
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }

// 	// // m := make(map[interface{}]interface{})

// 	// // err = yaml.Unmarshal([]byte(d), &m)
// 	// // if err != nil {
// 	// // 	log.Fatalf("error: %v", err)
// 	// // }
// 	// // // fmt.Printf("--- m:\n%v\n\n", m)

// 	// // d, err = yaml.Marshal(&m)
// 	// // if err != nil {
// 	// // 	log.Fatalf("error: %v", err)
// 	// // }
// 	// // fmt.Printf("--- m dump:\n%s\n\n", string(d))

// 	// f, err := os.Create("regexes.yaml.go")
// 	// if err != nil {
// 	// 	log.Fatalln(err)
// 	// }
// 	// defer func() {
// 	// 	err := f.Close()
// 	// 	if err != nil {
// 	// 		log.Fatalln(err)
// 	// 	}
// 	// }()

// 	// fmt.Fprintf(f, "package main\n\n")
// 	// fmt.Fprintf(f, "var regexes = %#v\n", v)
// 	// fmt.Printf("%+v", tff)
// 	return tff

// }
