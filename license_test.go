package licenseSDK

import (
	"github.com/ensaas/license-sdk/encryptor"
	"github.com/ensaas/license-sdk/store/postgre"
	"log"
	"testing"
)

func TestNew(t *testing.T) {
	licenseUrl := "http://api-license-ews.axa.wise-paas.com.cn/v1"

	entor, err := encryptor.New("12345678912345678932145678781231")
	if err != nil {
		log.Fatalf(err.Error())
	}

	storeParams := map[string]interface{}{
		postgre.Host:         "localhost",
		postgre.Port:         "5432",
		postgre.Username:     "postgres",
		postgre.Password:     "",
		postgre.DBName:       "listing",
		postgre.SSLMode:      "disable",
		postgre.MaxIdleConns: 10,
		postgre.MaxOpenConns: 10,
	}
	pgStore, err := postgre.New(storeParams)
	if err != nil {
		log.Fatalf(err.Error())
	}

	New(pgStore, entor, licenseUrl)
}

func TestNewWithDefault(t *testing.T) {
	licenseUrl := "http://api-license-ews.axa.wise-paas.com.cn/v1"
	salt := "12345678912345678932145678781231"
	storeParams := map[string]interface{}{
		postgre.Host:         "localhost",
		postgre.Port:         "5432",
		postgre.Username:     "postgres",
		postgre.Password:     "",
		postgre.DBName:       "listing",
		postgre.SSLMode:      "disable",
		postgre.MaxIdleConns: 10,
		postgre.MaxOpenConns: 10,
	}
	NewWithDefault(storeParams, salt, licenseUrl)
}

func TestNewWithDefaultEncryptor(t *testing.T) {
	licenseUrl := "http://api-license-ews.axa.wise-paas.com.cn/v1"
	salt := "12345678912345678932145678781231"
	storeParams := map[string]interface{}{
		postgre.Host:         "localhost",
		postgre.Port:         "5432",
		postgre.Username:     "postgres",
		postgre.Password:     "",
		postgre.DBName:       "listing",
		postgre.SSLMode:      "disable",
		postgre.MaxIdleConns: 10,
		postgre.MaxOpenConns: 10,
	}
	pgStore, err := postgre.New(storeParams)
	if err != nil {
		log.Fatalf(err.Error())
	}

	NewWithDefaultEncryptor(pgStore, salt, licenseUrl)
}

func TestNewWithDefaultStore(t *testing.T) {
	licenseUrl := "http://api-license-ews.axa.wise-paas.com.cn/v1"
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
	entor, err := encryptor.New("12345678912345678932145678781231")
	if err != nil {
		log.Fatalf(err.Error())
	}

	NewWithDefaultStore(entor, storeParams, licenseUrl)
}
