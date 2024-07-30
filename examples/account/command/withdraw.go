package command

import (
	"errors"
	"time"

	"github.com/GigaNerds/cqrs_es/examples/account"
	"github.com/GigaNerds/cqrs_es/examples/account/domain"
	"github.com/GigaNerds/cqrs_es/examples/account/event"
)

type Withdraw struct {
	AccountId domain.AccountId
	Amount    domain.AccountBalance
}

// ExecuteCommand describes logic of applying this command to the examples.Account object.
func (c Withdraw) ExecuteCommand(a *domain.Account) (account.Event, error) {
	if a.Balance < c.Amount {
		return nil, errors.New("not enough money")
	}

	withdrawal := event.AccountWithdrawal{
		AccountId: c.AccountId,
		Amount:    c.Amount,
		At:        event.WithdrawalTime(time.Now().UTC()),
	}
	return &withdrawal, nil
}

func (c Withdraw) GetRelatedId() (domain.AccountId, error) {
	return c.AccountId, nil
}
