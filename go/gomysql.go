// Date: 2015-06-07
// Authore: Coding5
// Purpose: Go 程式連接MySQL 資料庫 
// Websit: http://www.coding5.com/works?transno=1000000031

package main

import "database/sql"
import _ "github.com/go-sql-driver/mysql"
import "fmt"

func main() {
    db, err := sql.Open("mysql", "username:paw123@tcp(127.0.0.1:3306)/DBNAME")
    if err != nil {
        panic(err.Error())
    }

    defer db.Close()
	
    rows, err := db.Query("SELECT codeA FROM coding5")
    if err != nil {
        panic(err.Error())
    }

	for rows.Next() {
		var codeA string

		err = rows.Scan(&codeA)

		fmt.Println(codeA)
	}

    if err = rows.Err(); err != nil {
        panic(err.Error())
    }

}
