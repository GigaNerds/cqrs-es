package event

import "cqrs-es/examples/account/domain"

type AccountCreated struct {
	AccountId domain.AccountId
	Owner     domain.AccountOwner
	At        domain.CreationTime
}

func (ev *AccountCreated) ApplyTo(agg *domain.Account) {
	agg.Id = ev.AccountId
	agg.Owner = ev.Owner
	agg.Balance = 0
	agg.CreatedAt = ev.At
	agg.DeletedAt = ""
}

func (ev *AccountCreated) GetRelatedId() domain.AccountId {
	return ev.AccountId
}

type DepositTime string

type AccountDeposit struct {
	AccountId domain.AccountId
	Amount    domain.AccountBalance
	At        DepositTime
}

func (ev *AccountDeposit) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.Balance += ev.Amount
}

func (ev *AccountDeposit) GetRelatedId() domain.AccountId {
	return ev.AccountId
}

type WithdrawalTime string

type AccountWithdrawal struct {
	AccountId domain.AccountId
	Amount    domain.AccountBalance
	At        WithdrawalTime
}

func (ev *AccountWithdrawal) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.Balance -= ev.Amount
}

func (ev *AccountWithdrawal) GetRelatedId() domain.AccountId {
	return ev.AccountId
}

type AccountDeleted struct {
	AccountId domain.AccountId
	At        domain.DeletionTime
}

func (ev *AccountDeleted) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.DeletedAt = ev.At
}

func (ev *AccountDeleted) GetRelatedId() domain.AccountId {
	return ev.AccountId
}
