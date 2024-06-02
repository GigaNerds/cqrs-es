package example

type AccountCreated struct {
	Owner string
	At    CreationTime
}

func (ev *AccountCreated) ApplyTo(agg *Account) {
	agg.Id = NewId()
	agg.Owner = ev.Owner
	agg.Balance = 0
	agg.CreatedAt = ev.At
	agg.DeletedAt = ""
}

type DepositTime string

type AccountDeposit struct {
	AccountId Id
	Amount    int
	At        DepositTime
}

func (ev *AccountDeposit) ApplyTo(agg *Account) {
	if ev.AccountId != agg.Id {
		panic("`Id` mismatch")
	}
	agg.Balance += ev.Amount
}

type WithdrawalTime string

type AccountWithdrawal struct {
	AccountId Id
	Amount    int
	At        WithdrawalTime
}

func (ev *AccountWithdrawal) ApplyTo(agg *Account) {
	if ev.AccountId != agg.Id {
		panic("`Id` mismatch")
	}
	agg.Balance -= ev.Amount
}

type AccountDeleted struct {
	AccountId Id
	At        DeletionTime
}

func (ev *AccountDeleted) ApplyTo(agg *Account) {
	if ev.AccountId != agg.Id {
		panic("`Id` mismatch")
	}
	agg.DeletedAt = ev.At
}
