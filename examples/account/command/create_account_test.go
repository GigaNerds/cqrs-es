package command

import (
	"cqrs-es/examples/account"
	"cqrs-es/examples/account/domain"
	"testing"
)

func TestEmitsCreated(t *testing.T) {
	acc := domain.Account{}
	cmd := CreateAccount{
		Owner: "test",
	}
	ev, _ := cmd.ExecuteCommand(&acc)

	if ev.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, ev.Owner)
	}
}

func TestHandle(t *testing.T) {
	svc := account.NewService()
	cmd := CreateAccount{
		Owner: "test",
	}

	agg, err := cmd.HandleWith(svc)
	if err != nil {
		t.Errorf(err.Error())
	}

	if agg.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, agg.Owner)
	}
}
