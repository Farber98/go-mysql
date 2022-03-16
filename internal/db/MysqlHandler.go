package db

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

type MySQLHandler struct {
	Conn *sql.DB
}

func InitDb(user string, password string, host string, dbName string, maxOpenConnections, maxIddleConns, connMaxTTL int) (*MySQLHandler, error) {
	connString := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbName)

	conn, err := sql.Open("mysql", connString)
	if err != nil {
		return nil, err
	}
	conn.SetMaxOpenConns(maxOpenConnections)
	conn.SetConnMaxLifetime(time.Duration(connMaxTTL) * time.Second)
	conn.SetMaxIdleConns(maxIddleConns)
	return &MySQLHandler{conn}, nil
}
