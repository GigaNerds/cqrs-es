package account

import (
	"cqrs-es/examples/account/repository/posgre"
)

type Service struct {
	Repo posgre.Repository
}
