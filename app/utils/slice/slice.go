package sliceutils

func FindID(val int64, list []int64) int {
	for key, listVal := range list {
		if val == listVal {
			return key
		}
	}

	return -1
}

func Find(val int64, list []int64) bool {
	for _, listVal := range list {
		if val == listVal {
			return true
		}
	}

	return false
}

func remove(s []int, i int) []int {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

func ConvertInt32ArrayToInterfaceArray(intArray []int32) []interface{} {
	interfaceArray := make([]interface{}, len(intArray))
	for i, v := range intArray {
		interfaceArray[i] = v
	}
	return interfaceArray
}
