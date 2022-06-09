package postgre

const (
	Username     = "username"
	Password     = "password"
	Host         = "host"
	Port         = "port"
	DBName       = "dbname"
	SSLMode      = "sslmode"
	MaxIdleConns = "maxIdleConns"
	MaxOpenConns = "maxOpenConns"

	defaultSSLMode      = "disable"
	defaultMaxIdleConns = 10
	defaultMaxOpenConns = 10
)
