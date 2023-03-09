package mysqlutils

import (
	"fmt"
	"log"

	mysqldata "github.com/ravayak/copee_backend/app/data/mysql"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

func DeleteDatasFromDB() {
	for _, table := range mysqldata.Tables {
		query := fmt.Sprintf("DELETE FROM `%s`", table)
		_, err := mysql.Db.Exec(query)
		if err != nil {
			log.Fatal(err)
		}
	}

	fmt.Printf("Datas from all tables were deleted successfully\n")
}

func CompleteQuery(query string, args []int32) string {
	if len(args) == 0 {
		return query
	}
	query += "("
	if len(args) == 1 {
		return query + "?)"
	}
	for i := range args {
		query += "?"
		if i+1 < len(args) {
			query += ","
		}
	}
	query += ")"
	return query
}
