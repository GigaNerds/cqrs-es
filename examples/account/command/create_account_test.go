package command

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"cqrs-es/examples/account/event"
	"testing"
)

func TestEmitsCreated(t *testing.T) {
	acc := domain.Account{}
	cmd := CreateAccount{
		Owner: "test",
	}
	ev, _ := cmd.ExecuteCommand(&acc)
	created, ok := ev.(*event.AccountCreated)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	if created.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, created.Owner)
	}
}

func TestHandleCreateAccount(t *testing.T) {
	svc := account.NewService()
	cmd := CreateAccount{
		Owner: "test",
	}

	agg, ev, err := svc.HandleCommand(cmd)
	if err != nil {
		t.Errorf(err.Error())
	}

	created, ok := ev.(*event.AccountCreated)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	if created.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, created.Owner)
	}

	if agg.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, agg.Owner)
	}
}
