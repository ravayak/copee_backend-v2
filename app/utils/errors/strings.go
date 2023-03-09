package syserrors

func StringIsTooLong(err error) SysErrorI {
	return sysError{
		ErrMessage: "String is too long",
		ErrStatus:  STRING_TOO_LONG,
		ErrError:   err.Error(),
	}
}

func WrongStringSelected(err error) SysErrorI {
	return sysError{
		ErrMessage: "Wrong string selected",
		ErrStatus:  WRONG_STRING_SELECTED,
		ErrError:   err.Error(),
	}
}
