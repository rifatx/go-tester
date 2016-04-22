package sqltester

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
)

func Query() {
	db, err := sql.Open("mssql", "server=localhost;initial catalog=tester;user id=sa;password=osman.666")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sql conn ok")
	}

	rows, err := db.Query("SELECT name FROM sys.objects")

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("sql query ok")
		defer rows.Close()

		var name string

		for rows.Next() {
			err := rows.Scan(&name)

			if err != nil {
				fmt.Println(err)
			}

			fmt.Println(name)
		}

		err = rows.Err()

		if err != nil {
			fmt.Println(err)
		}
	}
}
