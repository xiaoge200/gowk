package gowk

import "cmp"

// 转指针
func Ptr[T cmp.Ordered | bool](v T) *T {
	return &v
}

func IntToBool(i int) bool {
	return i != 0
}

func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}
