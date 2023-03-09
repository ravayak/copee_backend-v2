package logs

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/ravayak/copee_backend/datasources/mysql"
	"github.com/stretchr/testify/assert"
)

func TestLogsDao(t *testing.T) {
	// Init mysql
	mysql.MysqlInit("copee-v2-test")

	// Create logs
	logCreateParams := mysql.CreateLogParams{
		LogEventDesc: sql.NullString{String: "Test", Valid: true},
		LogEventCode: sql.NullString{String: "Test", Valid: true},
		LogEventDate: sql.NullTime{Time: time.Now(), Valid: true},
	}

	resInsertClient, err := InsertLog(logCreateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	logId, err := resInsertClient.LastInsertId()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get log
	log, err := GetLog(int32(logId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, log.LogID, int32(logId))

	logUpdateParams := mysql.UpdateLogParams{
		LogEventDesc: sql.NullString{String: "Test updated", Valid: true},
		LogEventCode: sql.NullString{String: "Test updated", Valid: true},
		LogEventDate: sql.NullTime{Time: time.Now(), Valid: true},
		LogID:        int32(logId),
	}

	// Update log
	err = UpdateLog(logUpdateParams)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	// Get Updated log
	logs, err := ListLogs()
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
	assert.Equal(t, logs[0].LogEventDesc.String, "Test updated")

	// Delete log
	err = DeleteLog(int32(logId))
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}
}
