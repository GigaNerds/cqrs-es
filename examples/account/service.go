package account

import (
	"cqrs-es/examples/account/repository"
)

type Service struct {
	Repo repository.Repository
}
