package srp

import (
	"github.com/ensaas/license-sdk/check"
	"github.com/ensaas/license-sdk/common"
	"github.com/ensaas/license-sdk/store"
)

type checker struct {
	store.Store
	allowedFailedTimes int
	cronExpression     string
	isValid            bool
	trailDays          int
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
	return nil
}

// AvailableDays get trial left days
func (c *checker) AvailableDays() (int, error) {
	return 0, nil
}

func (c *checker) IsValid() bool {
	return false
}
