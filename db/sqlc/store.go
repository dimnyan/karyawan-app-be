package db

import "database/sql"

type Store interface {
	Querier
}

// SQLStore provides all functions to execute sql queries and transactions
type SQLStore struct {
	*Queries         // composition, instead of inheritance
	db       *sql.DB // database
}

// NewStore creates a new Store
func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
