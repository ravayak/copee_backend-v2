package syserrors

// Main Array errors file

func ArrayIndexOutOfBound(err error) SysErrorI {
	return sysError{
		ErrMessage: "Array Index is out of bound",
		ErrStatus:  STRING_TOO_LONG,
		ErrError:   err.Error(),
	}
}

func ArrayElementNotFound(err error) SysErrorI {
	return sysError{
		ErrMessage: "Array element was not found",
		ErrStatus:  ARRAY_ELEMENT_NOT_FOUND,
		ErrError:   err.Error(),
	}
}

func ArrayWrongNumberElement(err error) SysErrorI {
	return sysError{
		ErrMessage: "There isn't right number of element in array",
		ErrStatus:  ARRAY_WRONG_NUMBER_ELEMENT,
		ErrError:   err.Error(),
	}

}

func SwitchWarning(err error) SysErrorI {
	return sysError{
		ErrMessage: "Switch cannot choose a logic case",
		ErrStatus:  SWITCH_WARNING,
		ErrError:   err.Error(),
	}
}
