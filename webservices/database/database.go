package database

import (
	"database/sql"
	"fmt"
	"log"
	"time"
)

const (
	DBHost  = "127.0.0.1"
	DBPort  = ":3306"
	DBUser  = "root"
	DBPass  = ""
	DBDbase = "inventorydb"
)

var DbConn *sql.DB

func SetupDatabase() {
	// dbConn := fmt.Sprintf("%s:%s@tcp(%s)/%s", DBUser, DBPass, DBHost, DBDbase)
	// fmt.Println(dbConn)
	// db, err := sql.Open("mysql", dbConn)

	// if err != nil {
	// 	log.Println("Couldn't connect!")
	// }

	// log.Println("Success")
	// log.Println(db)
	var err error
	DbConn, err = sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/inventorydb")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("success")
	DbConn.SetMaxOpenConns(3)
	DbConn.SetMaxIdleConns(3)
	DbConn.SetConnMaxLifetime(60 * time.Second)
}
