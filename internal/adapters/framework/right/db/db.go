package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(driverName, dataSourceName string) (*Adapter, error) {
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatalf("db ping failure: %v", err)
	}

	return &Adapter{db: db}, nil
}

func (dba Adapter) CloseDbConnection() {
	err := dba.db.Close()
	log.Fatalf("db close failure: %v", err)
}

func (dba Adapter) AddToHistory(answer int32, operation string) error {
	insertQuery := sq.Insert("arith_history").Columns("date", "answer", "operation").
		Values(time.Now(), answer, operation)
	insertQuery = insertQuery.RunWith(dba.db)
	_, err := insertQuery.Exec()
	if err != nil {
		return err
	}
	return nil
}
