package helpers

import "strconv"

func HandleError(err error) {
	if err != nil {
		panic(err)
	}
}

func StringToInt(str string) int {
	num, err := strconv.Atoi(str)
	HandleError(err)
	return num
}
