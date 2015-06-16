// Date: 2015-06-08
// Authore: Coding5
// Purpose: Go 程式呼叫資料庫連線query 操作 
// Websit: http://www.coding5.com/works?transno=1000000032

package main

import "dbconn"
import "fmt"

func main() {
	var stmt string = "UPDATE coding5 SET codeB = '456' WHERE codeA = 'codeA'"

	num := dbconn.GetUpdate(stmt)

	fmt.Println(num)
}
