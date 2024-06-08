package posgre

import "github.com/jmoiron/sqlx"

// Repository is a repository for account domain.
type Repository struct {
	// DbConnection is a connection to the database.
	DbConnection sqlx.DB
}
