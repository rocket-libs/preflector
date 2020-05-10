package strings
//https://golangcode.com/substring-num-of-characters/
SubString(inputString string, startIndex int, length int) (outputString string,err error){
	if(len(inputString) == 0){
		return "",nil
	}
	asRune := []rune(myString)
	minSelectionLength = startIndex + length
	if(minSelectionLength > len(inputString)){
		return "", err
	}
}