package general

func IsDigit(c int32) bool {
	return '0' <= c && c <= '9'
}
func isSpace(c int32) bool { //소문자로 시작되는 function은 package 외부에서는 접근할 수 없다.
	switch c {
	case '\t', '\n', '\v', '\f', '\r', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
