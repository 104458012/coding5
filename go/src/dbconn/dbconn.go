package dbconn 

import "database/sql"
import _ "github.com/go-sql-driver/mysql"

var dbinfo string = "mysql:echoo710426@tcp(127.0.0.1:3306)/LETO"

func GetQuery(stmt string) []string {

	db, err := sql.Open("mysql", dbinfo)
	if err != nil {
    	panic(err.Error())
	}

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
