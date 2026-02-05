package config

import (
	"github.com/dgraph-io/badger/v4"
)

func NewDbConfig(db_path string) *badger.DB {
	opts := badger.DefaultOptions(db_path)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}

	return db
}
