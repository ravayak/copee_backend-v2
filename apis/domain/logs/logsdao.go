package logs

import (
	"context"
	"database/sql"

	"github.com/ravayak/copee_backend/datasources/mysql"
)

func ListLogs() ([]mysql.Log, error) {
	ctx := context.Background()
	logs, err := mysql.QueriesDb.ListLogs(ctx)
	return logs, err
}

func GetLog(logId int32) (mysql.Log, error) {
	ctx := context.Background()
	log, err := mysql.QueriesDb.GetLog(ctx, logId)
	return log, err
}

func InsertLog(logParams mysql.CreateLogParams) (sql.Result, error) {
	ctx := context.Background()
	res, err := mysql.QueriesDb.CreateLog(ctx, logParams)
	return res, err
}

func UpdateLog(logParams mysql.UpdateLogParams) error {
	ctx := context.Background()
	err := mysql.QueriesDb.UpdateLog(ctx, logParams)
	return err
}

func DeleteLog(logId int32) error {
	ctx := context.Background()
	err := mysql.QueriesDb.DeleteLog(ctx, logId)
	return err
}
