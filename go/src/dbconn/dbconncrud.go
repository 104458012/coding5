// Date: 2015-06-13
// Authore: Coding5
// Purpose: 資料庫連線 
// Websit: http://www.coding5.com/works?transno=1000000033

package dbconn 

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var dbinfo string = "username:pwd123@tcp(127.0.0.1:3306)/DBNAME"

func GetQuery(stmt string) []string {

	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
    	panic(err.Error())
	}

	defer db.Close()

	counts, err := db.Query(stmt)
	if err != nil {
		panic(err.Error())
	}

	var ii int = 0
	for counts.Next() {
		ii++
	}

	rows, err := db.Query(stmt)
	if err != nil {
		panic(err.Error())
	}

	transno := make([]string, ii)
	var cc int = 0	
	for rows.Next() {
		err = rows.Scan(&transno[cc])
		cc++
	}

    if err = rows.Err(); err != nil {
        panic(err.Error())
    }

	return transno	
}

func GetInsert(stmt string) int64 {
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
    	panic(err.Error())
	}

	defer db.Close()

	ret, err := db.Exec(stmt)
	if err != nil {
		panic(err.Error())
	}

	id, err := ret.LastInsertId()

	return id
}

func GetUpdate(stmt string) int64 {
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
    	panic(err.Error())
	}

	defer db.Close()
	
	ret, err := db.Exec(stmt)
	if err != nil {
		panic(err.Error())
	}

	num, err := ret.RowsAffected()

	return num
}

func GetDelete(stmt string) int64 {
	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
    	panic(err.Error())
	}

	defer db.Close()
	
	ret, err := db.Exec(stmt)
	if err != nil {
		panic(err.Error())
	}

	num, err := ret.RowsAffected()

	return num
}
