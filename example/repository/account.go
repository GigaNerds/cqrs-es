package repository

import (
	"cqrs-es/example"
)

func (r *Repository) SaveAggregate(agg example.Account) error {
	statement := `INSERT INTO accounts (id, owner, balance, created_at, deleted_at)
    			  VALUES ($1, $2, $3, $4, $5)
    		      ON CONFLICT(id)
    			  DO UPDATE SET
        			  owner = EXCLUDED.owner,
        			  balance = EXCLUDED.balance,
        			  created_at = EXCLUDED.created_at,
        			  deleted_at = EXCLUDED.deleted_at;`

	// TODO: Move this transaction start to previous layer. Also we need to divide transacted/non-transacted calls.
	conn, err := r.DbConnection.Beginx()
	if err != nil {
		return err
	}

	_, err = conn.Exec(statement,
		agg.Id,
		agg.Owner,
		agg.Balance,
		agg.CreatedAt,
		agg.DeletedAt)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) LoadAggregate(id example.Id) (example.Account, error) {
	statement := "SELECT id, owner, balance, created_at, deleted_at FROM accounts WHERE id = $1"

	var resId example.Id
	var owner example.Owner
	var balance example.Balance
	var createdAt example.CreationTime
	var deletedAt example.DeletionTime

	conn, err := r.DbConnection.Beginx()
	if err != nil {
		return example.Account{}, err
	}

	err = conn.QueryRow(statement, id).Scan(&resId, &owner, &balance, &createdAt, &deletedAt)
	if err != nil {
		return example.Account{}, err
	}

	account := example.Account{
		Id:        resId,
		Owner:     owner,
		Balance:   balance,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}

	return account, nil
}
