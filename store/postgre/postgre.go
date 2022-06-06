package postgre

import (
	"fmt"
	"github.com/ensaas/license-sdk/common"
	"github.com/ensaas/license-sdk/store"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"sync"
)

var once sync.Once

const PostgresStoreType = "postgres"

type postgres struct {
	DBConn *gorm.DB
}

func (p *postgres) Load() (interface{}, error) {
	return nil, nil
}

func (p *postgres) Save(val interface{}) error {
	return nil
}

func (p *postgres) Initialize() error {
	p.DBConn.AutoMigrate(store.License{})
	return nil
}

func New(pg *common.Postgres) store.Store {
	dbUrl := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		pg.Host, pg.Port, pg.Username, pg.Database, pg.Password, pg.SSLMode)

	var (
		err    error
		dbConn *gorm.DB
	)
	once.Do(func() {
		dbConn, err = gorm.Open("postgres", dbUrl)
		if err != nil {
			panic("open postgres connection failed")
		}
		dbConn.LogMode(false)
		dbConn.SingularTable(true)
	})

	return &postgres{DBConn: dbConn}
}
