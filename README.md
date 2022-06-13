# License SDK

SDK for Go language offers a library for validate app license. It realize validate license locally by pass parameters of related and 
if check failed then will call license api which configured.

It's SDK could prevent tamper data from license API,also compatible with old license api version.

## Install

```
 go get -d github.com/ensaas/license-sdk
```

## Usage

```go
package main

import (
	"fmt"
	licenseSDK "github.com/ensaas/license-sdk"
	"github.com/ensaas/license-sdk/encryptor"
	"github.com/ensaas/license-sdk/store/postgre"
	"log"
)

func main() {
	// salt must has 32 byte
	salt := "12345678123456781234567812345678"
	// create a encryptor for encrypt data is store
	entor, err := encryptor.New(salt)
	if err != nil {
		log.Fatalf("init encryptor failed:[%v]", err)
	}

	// postgres param
	pgParams := map[string]interface{}{
		postgre.Host:         "172.21.84.188",
		postgre.Port:         "5432",
		postgre.Username:     "postgres",
		postgre.Password:     "123456",
		postgre.DBName:       "listing",
		postgre.SSLMode:      "disable",
		postgre.MaxIdleConns: 10,
		postgre.MaxOpenConns: 10,
	}
	// init a postgre store
	pgStore, err := postgre.New(pgParams)
	if err != nil {
		log.Fatalf("init postgre store failed:[%v]", err)
	}
	licenseUrl := "http://localhost:8080/v1"

	// create a license Mgr
	licenseMgr := licenseSDK.New(pgStore, entor, licenseUrl)
	// init license app related params
	licenseMgr.InitAppParams("hh", "dddaaacacsa", "")

	// license which will be checked
	checkedLicense := []*licenseSDK.License{
		{
			Pn:              "dd111ddd",
			Authcode:        "fc51-5c23-0002",
			Number:          2,
			ExpireTimestamp: 0,
		},
	}
	// start validate license job
	if err := licenseMgr.StartValidate(checkedLicense); err != nil {
		log.Fatalf("validate license start job failed:[%s]", err.Error())
	}
	// get license available days by pn
	availableDays, err := licenseMgr.GetAvailableDays("9806WPDASH")
	if err != nil {
		log.Fatalf("get license available days failed:[%s]", err.Error())
	}
	fmt.Println(availableDays)
	// check license is legal or not
	isValid,err := licenseMgr.IsLegalLicense("9806WPDASH")
	if err != nil{
		log.Fatalf("is license legal failed:[%v]",err)
	}
	fmt.Println(isValid)
}
```

## Note

LicenseSDK need binary program Liv for linux or Liv.exe for windows,so need put this binary program in /usr/local/bin for linux or
set path for windows

There is another way to make this binary enabled that is make dir "bin" under app program root directory,then put "Liv" or "Liv.exe"
in "bin" directory

