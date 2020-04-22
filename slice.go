package gox

func IndexOfInt(a []int, i int) int {
	for idx, v := range a {
		if v == i {
			return idx
		}
	}

	return -1
}

func IndexOfInt64(a []int64, i int64) int {
	for idx, v := range a {
		if v == i {
			return idx
		}
	}

	return -1
}

func IndexOfString(a []string, s string) int {
	for i, str := range a {
		if str == s {
			return i
		}
	}

	return -1
}

func ReverseIntSlice(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func ReverseInt64Slice(s []int64) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}

func RemoveInt64(s []int64, v int64) []int64 {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func RemoveFloat64(s []float64, v float64) []float64 {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}

func RemoveString(s []string, v string) []string {
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == v {
			s = append(s[:i], s[i+1:]...)
		}
	}
	return s
}
