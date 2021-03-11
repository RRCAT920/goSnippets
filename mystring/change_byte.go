package mystring

// 修改字符串中的字符
func Set(s *string, index int, char byte) {
	temp := []byte(*s)
	temp[index] = char
	*s = string(temp)
}
