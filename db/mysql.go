package db

import (
	"database/sql"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewMysql(dataSource string, maxOpenConns, maxIdleConns, maxLifeTimeSeconds int) (*sql.DB, error) {
	db, err := sql.Open("mysql", dataSource)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(maxOpenConns)
	db.SetMaxIdleConns(maxIdleConns)
	db.SetConnMaxLifetime(time.Second * time.Duration(maxLifeTimeSeconds))

	err = db.Ping()
	return db, err
}
