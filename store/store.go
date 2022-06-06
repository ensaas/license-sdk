package store

import (
	"errors"
	"log"
)

var register = map[string]Store{}

type Store interface {
	Save(val interface{}) error
	Load() (interface{}, error)
	Initialize() error
}

func Register(typ string, storeObj Store) error {
	if len(typ) == 0 {
		return errors.New("invalid register data source type")
	}
	if storeObj == nil {
		return errors.New("empty data source")
	}

	if HasDataSource(typ) {
		log.Printf("data source type %s has registered", typ)
		return nil
	}
	register[typ] = storeObj
	return nil
}

func GetStore(typ string) (Store, bool) {
	if d, ok := register[typ]; ok && d != nil {
		return d, ok
	}
	return nil, false
}

func HasDataSource(typ string) bool {
	if s, ok := register[typ]; ok && s != nil {
		return true
	}
	return false
}
