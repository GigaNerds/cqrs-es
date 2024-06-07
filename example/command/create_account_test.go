package command

import (
	"cqrs-es/example"
	"testing"
)

func TestEmitsCreated(t *testing.T) {
	acc := example.Account{}
	cmd := CreateAccount{
		Owner: "test",
	}
	ev, _ := cmd.ExecuteCommand(&acc)

	if ev.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, ev.Owner)
	}
}