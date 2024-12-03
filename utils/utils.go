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

type Inty interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Numerical interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~float32 | ~float64
}

func ParseInt[T Inty](s string) T {
	var result T
	for i := 0; i < len(s); i++ {
		ch := s[i]
		digit := T(ch - '0')
		result = result*10 + digit
	}
	return result
}
