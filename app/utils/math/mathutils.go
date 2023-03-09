package mathutils

import (
	"errors"
	"math"
	"strconv"
)

func IsValidNumberParam(param string) bool {
	if param != "" {
		_, err := GetIntValue(param)
		if err != nil {
			return false
		}
	}

	return true
}

func GetIntValue(value string) (int, error) {
	intValue, err := strconv.Atoi(value)
	if err != nil {
		err := errors.New("invalid ID number")
		return 0, err
	}
	return intValue, nil
}

func ComputeAmount(comClientPayment, comCeePayment, coef float64) float64 {
	return (comClientPayment + comCeePayment) * coef
}

func ComputePercentage(n float64, percentage float64) float64 {
	return n * (percentage / 100)
}

func RoundToNearestNDecimals(n float64, decimals int64) float64 {
	return n * float64(decimals) / float64(decimals)
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
