package database

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/trinodb/trino-go-client/trino"
)

const (
	DB_TRINO    = "trino"
	DB_MYSQL    = "mysql"
	DB_POSTGRES = "postgres"
)
