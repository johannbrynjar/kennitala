package utils

import "strconv"

func StringToInt(input string) (int8, error) {
	intVar, err := strconv.ParseInt(input, 10, 8)
	if err != nil {
		return 0, err
	}
	return int8(intVar), nil
}
