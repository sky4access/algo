package array

func reverseString(str string) string {
	if len(str) < 2 {
		return str
	}
	arr := []rune(str)
	res := make([]rune, len(arr))
	j := 0
	for i := len(arr) - 1; i >= 0; i-- {
		res[j] = arr[i]
		j++
	}
	return string(res)
}

func reverseString2(str string) string {
	runes := []rune(str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
