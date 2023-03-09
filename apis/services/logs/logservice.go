package log

import (
	"database/sql"

	"github.com/ravayak/copee_backend/apis/domain/logs"
	"github.com/ravayak/copee_backend/datasources/mysql"
)

var (
	LogService logServiceInterface = &logService{}
)

type logService struct{}
type logServiceInterface interface {
	InsertLog(logParams mysql.CreateLogParams) (sql.Result, error)
	UpdateLog(logParams mysql.UpdateLogParams) error
	DeleteLog(logId int32) error
	GetLog(logId int32) (mysql.Log, error)
	ListLogs() ([]mysql.Log, error)
}

func (cs *logService) InsertLog(logParams mysql.CreateLogParams) (sql.Result, error) {
	return logs.InsertLog(logParams)
}

func (cs *logService) UpdateLog(logParams mysql.UpdateLogParams) error {
	return logs.UpdateLog(logParams)
}

func (cs *logService) DeleteLog(logId int32) error {
	return logs.DeleteLog(logId)
}

func (cs *logService) GetLog(logId int32) (mysql.Log, error) {
	return logs.GetLog(logId)
}

func (cs *logService) ListLogs() ([]mysql.Log, error) {
	return logs.ListLogs()
}
