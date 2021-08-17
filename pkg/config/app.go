package config

import (
	"gorm.io/driver/postgres"
	//"database/sql"
	"fmt"
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

const (
	host = "localhost"
	dbname = "todo"
	port = 5432
	user =  "postgres"
	password = "enoch@1992"
)

func Connect() {
	psqlConn := fmt.Sprintf("host = %s port= %d user= %s password= %s dbname= %s sslmode=disable", host, port, user, password, dbname)
	var err error
	db, err = gorm.Open(postgres.Open(psqlConn), &gorm.Config{})
	checkErr(err)
}

func GetDB() *gorm.DB {
	return db
}
func checkErr(err error){
	if err!= nil{
		panic(err)
	}
}
