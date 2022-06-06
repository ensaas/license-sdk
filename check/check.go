package check

import (
	"github.com/ensaas/license-sdk/common"
)

type Checker interface {
	CheckLicense(lic *common.License) error
	AvailableDays() (int, error)
	IsValid() bool
}
