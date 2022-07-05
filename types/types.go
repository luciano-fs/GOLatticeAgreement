package main

type Elem interface {
	Join(Elem) Elem
	Leq(Elem)  bool
}

type IntSet map[int32]bool

func (a IntSet) Join(b IntSet) IntSet{
	c := make(map[int32]bool)

	for elemA,_ := range a {
		c[elemA] = true
	}
	for elemB,_ := range b {
		c[elemB] = true
	}

	return c
}

func (a IntSet) Leq(b IntSet) bool {
	for elemA,_ := range a {
		if !b[elemA] {
			return false
		}
	}
	return true
}
