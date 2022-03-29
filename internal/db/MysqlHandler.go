package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"go-mysql/internal/structs"
	"log"
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

func (h *MySQLHandler) CallSP(sp string, obj interface{}) (*[]byte, error) {
	input, err := json.Marshal(obj)
	if err != nil {
		return nil, err
	}
	var output *[]byte
	startTime := time.Now()
	if obj != nil {
		query := fmt.Sprintf("CALL %s (?)", sp)
		err = h.Conn.QueryRow(query, string(input)).Scan(&output)
	} else {
		query := fmt.Sprintf("CALL %s ()", sp)
		err = h.Conn.QueryRow(query).Scan(&output)
	}
	called := fmt.Sprintf("CALL %s ('%s'); (%s)", sp, string(input), time.Since(startTime))
	log.Println("[mysql]" + called)

	if err == sql.ErrNoRows {
		return nil, nil
	}

	if err != nil {
		return nil, errors.New(called + " " + err.Error())
	}

	var respuestaMySQL structs.ResponseMySQL

	err = json.Unmarshal(*output, &respuestaMySQL)
	if err != nil {
		return nil, err
	}

	if respuestaMySQL.Error != nil {
		return nil, errors.New(*(respuestaMySQL.Error))
	}

	b, err := json.Marshal(respuestaMySQL.Respuesta)

	if err != nil {
		return nil, err

	}
	return &b, nil

}
