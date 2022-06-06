package rest

import (
	"github.com/ensaas/license-sdk/datasource"
	"log"
	"os"
)

const DataSourceType = "rest"

func init() {
	if err := datasource.Register(DataSourceType, newRest()); err != nil {
		log.Panicf("register rest data source failed:%s", err.Error())
	}
}

type rest struct {
	LicenseURL string
}

func (r *rest) GetBySrp(service, licenseID string) (datasource.SrpLicenseList, error) {

}

func (r *rest) GetByEnSaaS(service, licenseID string) (datasource.EnSaaSLicenseList, error) {

}

func newRest() *rest {
	licenseURL, exist := os.LookupEnv("LICENSE_URL")
	if !exist {
		return nil
	}
	return &rest{LicenseURL: licenseURL}
}
