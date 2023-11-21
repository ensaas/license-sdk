package sqlite

import (
	"fmt"
	"github.com/ensaas/license-sdk/store"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

const (
	DatabaseName = "database"
)

type sqlite struct {
	DBConn *gorm.DB
}

type license struct {
	Id          int    `json:"id" gorm:"AUTO_INCREMENT"`
	Pn          string `gorm:"uniqueIndex"`
	CheckResult string
}

func (l *license) TableName() string {
	return "license"
}

func (s *sqlite) Save(val map[string]interface{}) error {
	pn, ok := val["pn"]
	if !ok {
		return fmt.Errorf("pn not exist")
	}
	cr, ok := val["checkResult"]
	if !ok {
		return fmt.Errorf("check result not found")
	}

	lic := &license{}
	if err := s.DBConn.Where("pn=?", pn).First(lic).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return fmt.Errorf("sqlite take license failed:[%v]", err)
		}
		lic.Pn = pn.(string)
	}
	lic.CheckResult = cr.(string)

	if err := s.DBConn.Save(lic).Error; err != nil {
		return fmt.Errorf("sqlite save license failed:[%v]", err)
	}
	return nil
}

func (s *sqlite) Load(params map[string]interface{}) (interface{}, error) {
	lic := &license{}
	if err := s.DBConn.Where(params).First(lic).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("postgres load license failed:[%v]", err)
	}
	return lic.CheckResult, nil
}

func (s *sqlite) Initialize(params map[string]interface{}) error {
	return nil
}

func New(params map[string]interface{}) (store.Store, error) {
	dbConn, err := gorm.Open("sqlite3", params[DatabaseName])
	if err != nil {
		return nil, err
	}
	dbConn.LogMode(false)
	dbConn.SingularTable(true)
	// initialize license table
	dbConn.AutoMigrate(&license{})

	return &sqlite{DBConn: dbConn}, nil
}
