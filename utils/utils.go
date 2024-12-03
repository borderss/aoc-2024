package utils

func PtrBool(val bool) *bool { return &val }

func ParseInt8(s string) int8 {
	var result int8
	for i := 0; i < len(s); i++ {
		ch := s[i]
		digit := int8(ch - '0')
		result = result*10 + digit
	}
	return result
}
