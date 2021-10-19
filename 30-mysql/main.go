package main

import (
	"database/sql"

	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	dsn := "root:root@tcp(127.0.0.1:3306)/goday"

	db, err := sql.Open("mysql", dsn)

	con, err := sqlx.Connect("mysql", dsn)
	if err != nil {
		fmt.Printf("open %s failed,err : %s\n", dsn, err)
		return
	}

	fmt.Println("connect mysql success")
	err = db.Ping()
	if err != nil {
		fmt.Printf("open %s failed,err : %s\n", dsn, err)
	}

	fmt.Println("connect mysql success")

}
