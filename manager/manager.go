package manager

import (
	"github.com/ensaas/license-sdk/check"
	"github.com/ensaas/license-sdk/datasource"
	"github.com/ensaas/license-sdk/store"
)

type Manager struct {
	check.Checker
	store.Store
	datasource.DataSource
}

func NewManager() {

}
