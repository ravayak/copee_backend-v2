package dateutils

import (
	"errors"
	"strconv"
	"strings"
	"time"

	"github.com/ravayak/copee_backend/app/data/global"
	mathutils "github.com/ravayak/copee_backend/app/utils/math"
)

const (
	apiDateLayout     = "2006-01-02T15:04:05Z"
	apiDBSQLLayout    = "2006-01-02 15:04:05"
	INITIAL_DATE_FROM = "2021-01-01"
	INITIAL_DATE_TO   = "2021-12-12"
)

func GetNow() time.Time {
	// Wrong hour with UTC ?
	return time.Now()
}

func GetNowString() string {
	dateCreated := GetNow().Format(apiDateLayout)
	return dateCreated
}

func GetNowDBFormatString() string {
	return GetNow().Format(apiDBSQLLayout)
}

func SetPeriod(day, month int, periodDay int) (string, string) {
	var monthFrom string
	var monthTo string

	if day >= periodDay {
		monthFrom = strconv.Itoa(month)
		monthTo = strconv.Itoa(month + 1)
	}

	if day < periodDay {
		monthFrom = strconv.Itoa(month - 1)
		monthTo = strconv.Itoa(month)
	}

	return monthFrom, monthTo
}

// getDate get year/month/day from a date
func GetDate(date string) (int, int, int, error) {
	dateSpaceSplited := strings.Split(date, " ")
	year := 0
	month := 0
	day := 0
	var err error

	if len(dateSpaceSplited) > 0 {
		yearMonthDay := strings.Split(dateSpaceSplited[0], "-")
		if len(yearMonthDay) > 2 {
			year, err = mathutils.GetIntValue(yearMonthDay[0])
			if err != nil {
				return year, month, day, err
			}
			month, err = mathutils.GetIntValue(yearMonthDay[1])
			if err != nil {
				return year, month, day, err
			}
			day, err = mathutils.GetIntValue(yearMonthDay[2])
			if err != nil {
				return year, month, day, err
			}
		} else {
			err = errors.New("invalid date splited")
		}
	} else {
		err = errors.New("invalid date splited")
	}

	return year, month, day, err
}

func FormatValuesToMySqlDate(year, day int, month string) string {
	if month == "0" {
		year -= 1
		month = "12"
	}
	return strconv.Itoa(year) + "-" + month + "-" + strconv.Itoa(global.DAY_PERIOD)
}
