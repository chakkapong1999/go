package db

import (
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

func connectDB() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "root",
		Addr:   "localhost:3307",
		DBName: "shopping-table",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	defer db.Close()
	if err != nil {
		fmt.Println("Fail to Connect to DB.")
	} else {
		fmt.Println("Connect to DB Success.")
	}
}
