package command

import (
	"time"

	"github.com/GigaNerds/cqrs_es/examples/account"
	"github.com/GigaNerds/cqrs_es/examples/account/domain"
	"github.com/GigaNerds/cqrs_es/examples/account/event"
)

type Deposit struct {
	AccountId domain.AccountId
	Amount    domain.AccountBalance
}

// ExecuteCommand describes logic of applying this command to the examples.Account object.
func (c Deposit) ExecuteCommand(_ *domain.Account) (account.Event, error) {
	deposit := event.AccountDeposit{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        event.DepositTime(time.Now().UTC().String()),
	}
	return &deposit, nil
}

func (c Deposit) GetRelatedId() (domain.AccountId, error) {
	return c.AccountId, nil
}
