package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	"github.com/ravayak/copee_backend/app/utils/logger"
)

const (
	mysqlUsersUsername = "root"
	mysqlUsersPassword = ""
	mysqlUsersHost     = "localhost:3306"
)

var (
	Db        *sql.DB
	QueriesDb *Queries
)

func MysqlInit(dbName string) {
	if dbName != "" {
		dataSourceName := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=true",
			mysqlUsersUsername,
			mysqlUsersPassword,
			mysqlUsersHost,
			dbName,
		)

		var err error
		Db, err = sql.Open("mysql", dataSourceName)
		if err != nil {
			log.Fatalf("Connection to the mysql DB couldn't be established, error:%v", err)
		}

		QueriesDb = New(Db)

		if err = Db.Ping(); err != nil {
			log.Fatalf("Connection to the mysql DB couldn't be established, error: %v", err)
		}

		mysql.SetLogger(logger.GetLogger())
		log.Println("database successfully configured")
	} else {
		log.Fatalf("Connection to the mysql DB couldn't be established, no dbName provided")
	}
}
