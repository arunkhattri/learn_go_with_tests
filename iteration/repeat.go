package iteration

func Repeat(str1 string, n int) (result string) {
	for i := 0; i < n; i++ {
		result += str1
	}
	return result
}
