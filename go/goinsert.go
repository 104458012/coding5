// Date: 2015-06-08
// Authore: Coding5
// Purpose: Go 程式呼叫資料庫連線INSERT 操作 
// Websit: http://www.coding5.com/works?transno=1000000033

package main

import "dbconn"
import "fmt"

func main() {
	var stmt string = "INSERT INTO codeing5 (codeA) VALUE ('codeA')"

	id := dbconn.GetInsert(stmt)

	fmt.Println(id)
}
