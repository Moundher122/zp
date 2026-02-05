package config

import (
	"github.com/dgraph-io/badger/v4"
)

type DbConfig struct {
	Db badger.DB
}

func NewDbConfig(db_path string) *DbConfig {
	opts := badger.DefaultOptions(db_path)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}

	return &DbConfig{
		Db: *db,
	}
}
