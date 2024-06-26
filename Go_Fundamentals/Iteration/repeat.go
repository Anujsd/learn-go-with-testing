package iteration

func Repeat(ch string, cnt int) string {
	var tmp string
	for i := 0; i < cnt; i++ {
		tmp += ch
	}
	return tmp
}
