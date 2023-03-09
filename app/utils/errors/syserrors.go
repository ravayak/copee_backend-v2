package syserrors

import (
	"fmt"
)

var (
	SysError sysError = sysError{}
)

const WRONG_STRING_SELECTED = "W_STR_S"
const IMPOSSIBLE_RESULT = "IMP_RES"
const SWITCH_WARNING = "S_W"
const DIVISION_BY_ZERO = "N_DBZ0"
const NEGATIVE_NUMBER = "N_NEGN"
const STRING_TOO_LONG = "STR_TL"
const ARRAY_INDEX_OUT_OF_BOUND = "ARR_IOOB"
const ARRAY_ELEMENT_NOT_FOUND = "ARR_ENF"
const ARRAY_WRONG_NUMBER_ELEMENT = "ARR_WNE"
const MAP_INDEX_OUT_OF_BOUND = "MAP_IOB"
const MAP_ELEMENT_NOT_FOUND = "MAP_ENF"

type sysError struct {
	ErrMessage string `json:"message"`
	ErrStatus  string `json:"code"`
	ErrError   string `json:"error"`
}

type SysErrorI interface {
	Message() string
	Status() string
	Error() string
}

func (e sysError) Message() string {
	return e.ErrMessage
}

func (e sysError) Status() string {
	return e.ErrStatus
}

func (e sysError) Error() string {
	return fmt.Sprintf("message: %s - status: %s - error: %s",
		e.ErrMessage, e.ErrStatus, e.ErrError)
}
