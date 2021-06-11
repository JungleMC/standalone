package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/storage"
)

var DB *leveldb.DB

func Load() (err error) {
	DB, err = leveldb.Open(storage.NewMemStorage(), nil)
	return
}
