package cmd

import (
	"backend-server/config"
	"fmt"
)

type Cmd struct {
	config *config.Config
}

func NewCmd(filePath string) *Cmd {
	c := &Cmd{
		config: config.NewConfig(filePath),
	}

	fmt.Printf("NewCmd > %s\n", c.config.Server.Port)

	return c
}
