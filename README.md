# License SDK

SDK for Go language offers a library for validate app license. It realize validate license locally by pass parameters of related and 
if check failed then will call license api which configured.

It's SDK could prevent tamper data from license API,also compatible with old license api version.

## Install

```
 go get -d github.com/ensaas/license-sdk
```

## Usage

```
package main

import (
	"fmt"
	licenseSDK "github.com/ensaas/license-sdk"
	"github.com/ensaas/license-sdk/encryptor"
	"github.com/ensaas/license-sdk/store/postgre"
	"log"
)

func main() {
	salt := "12345678123456781234567812345678"
	entor, err := encryptor.New(salt)
	if err != nil {
		log.Fatalf("init encryptor failed:[%v]", err)
	}

	pgParams := map[string]interface{}{
		postgre.Host:         "localhost",
		postgre.Port:         "5432",
		postgre.Username:     "postgres",
		postgre.Password:     "123456",
		postgre.DBName:       "listing",
		postgre.SSLMode:      "disable",
		postgre.MaxIdleConns: 10,
		postgre.MaxOpenConns: 10,
	}
	pgStore, err := postgre.New(pgParams)
	if err != nil {
		log.Fatalf("init postgre store failed:[%v]", err)
	}
	licenseUrl := "http://localhost:8080/v1"

	// create a license Mgr
	licenseMgr := licenseSDK.New(pgStore, entor, licenseUrl)
	licenseMgr.InitAppParams("LicenseServer", "cluster001-workspaceId21d21d-ensaas", "1dwd12ijijdq")

	checkedLicense := []*licenseSDK.License{
		{
			Pn:              "9806WPDASH",
			Authcode:        "5083-2101-0001",
			Number:          1,
			ExpireTimestamp: 0,
		},
	}
	if err := licenseMgr.StartValidate(checkedLicense); err != nil {
		log.Fatalf("validate license start job failed:[%s]", err.Error())
	}

	availableDays, err := licenseMgr.GetAvailableDays("9806WPDASH")
	if err != nil {
		log.Fatalf("get license available days failed:[%s]", err.Error())
	}
	fmt.Println(availableDays)
}
```

