package postgre

import (
	"fmt"
	"log"
	"testing"
)

var m = map[string]interface{}{
	Host:     "localhost",
	Port:     "5432",
	Username: "postgres",
	Password: "123456",
	DBName:   "license",
}

func TestPostgres_Load(t *testing.T) {
	pg, err := New(m)
	if err != nil {
		log.Fatalf(err.Error())
	}
	if err := pg.Save(map[string]interface{}{"pn": "123456", "checkResult": "ccvvv"}); err != nil {
		log.Fatalf(err.Error())
	}

	val, err := pg.Load(nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(val)
}
