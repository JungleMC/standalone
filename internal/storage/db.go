package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"github.com/syndtr/goleveldb/leveldb/storage"
	"reflect"

	. "github.com/junglemc/JungleTree/pkg/util"
)

var DB *leveldb.DB

func Load() (err error) {
	DB, err = leveldb.Open(storage.NewMemStorage(), nil)
	return
}

func Get(key Identifier, value interface{}, opts *opt.ReadOptions) error {
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

	val := v.Elem().Interface()
	err = decoder.Decode(val)
	if err != nil {
		return err
	}
	v.Elem().Set(reflect.ValueOf(val))
	return nil
}

func Put(key Identifier, value interface{}, opts *opt.WriteOptions) error {
	buf := &bytes.Buffer{}
	encoder := gob.NewEncoder(buf)
	if err := encoder.Encode(value); err != nil {
		return err
	}
	return DB.Put([]byte(key), buf.Bytes(), opts)
}

func Has(key Identifier, opts *opt.ReadOptions) (bool, error) {
	return DB.Has([]byte(key), opts)
}
