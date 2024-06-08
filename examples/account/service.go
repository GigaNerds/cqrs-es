package account

import (
	"cqrs-es/examples/account/repository/in_memory"
)

type Service struct {
	Repo in_memory.Repository
}

func NewService() Service {
	repo := in_memory.NewRepository()

	svc := Service{
		Repo: repo,
	}

	return svc
}
