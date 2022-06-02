package Setting

import (
	"database/sql"

	"fmt"
)

//Init ...

func Constr() *sql.DB {

	Condb, errdb := sql.Open("mssql", "server=172.104.170.123;database=test;user id=psiber-sql-admin;password=ZhJhE6C&cx87W7@!;")

	if errdb != nil {

		fmt.Println(" Error open db:", errdb.Error())

	}

	return Condb

}
