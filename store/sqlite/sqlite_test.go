package sqlite

import (
	"fmt"
	"github.com/ensaas/license-sdk/store"
	"github.com/jinzhu/gorm"
	"log"
	"reflect"
	"testing"
)

func TestSqlite_Load(t *testing.T) {
	m := map[string]interface{}{DatabaseName: "/home/sqlite/license"}
	s, _ := New(m)

	if err := s.Save(map[string]interface{}{"pn": "123456", "checkResult": "ccvvv"}); err != nil {
		log.Fatalf(err.Error())
	}

	val, err := s.Load(nil)
	if err != nil {
		log.Fatalf(err.Error())
	}
	fmt.Println(val)
}

func TestNew(t *testing.T) {
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    store.Store
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := New(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("New() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlite_Initialize(t *testing.T) {
	type fields struct {
		DBConn *gorm.DB
	}
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlite{
				DBConn: tt.fields.DBConn,
			}
			if err := s.Initialize(tt.args.params); (err != nil) != tt.wantErr {
				t.Errorf("Initialize() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_sqlite_Load(t *testing.T) {
	type fields struct {
		DBConn *gorm.DB
	}
	type args struct {
		params map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlite{
				DBConn: tt.fields.DBConn,
			}
			got, err := s.Load(tt.args.params)
			if (err != nil) != tt.wantErr {
				t.Errorf("Load() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Load() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_sqlite_Save(t *testing.T) {
	type fields struct {
		DBConn *gorm.DB
	}
	type args struct {
		val map[string]interface{}
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &sqlite{
				DBConn: tt.fields.DBConn,
			}
			if err := s.Save(tt.args.val); (err != nil) != tt.wantErr {
				t.Errorf("Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
