package operations

func Join(a, b map[int32]bool) map[int32]bool {
	c := make(map[int32]bool)

	for elemA,_ := range a {
		c[elemA] = true
	}
	for elemB,_ := range b {
		c[elemB] = true
	}

	return c
}

func Leq(a, b map[int32]bool) bool {
	for elemA,_ := range a {
		if !b[elemA] {
			return false
		}
	}
	return true
}
