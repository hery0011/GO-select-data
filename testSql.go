package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql" //mila instalena aloha ito 'go get -u github.com/go-sql-driver/mysql'
)

func main() {
	fmt.Println("resulta sql select")

	db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/golangtest")

	if err != nil {
		panic(err.Error())
	}

	rows, err := db.Query("SELECT * FROM user")

	checkErr(err)

	for rows.Next() {
		var id int
		var nom string
		var prenom string
		var tel string
		err = rows.Scan(&id, &nom, &prenom, &tel)
		checkErr(err)
		fmt.Println(id)
		fmt.Println(nom)
		fmt.Println(prenom)
		fmt.Println(tel)
	}

	defer db.Close()
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
