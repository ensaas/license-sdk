package checker

import (
	"fmt"
	"github.com/ensaas/license-sdk/encryptor"
	"github.com/ensaas/license-sdk/liv"
	"github.com/ensaas/license-sdk/models"
	"github.com/ensaas/license-sdk/store"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

const (
	allowedFailedTimes = 42
	availableDays      = 12
)

// Checker for check license validator
type Checker interface {
	// GetAvailableDays get license available days
	GetAvailableDays(pn string) (int, error)
	// ValidateLicense check license is valid
	ValidateLicense(lic *models.License) (bool, error)
	// RecordLicenseStatus record license status
	RecordLicenseStatus(pn string, isValid, isAppBootCheck bool) error
}

type defaultChecker struct {
	store      store.Store
	livWrapper liv.Wrapper
	encryptor  encryptor.Encryptor
}

func New(store store.Store, encryptor encryptor.Encryptor) Checker {
	return &defaultChecker{
		store:      store,
		livWrapper: liv.New(),
		encryptor:  encryptor,
	}
}

// ValidateLicense is validate license authcode is correct and check expire timestamp
func (d *defaultChecker) ValidateLicense(lic *models.License) (bool, error) {
	isValid, err := d.livWrapper.CheckAuthCode(lic)
	if err != nil {
		logrus.Errorf("checker check authcode failed:[%v]", err)
		return false, fmt.Errorf("checker check authcode failed:[%v]", err)
	}
	return isValid, nil
}

// RecordLicenseStatus record license status
func (d *defaultChecker) RecordLicenseStatus(pn string, isValid, isAppBootCheck bool) error {
	// get license check failed times
	failedTimes, err := d.loadFailedTimes(pn)
	if err != nil && !isValid {
		return fmt.Errorf("checker load license status failed:[%v]", err)
	} else if err != nil && isValid {
		logrus.Warningf("load license status failed:[%v]", err)
		return nil
	}

	if isValid && failedTimes <= 0 {
		return nil
	} else if isValid && failedTimes > 0 {
		if err := d.resetFailedTimes(pn); err != nil {
			logrus.Warningf("reset license status failed:[%v]", err)
		}
		return nil
	} else if !isValid && failedTimes >= allowedFailedTimes {
		return fmt.Errorf("check license failed times already over allowed times")
	} else if !isValid && failedTimes < allowedFailedTimes {
		if isAppBootCheck { // if is first boot check,dont accumulate failed times
			logrus.Warningf("check license failed,pn is [%s]", pn)
			return nil
		}
		if err := d.saveFailedTimes(pn, failedTimes); err != nil {
			return err
		}
		return nil
	}
	return nil
}

// GetStatus is get status of license by pn
func (d *defaultChecker) GetAvailableDays(pn string) (int, error) {
	failedTimes, err := d.loadFailedTimes(pn)
	if err != nil {
		return 0, err
	}
	if failedTimes > allowedFailedTimes {
		logrus.Warnf("license failed times already over allowed failed times...")
		return 0, nil
	}

	inactiveDays := failedTimes / 3 //check three time a day
	if (failedTimes % 3) != 0 {
		inactiveDays++
	}
	avaiDays := availableDays - inactiveDays
	if avaiDays < 0 {
		return 0, nil
	}
	return avaiDays, nil
}

// saveFailedTimes for save license failed time and accumulate is
func (d *defaultChecker) saveFailedTimes(pn string, failedTimes int) error {
	cipherFailedTimes, err := d.encryptor.Encrypt([]byte(fmt.Sprintf("%s:%d", pn, failedTimes+1)))
	if err != nil {
		return fmt.Errorf("encrypt license status error:[%v]", err)
	}
	if err := d.store.Save(map[string]interface{}{pn: cipherFailedTimes}); err != nil {
		return fmt.Errorf("save license status failed:[%v]", err)
	}
	return nil
}

// resetFailedTimes for reset failed times when check license success
func (d *defaultChecker) resetFailedTimes(pn string) error {
	cipherFailedTimes, err := d.encryptor.Encrypt([]byte(fmt.Sprintf("%s:%d", pn, 0)))
	if err != nil {
		return fmt.Errorf("encrypt license status error:[%v]", err)
	}
	if err := d.store.Save(map[string]interface{}{pn: cipherFailedTimes}); err != nil {
		return fmt.Errorf("save license status failed:[%v]", err)
	}
	return nil
}

// loadFailedTimes for get check license already failed times
func (d *defaultChecker) loadFailedTimes(pn string) (int, error) {
	cipher, err := d.store.Load(map[string]interface{}{"pn": pn})
	if err != nil {
		return 0, fmt.Errorf("get license status failed:[%v]", err)
	}
	if cipher == nil {
		return 0, nil
	}

	val, err := d.encryptor.Decrypt(cipher.(string))
	if err != nil {
		return 0, fmt.Errorf("get license status failed:[%v]", err)
	}
	arr := strings.Split(val.(string), ":")
	if len(arr) != 2 {
		return 0, fmt.Errorf("get license status failed: check license data [%s] not correct", val)
	}
	failedTimes, err := strconv.Atoi(arr[1])
	if err != nil {
		return 0, fmt.Errorf("get license status failed: [%v]", err)
	}
	return failedTimes, nil
}
