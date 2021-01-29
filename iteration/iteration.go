package iteration

func Repeat(char string) string {
	var str string
	for i := 0; i < 5; i++ {
		str += char
	}

	return str
}
