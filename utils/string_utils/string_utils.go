package string_utils

import "errors"

//https://golangcode.com/substring-num-of-characters/

func SubString(inputString string, startIndex int, length int) (outputString string, err error) {
	if len(inputString) == 0 {
		return "", nil
	}
	asRune := []rune(inputString)
	endIndex := startIndex + length

	if endIndex > len(inputString) {
		return "", errors.New("Specified substring length would go beyond length of string.")
	}
	return string(asRune[startIndex:endIndex]), nil
}
