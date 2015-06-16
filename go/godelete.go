// Date: 2015-06-08
// Authore: Coding5
// Purpose: Go 程式呼叫資料庫連線DELETE 操作 
// Websit: http://www.coding5.com/works?transno=1000000033

package main

import "dbconn"
import "fmt"

func main() {
	var stmt string = "DELETE FROM coding5 WHERE codeA ='codeA'"

	num := dbconn.GetDelete(stmt)

	fmt.Println(num)
}
