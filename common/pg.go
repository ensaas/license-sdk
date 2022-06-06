package common

import "errors"

type Postgres struct {
	Username string
	Password string
	Database string
	Host     string
	Port     string
	SSLMode  string
}

func NewPostgres(user, pwd, dbName, host, port, sslMode string) (*Postgres, error) {
	if len(user) == 0 {
		return nil, errors.New("license pg username is required")
	}
	if len(pwd) == 0 {
		return nil, errors.New("license pg password is required")
	}
	if len(dbName) == 0 {
		return nil, errors.New("license pg database is required")
	}
	if len(host) == 0 {
		return nil, errors.New("license pg host is required")
	}
	if len(port) == 0 {
		return nil, errors.New("license pg port is required")
	}
	if len(sslMode) == 0 {
		sslMode = "disable"
	}
	return &Postgres{
		Username: user,
		Password: pwd,
		Database: dbName,
		Host:     host,
		Port:     port,
		SSLMode:  sslMode,
	}, nil
}
