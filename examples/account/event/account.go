package event

import "cqrs-es/examples/account/domain"

type AccountCreated struct {
	AccountId domain.Id
	Owner     domain.Owner
	At        domain.CreationTime
}

func (ev *AccountCreated) ApplyTo(agg *domain.Account) {
	agg.Id = ev.AccountId
	agg.Owner = ev.Owner
	agg.Balance = 0
	agg.CreatedAt = ev.At
	agg.DeletedAt = ""
}

type DepositTime string

type AccountDeposit struct {
	AccountId domain.Id
	Amount    domain.Balance
	At        DepositTime
}

func (ev *AccountDeposit) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.Balance += ev.Amount
}

type WithdrawalTime string

type AccountWithdrawal struct {
	AccountId domain.Id
	Amount    domain.Balance
	At        WithdrawalTime
}

func (ev *AccountWithdrawal) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.Balance -= ev.Amount
}

type AccountDeleted struct {
	AccountId domain.Id
	At        domain.DeletionTime
}

func (ev *AccountDeleted) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.DeletedAt = ev.At
}
