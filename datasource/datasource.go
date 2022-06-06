package datasource

import (
	"errors"
	_ "github.com/ensaas/license-sdk/datasource/env"
	_ "github.com/ensaas/license-sdk/datasource/rest"
	"log"
)

var register = map[string]DataSource{}

type DataSource interface {
	GetBySrp(service, licenseID string) (SrpLicenseList, error)
	GetByEnSaaS(service, licenseID string) (EnSaaSLicenseList, error)
}

func Register(typ string, dataSource DataSource) error {
	if len(typ) == 0 {
		return errors.New("invalid register data source type")
	}
	if dataSource == nil {
		return errors.New("empty data source")
	}

	if HasDataSource(typ) {
		log.Printf("data source type %s has registered", typ)
		return nil
	}
	register[typ] = dataSource
	return nil
}

func GetDataSource(typ string) (DataSource, bool) {
	if d, ok := register[typ]; ok {
		return d, ok
	}
	return nil, false
}

func HasDataSource(typ string) bool {
	if _, ok := register[typ]; ok {
		return true
	}
	return false
}

func ListDataSource() map[string]DataSource {
	return register
}
