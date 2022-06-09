package main

import (
	"fmt"
	"github.com/ensaas/license-sdk/checker"
	"github.com/ensaas/license-sdk/encryptor"
	"github.com/ensaas/license-sdk/liv"
	"github.com/ensaas/license-sdk/models"
	"github.com/ensaas/license-sdk/retrieve"
	"github.com/ensaas/license-sdk/store"
	"github.com/ensaas/license-sdk/store/postgre"
	"github.com/robfig/cron"
	"github.com/sirupsen/logrus"
	"log"
)

const (
	version          = "1.0.1"
	checkLicenseCron = "0 0 0,8,16 * * ?"
)

var appParam *appParams

// License is app license used be validate
type License struct {
	Pn              string
	Authcode        string
	ExpireTimestamp int64
	Number          int
}

// AppParams is parameter of app and which are fixed
type appParams struct {
	ServiceName string
	LicenseId   string
	ActiveInfo  string
}

// LicenseManager is for manager app license lifecycle
type LicenseManager interface {
	// GetStatus get current license check status by pn
	GetAvailableDays(pn string) (int, error)
	//InitAppParams initialize app parameters which are fixed value
	InitAppParams(serviceName, licenseId, activeInfo string)
	// StartValidate is validate license status
	StartValidate(licenses []*License) error
	// Version print sdk version
	Version() string
	// LivVersion
	LivVersion() (string, error)
}

type licenseManager struct {
	checker   checker.Checker
	retriever retrieve.Retriever
	licenses  []*License
}

func (lic *licenseManager) InitAppParams(serviceName, licenseId, activeInfo string) {
	appParam = &appParams{
		ServiceName: serviceName,
		LicenseId:   licenseId,
		ActiveInfo:  activeInfo,
	}
}

// GetStatus is get check license status,data in db format is serviceName:failedTimes
func (lic *licenseManager) GetAvailableDays(pn string) (int, error) {
	return lic.checker.GetAvailableDays(pn)
}

// StartValidate start validate license jobs
func (lic *licenseManager) StartValidate(licenses []*License) (err error) {
	lic.licenses = licenses

	// when app start check license
	if err = lic.licenseCheck(licenses, true); err != nil {
		return
	}

	// check license in fixed time
	cronObj := cron.New()
	if err := cronObj.AddJob(checkLicenseCron, lic); err != nil {
		log.Printf("license start check license job failed:%s", err.Error())
		return err
	}
	cronObj.Start()

	return
}

// CheckLicenseJob
func (lic *licenseManager) Run() {
	if err := lic.licenseCheck(lic.licenses, false); err != nil {
		logrus.Errorf("[license check]- check license failed:[%v]", err)
	}
}

// licenseCheck is check license
func (lic *licenseManager) licenseCheck(licenses []*License, isAppBootCheck bool) error {
	// if all license from app check failed
	isAllFailed := true
	for _, license := range licenses {
		commonLic := &models.License{
			LicenseID:       appParam.LicenseId,
			Pn:              license.Pn,
			ActiveInfo:      appParam.ActiveInfo,
			Authcode:        license.Authcode,
			Number:          license.Number,
			ExpireTimestamp: license.ExpireTimestamp,
		}
		isValid, err := lic.checker.ValidateLicense(commonLic)
		if err != nil {
			return err
		}
		// record license status
		if err := lic.checker.RecordLicenseStatus(license.Pn, isValid, isAppBootCheck); err != nil {
			logrus.Warnf("record license status failed:[%v]", err)
		}
		if isValid {
			isAllFailed = false
		}
	}
	// if not all license check failed
	if !isAllFailed {
		return nil
	}

	// get license from api by service name
	apiLicenses, err := lic.retriever.LicenseWithActiveInfoBy(appParam.LicenseId, appParam.ServiceName)
	if err != nil {
		logrus.Warningf("get license from license API failed:[%v]", err)
	}
	for _, license := range apiLicenses {
		commonLic := &models.License{
			LicenseID:       appParam.LicenseId,
			Pn:              license.Pn,
			ActiveInfo:      appParam.ActiveInfo,
			Authcode:        license.AuthCode,
			Number:          license.Number,
			ExpireTimestamp: license.ExpireTimestamp,
		}
		isValid, err := lic.checker.ValidateLicense(commonLic)
		if err != nil {
			return err
		}
		// record license status
		if err := lic.checker.RecordLicenseStatus(license.Pn, isValid, isAppBootCheck); err != nil {
			logrus.Warnf("record license status failed:[%v]", err)
		}
	}
	return nil
}

func (lic *licenseManager) Version() string {
	return version
}

func (lic *licenseManager) LivVersion() (string, error) {
	version, err := liv.New().GetVersion()
	if err != nil {
		return "", fmt.Errorf("get liv version failed:[%v]", err)
	}
	return version, nil
}

// NewLicenseManager create a license manager with store,encryptor,licenseUrl
func New(store store.Store, encryptor encryptor.Encryptor, licenseUrl string) *licenseManager {
	return &licenseManager{
		checker:   checker.New(store, encryptor),
		retriever: retrieve.NewRetriever(licenseUrl),
	}
}

//NewWithDefaultEncryptor create a license manager with default encryptor
func NewWithDefaultEncryptor(store store.Store, salt, licenseUrl string) (*licenseManager, error) {
	etor, err := encryptor.New(salt)
	if err != nil {
		return nil, fmt.Errorf("init encryptor failed:%v", err)
	}
	return New(store, etor, licenseUrl), nil
}

//NewWithDefaultStore create a license manager with default store
func NewWithDefaultStore(encryptor encryptor.Encryptor, storeParams map[string]interface{}, licenseUrl string) (*licenseManager, error) {
	pg, err := postgre.New(storeParams)
	if err != nil {
		return nil, fmt.Errorf("create postgre store failed:[%v]", err)
	}
	return New(pg, encryptor, licenseUrl), nil
}

//NewLicenseManagerWithDefault create license manager with default component
func NewWithDefault(storeParams map[string]interface{}, aesSalt, licenseUrl string) (*licenseManager, error) {
	pg, err := postgre.New(storeParams)
	if err != nil {
		return nil, fmt.Errorf("create postgre store failed:[%v]", err)
	}
	etor, err := encryptor.New(aesSalt)
	if err != nil {
		return nil, fmt.Errorf("init encryptor failed:%v", err)
	}
	return New(pg, etor, licenseUrl), nil
}
