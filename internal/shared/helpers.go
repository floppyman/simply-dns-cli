package shared

func Index(vs []any, t any) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}
