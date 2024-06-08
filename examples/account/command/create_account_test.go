package command

import (
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
