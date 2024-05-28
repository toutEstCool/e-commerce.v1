package main

import (
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/toutEstCool/e-commerce.v1/cmd/api"
	"github.com/toutEstCool/e-commerce.v1/cmd/db"
)

func main() {
	db, err := db.NewMySqlStorage(mysql.Config{
		User:                 "root",
		Passwd:               "root",
		Addr:                 "127.0.1:3306",
		DBName:               "ecom",
		Net:                  "tcp",
		AllowNativePasswords: true,
		ParseTime:            true,
	})

	server := api.NewAPIServer(":8080", db)
	if err := server.Run(); err != nil {
		log.Fatal(err)
	}
}
