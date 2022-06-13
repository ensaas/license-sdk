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
func main(){
    licenseUrl := "http://api-license-ews.axa.wise-paas.com.cn/v1"
    
    encyptor, err := encryptor.New("12345678912345678932145678781231")
    if err != nil {
        log.Fatalf(err.Error())
    }

    storeParams := map[string]interface{}{
        postgre.Host:         "localhost",
        postgre.Port:         "5432",
        postgre.Username:     "postgres",
        postgre.Password:     "123456",
        postgre.DBName:       "listing",
        postgre.SSLMode:      "disable",
        postgre.MaxIdleConns: 10,
        postgre.MaxOpenConns: 10,
    }
    pgStore, err := postgre.New(storeParams)
    if err != nil {
        log.Fatalf(err.Error())
    }

    New(pgStore, encyptor, licenseUrl)
}
```

