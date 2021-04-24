package db

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var err error

type userDB struct {
	Conn *sql.DB
}

type db struct {
	UserDB userDB
}

var DBConn = &db{}

// MYSQL 앱 부팅 시에만 실행 해서 연결
func MySQLConnect(connType uint16, dsn string) {
	switch connType {
	case 0:
		DBConn.UserDB.Conn, err = sql.Open("mysql", dsn)
		if err != nil || DBConn.UserDB.Conn.Ping() != nil {
			log.Fatal(err)
		}
		DBConn.UserDB.Conn.SetConnMaxLifetime(time.Minute * 3)
		DBConn.UserDB.Conn.SetMaxOpenConns(10)
		DBConn.UserDB.Conn.SetMaxIdleConns(10)
		fmt.Println("UserDB 연결 완료!")
	}
}
