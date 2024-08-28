package config

import (
	"os"

	"github.com/naoina/toml"
)

// 설정 구조체
type Config struct {
	Server struct {
		Port string
	}
}

func NewConfig(filePath string) *Config {
	// var c *Config
	c := new(Config)

	if file, err := os.Open(filePath); err != nil {
		panic(err)
	} else if err = toml.NewDecoder(file).Decode(c); err != nil {
		panic(err)
	} else {
		return c
	}

}
