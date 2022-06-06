package ensaas

import (
	"github.com/ensaas/license-sdk/check"
	"github.com/ensaas/license-sdk/common"
	"github.com/ensaas/license-sdk/datasource"
	"github.com/ensaas/license-sdk/store"
	"github.com/robfig/cron"
	"log"
)

const zero_failed_time = 0

type checker struct {
	store.Store
	allowedFailedTimes int
	cronExpression     string
	isValid            bool
	trailDays          int // license invalid,user trail days
	licenseDataList    []*datasource.EnSaaSLicense
}

func NewChecker(store store.Store, allowedFailedTimes, trailDays int, cronExp string) check.Checker {
	return &checker{
		Store:              store,
		allowedFailedTimes: allowedFailedTimes,
		cronExpression:     cronExp,
		trailDays:          trailDays,
	}
}

// CheckLicense
func (c *checker) CheckLicense(lic *common.License) error {
	// initialize store
	if err := c.Store.Initialize(); err != nil {
		log.Printf("license initialize store failed:%s", err.Error())
		return err
	}

	//check license current status
	if err := c.checkStatus(); err != nil {
		log.Printf("license check status failed:%s", err.Error())
		return err
	}

	// collect license data
	var (
		dataSourceList    = datasource.ListDataSource()
		ensaasLicenseList = make([]*datasource.EnSaaSLicense, 0)
	)
	for _, ds := range dataSourceList {
		ensaaslicenses, err := ds.GetByEnSaaS(lic.ServiceName, lic.LicenseID)
		if err != nil {
			log.Printf("license load ensaas license data failed:%s", err.Error())
			return err
		}
		ensaasLicenseList = append(ensaasLicenseList, ensaaslicenses...)
	}
	c.licenseDataList = ensaasLicenseList

	// check license at regular time
	if err := c.checkLicenseAtRegularTime(); err != nil {
		log.Printf("license check at regular time failed:%s", err.Error())
		return err
	}
	return nil
}

// AvailableDays get trial left days
func (c *checker) AvailableDays() (int, error) {
	intVal, err := c.loadLicenseVal()
	if err != nil {
		return 0, err
	}

	inactiveDays := intVal / 3 //check three time a day
	if (intVal % 3) != 0 {
		inactiveDays++
	}
	trailDays := check.Sub(c.trailDays, inactiveDays)
	if trailDays < 0 {
		return 0, nil
	}
	return trailDays, nil
}

func (c *checker) IsValid() bool {
	return c.isValid
}

func (c *checker) Run() {
	if len(c.licenseDataList) == 0 {
		if err := c.addFailedTimes(); err != nil {
			log.Printf("record license info failed:%s", err.Error())
			return
		}
		return
	}

	hasValidLicense := false
	for _, licenseData := range c.licenseDataList {
		checkData := &license{
			activeInfo:      licenseData.ActiveInfo,
			licenseID:       licenseData.LicenseID,
			pn:              licenseData.Pn,
			authCode:        licenseData.AuthCode,
			expireTimestamp: licenseData.ExpireTimestamp,
			number:          licenseData.Number}
		if err := checkData.validateAuthCode(); err != nil {
			continue
		}
		hasValidLicense = true
		break
	}
	if !hasValidLicense {
		if err := c.addFailedTimes(); err != nil {
			log.Printf("record license info failed:%s", err.Error())
			return
		}
		return
	}

	// reset failed times
	if err := c.resetFailedTimes(); err != nil {
		log.Printf("record license info failed:%s", err.Error())
	}
}

func (c *checker) resetFailedTimes() error {
	c.isValid = true
	if err := c.Store.Save(zero_failed_time); err != nil {
		return err
	}
	return nil
}

func (c *checker) addFailedTimes() error {
	oldFailedTimes, err := c.loadLicenseVal()
	if err != nil {
		c.isValid = false
		return err
	}
	if oldFailedTimes >= c.allowedFailedTimes {
		c.isValid = false
		return nil
	}

	newFailedTimes := check.Add(oldFailedTimes, 1)
	if newFailedTimes >= c.allowedFailedTimes {
		c.isValid = false
	}
	if err := c.Store.Save(newFailedTimes); err != nil {
		return err
	}
	return nil
}

func (c *checker) loadLicenseVal() (int, error) {
	val, err := c.Store.Load()
	if err != nil {
		log.Printf("license load store value failed:%s", err.Error())
		return 0, err
	}

	intVal, ok := val.(int)
	if !ok {
		log.Printf("license store value %s format incorrect", val)
		return 0, err
	}
	return intVal, nil
}

// checkLicenseAtRegularTime
func (c *checker) checkLicenseAtRegularTime() error {
	cronObj := cron.New()
	if err := cronObj.AddJob(c.cronExpression, c); err != nil {
		log.Printf("license start check license job failed:%s", err.Error())
		return err
	}
	cronObj.Start()
	return nil
}

// checkStatus
func (c *checker) checkStatus() error {
	intVal, err := c.loadLicenseVal()
	if err != nil {
		return err
	}
	if intVal >= c.allowedFailedTimes {
		c.isValid = false
	} else {
		c.isValid = true
	}
	return nil
}
