package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DbConn *sql.DB
var server = "dynamicsbackup.database.windows.net"
var port = 1433
var user = "shreejit"
var password = "123456798rockoN."
var database = "backup"

func SetupDatabase() {
	var err error

	DbConn, err = sql.Open("mysql", "shreejit@dynamicsbackup:123456798rockoN.@tcp(dynamicsbackup.mysql.database.azure.com:3306)/svna")

	if err != nil {
		fmt.Println("this is an error")
		log.Fatal(err)

		return
	}

	fmt.Println("Connection successful")

}
