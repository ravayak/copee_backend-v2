package syserrors

func DivideByZero(err error) SysErrorI {
	return sysError{
		ErrMessage: "Division by zero",
		ErrStatus:  DIVISION_BY_ZERO,
		ErrError:   err.Error(),
	}
}

func NegativeNumber(err error) SysErrorI {
	return sysError{
		ErrMessage: "negative number",
		ErrStatus:  NEGATIVE_NUMBER,
		ErrError:   err.Error(),
	}
}

func ZeroOrNegativeNumber(err error) SysErrorI {
	return sysError{
		ErrMessage: "zero or negative number",
		ErrStatus:  DIVISION_BY_ZERO + " OR " + NEGATIVE_NUMBER,
		ErrError:   err.Error(),
	}
}

func ImpossibleResult(err error) SysErrorI {
	return sysError{
		ErrMessage: "Impossible result or return value",
		ErrStatus:  IMPOSSIBLE_RESULT,
		ErrError:   err.Error(),
	}
}
