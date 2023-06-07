package db

import "database/sql"

type SQLStore struct {
	*Queries
	db *sql.DB
}

type Store interface {
	Querier
}

func NewStore(db *sql.DB) Store {
	return SQLStore{
		db:      db,
		Queries: New(db),
	}
}
