package env

import "github.com/ensaas/license-sdk/datasource"

const DataSourceType = "environ"

func init() {
	if err := datasource.Register(DataSourceType, new(environ)); err != nil {
		panic("license register env data source failed")
	}
}

type environ struct{}

func (r *environ) GetBySrp(service, licenseID string) (datasource.SrpLicenseList, error) {

}

func (r *environ) GetByEnSaaS(service, licenseID string) (datasource.EnSaaSLicenseList, error) {

}
