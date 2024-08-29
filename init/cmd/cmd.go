package cmd

import (
	"backend-server/config"
	"backend-server/network"
	"backend-server/repository"
	"backend-server/service"
	"fmt"
)

type Cmd struct {
	config     *config.Config
	network    *network.Network
	repository *repository.Repository
	service    *service.Service
}

func NewCmd(filePath string) *Cmd {

	c := &Cmd{
		config: config.NewConfig(filePath),
		// network: network.NewNetwork(),
	}

	c.repository = repository.NewRepository()
	c.service = service.NewServcie(c.repository)
	c.network = network.NewNetwork(c.service)

	// 이 아래는 동작안하는 부분을 ServerStart() 추가로 변경
	fmt.Printf("NewCmd > %s\n", c.config.Server.Port)

	c.network.ServerStart(c.config.Server.Port)

	return c
}
