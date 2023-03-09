package syserrors

func MapIndexOutOfBound(err error) SysErrorI {
	return sysError{
		ErrMessage: "Map Index is out of bound",
		ErrStatus:  MAP_INDEX_OUT_OF_BOUND,
		ErrError:   err.Error(),
	}
}

func MapElementNotFound(err error) SysErrorI {
	return sysError{
		ErrMessage: "Map element was not found",
		ErrStatus:  MAP_ELEMENT_NOT_FOUND,
		ErrError:   err.Error(),
	}
}
