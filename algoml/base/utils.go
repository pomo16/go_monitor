package base

//IsDigitByte 字节是否为数字
func IsDigitByte(b byte) bool {
	return '0' <= b && b <= '9'
}

//IsLikeDigit 判定是否像数字字符串组合特征，时间，日期，浮点数等等
func IsLikeDigit(b byte) bool {
	return ('0' <= b && b <= '9') || b == '.' || b == '-' || b == ':'
}

//IsDigitForm 判定分词的前两个字符是不是满足普通数字字符串特征
func IsDigitForm(s string) bool {
	if len(s) == 1 && IsDigitByte(s[0]) {
		return true
	}

	if len(s) > 1 && IsDigitByte(s[0]) && IsLikeDigit(s[1]) {
		return true
	}

	return false
}

//IsEngChar 检查是否为英文字符
func IsEngChar(b byte) bool {
	return ('A' <= b && b <= 'Z') || ('a' <= b && b <= 'z')
}
