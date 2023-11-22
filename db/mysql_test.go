package db

import (
	"database/sql"
	"log"
	"testing"
)

var (
	mydb *sql.DB
)

func TestNewMysql(t *testing.T) {
	dbtmp, err := NewMysql("user:password@tcp(127.0.0.1:3306)/DATABASE?charset=utf8", 4, 4, 0)
	if err != nil {
		t.Errorf("connect to mysql error:%v", err)
	}

	mydb = dbtmp
}

func TestQuery(t *testing.T) {
	id := 0
	row := mydb.QueryRow("select CONNECTION_ID();")
	err := row.Scan(&id)
	if err != nil {
		t.Errorf("row.Scan error: %v\n", err)
	}
	log.Printf("connection id: %d\n", id)
}
