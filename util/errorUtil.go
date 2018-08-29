package util

/**
	internetUtil.go调用
 */

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

type decodeError struct {
	info string
}

func (e *decodeError) Error() string {
	return e.info
}
