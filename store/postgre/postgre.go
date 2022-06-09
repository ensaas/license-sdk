package postgre

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type postgres struct {
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

func (p *postgres) Load(params map[string]interface{}) (interface{}, error) {
	lic := &license{}
	if err := p.DBConn.Where(params).First(lic).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, fmt.Errorf("postgres load license failed:[%v]", err)
	}
	return lic.CheckResult, nil
}

func (p *postgres) Save(val map[string]interface{}) error {
	pn, ok := val["pn"]
	if !ok {
		return fmt.Errorf("pn not exist")
	}
	cr, ok := val["checkResult"]
	if !ok {
		return fmt.Errorf("check result not found")
	}

	lic := &license{}
	if err := p.DBConn.Where("pn=?", pn).First(lic).Error; err != nil {
		if err != gorm.ErrRecordNotFound {
			return fmt.Errorf("postgres take license failed:[%v]", err)
		}
		lic.Pn = pn.(string)
	}
	lic.CheckResult = cr.(string)

	if err := p.DBConn.Save(lic).Error; err != nil {
		return fmt.Errorf("postgres save license failed:[%v]", err)
	}
	return nil
}

func (p *postgres) Initialize(params map[string]interface{}) error {
	return nil
}

func New(params map[string]interface{}) (*postgres, error) {
	host, ok := params[Host]
	if !ok {
		return nil, fmt.Errorf("postgre host not found")
	}
	port, ok := params[Port]
	if !ok {
		port = "5432"
	}
	username, ok := params[Username]
	if !ok {
		return nil, fmt.Errorf("postgre username not found")
	}
	password, ok := params[Password]
	if !ok {
		return nil, fmt.Errorf("postgre password not found")
	}
	dbName, ok := params[DBName]
	if !ok {
		return nil, fmt.Errorf("postgre dbname not found")
	}
	sslMode, ok := params[SSLMode]
	if !ok {
		sslMode = defaultSSLMode
	}
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		host, port, username, dbName, password, sslMode)

	dbConn, err := gorm.Open("postgres", dbUrl)
	if err != nil {
		panic("open postgres connection failed")
	}

	dbConn.LogMode(false)
	dbConn.SingularTable(true)
	maxIdleConns, ok := params[MaxIdleConns]
	if !ok {
		maxIdleConns = defaultMaxIdleConns
	} else {
		_, ok := maxIdleConns.(int)
		if !ok {
			return nil, fmt.Errorf("postgre maxIdleConns format error")
		}
	}
	maxOpenConns, ok := params[MaxOpenConns]
	if !ok {
		maxOpenConns = defaultMaxOpenConns
	} else {
		_, ok := maxOpenConns.(int)
		if !ok {
			return nil, fmt.Errorf("postgre maxOpenConns format error")
		}
	}
	dbConn.DB().SetMaxIdleConns(maxIdleConns.(int))
	dbConn.DB().SetMaxOpenConns(maxOpenConns.(int))

	// initialize license table
	dbConn.AutoMigrate(&license{})

	return &postgres{DBConn: dbConn}, nil
}
