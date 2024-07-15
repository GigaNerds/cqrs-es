package command

import (
	"testing"

	"github.com/GigaNerds/cqrs_es/examples/account"
	"github.com/GigaNerds/cqrs_es/examples/account/domain"
	"github.com/GigaNerds/cqrs_es/examples/account/event"
)

func TestEmitsCreated(t *testing.T) {
	acc := domain.Account{}
	cmd := CreateAccount{
		Owner: "test",
	}
	ev, _ := cmd.ExecuteCommand(&acc)
	evs, ok := ev.(*account.EventSet)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	created := evs.Events[0].(*event.AccountCreated)
	if created.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, created.Owner)
	}

	activated := evs.Events[1].(*event.AccountActivated)
	if activated.AccountId != created.AccountId {
		t.Errorf("Expected id to be '%d', got '%d'", created.AccountId, activated.AccountId)
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

	evs, ok := ev.(*account.EventSet)
	if !ok {
		t.Errorf("Wrong event type returned")
	}

	created := evs.Events[0].(*event.AccountCreated)
	if created.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, created.Owner)
	}
	activated := evs.Events[1].(*event.AccountActivated)
	if activated.AccountId != created.AccountId {
		t.Errorf("Expected id to be '%d', got '%d'", created.AccountId, activated.AccountId)
	}

	if agg.Owner != cmd.Owner {
		t.Errorf("Expected owner to be '%s', got '%s'", cmd.Owner, agg.Owner)
	}
}
