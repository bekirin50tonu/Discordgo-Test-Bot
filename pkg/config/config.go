package config

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Yml *YmlConfig
}

type YmlConfig struct {
	Name    string   `yml:"name"`
	Version float32  `yml:"version"`
	Token   string   `yml:"token"`
	Owners  []string `yml:"owners"`
}

func New(path string) *Config {
	return &Config{
		Yml: readYmlConfig(path),
	}
}

func readYmlConfig[T YmlConfig](filePath string) *T {

	base_dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path := filepath.Join(base_dir, filePath)
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		panic(err)
	}
	data := T{}
	err = yaml.Unmarshal(buf, &data)
	if err != nil {
		panic(err)
	}
	return &data

}
