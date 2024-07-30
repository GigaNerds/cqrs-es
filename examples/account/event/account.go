package event

import (
	"time"

	"github.com/GigaNerds/cqrs_es/examples/account/domain"
)

type AccountActivated struct {
	AccountId domain.AccountId
	At        domain.ActivationTime
}

func (ev *AccountActivated) ApplyTo(agg *domain.Account) {
	agg.ActivatedAt = ev.At
}

func (ev *AccountActivated) GetRelatedId() domain.AccountId {
	return ev.AccountId
}

func (ev *AccountActivated) GetEventType() string {
	return "account.activated"
}

func (ev *AccountActivated) GetHappenedAt() time.Time {
	return time.Time(ev.At)
}

type AccountCreated struct {
	AccountId domain.AccountId
	Owner     domain.AccountOwner
	At        domain.CreationTime
}

func (ev *AccountCreated) ApplyTo(agg *domain.Account) {
	agg.Id = ev.AccountId
	agg.Owner = ev.Owner
	agg.Balance = 0
	agg.ActivatedAt = domain.ActivationTime{}
	agg.CreatedAt = ev.At
	agg.DeletedAt = domain.DeactivationTime{}
}

func (ev *AccountCreated) GetRelatedId() domain.AccountId {
	return ev.AccountId
}

func (ev *AccountCreated) GetEventType() string {
	return "account.created"
}

func (ev *AccountCreated) GetHappenedAt() time.Time {
	return time.Time(ev.At)
}

type DepositTime time.Time

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

func (ev *AccountDeposit) GetEventType() string {
	return "account.deposit"
}

func (ev *AccountDeposit) GetHappenedAt() time.Time {
	return time.Time(ev.At)
}

type WithdrawalTime time.Time

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

func (ev *AccountWithdrawal) GetEventType() string {
	return "account.withdrawal"
}

func (ev *AccountWithdrawal) GetHappenedAt() time.Time {
	return time.Time(ev.At)
}

type AccountDeactivated struct {
	AccountId domain.AccountId
	At        domain.DeactivationTime
}

func (ev *AccountDeactivated) ApplyTo(agg *domain.Account) {
	if ev.AccountId != agg.Id {
		panic("`AccountId` mismatch")
	}
	agg.DeletedAt = ev.At
}

func (ev *AccountDeactivated) GetRelatedId() domain.AccountId {
	return ev.AccountId
}

func (ev *AccountDeactivated) GetEventType() string {
	return "account.deactivated"
}

func (ev *AccountDeactivated) GetHappenAt() time.Time {
	return time.Time(ev.At)
}
