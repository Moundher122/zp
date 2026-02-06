package config

import (
	"github.com/dgraph-io/badger/v4"
)

type DB interface {
	Update(func(*badger.Txn) error) error
	View(func(*badger.Txn) error) error
}

func NewDbConfig(db_path string) *badger.DB {
	opts := badger.DefaultOptions(db_path)
	db, err := badger.Open(opts)
	if err != nil {
		panic(err)
	}
	return db
}

func AddToDb(db DB, key, value []byte) error {
	return db.Update(func(txn *badger.Txn) error {
		err := txn.Set(key, value)
		return err
	})
}
func GetFromDb(db DB, key []byte) ([]byte, error) {
	var value []byte
	err := db.View(func(txn *badger.Txn) error {
		item, err := txn.Get(key)
		if err != nil {
			return err
		}
		return item.Value(func(val []byte) error {
			value = append([]byte{}, val...)
			return nil

		})
	})
	return value, err
}
