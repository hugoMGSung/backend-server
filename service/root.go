package service

import (
	"backend-server/repository"
	"sync"
)

var (
	serviceInit     sync.Once
	serviceInstance *Service
)

// Network, Repository 다리역할
type Service struct {
	// repository
	repository *repository.Repository

	User *User
}

func NewServcie(rep *repository.Repository) *Service {
	serviceInit.Do(func() {
		serviceInstance = &Service{
			repository: rep,
		}

		serviceInstance.User = newUserService(rep.User)
	})

	return serviceInstance
}
