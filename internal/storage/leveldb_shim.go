package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/storage"
	"reflect"

	"github.com/junglemc/JungleTree/internal/configuration"
)

var DB *leveldb.DB

func Load() (err error) {
	if configuration.Config().DebugMode {
		DB, err = leveldb.Open(storage.NewMemStorage(), nil)
	} else {
		DB, err = leveldb.OpenFile("data", nil)
	}
	return
}

func Get(key string, value interface{}, opts *opt.ReadOptions) error {
	v := reflect.ValueOf(value)
	if v.Kind() != reflect.Ptr || v.IsNil() {
		return errors.New("ptr required")
	}

	data, err := DB.Get([]byte(key), opts)
	if err != nil {
		return err
	}

	buf := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buf)

	err = decoder.Decode(value)
	if err != nil {
		return err
	}
	return nil
}

func Put(key string, value interface{}, opts *opt.WriteOptions) error {
	buf := &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return DB.Put([]byte(key), buf.Bytes(), opts)
}

func Has(key string, opts *opt.ReadOptions) (bool, error) {
	return DB.Has([]byte(key), opts)
}
