package common

func Trenary(statement bool, a, b interface{}) interface{} {
	if statement {
		return a
	}

	return b
}
