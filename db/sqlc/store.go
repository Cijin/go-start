package db

import "database/sql"

type Store interface {
	Querier
}

// Store provides methods to execute SQL db queries and transactions
type SQLStore struct {
	*Queries
	db *sql.DB
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
