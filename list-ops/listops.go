package listops

type binFunc func(x, y int) int
type predFunc func(int) bool
type unaryFunc func(int) int
type IntList []int

func (il IntList) Append(int IntList) IntList {
	return append(il, int...)
}

func (il IntList) Concat(intLs []IntList) IntList {
	for _, intL := range intLs {
		if intL.Length() == 0 {
			continue
		}
		il = append(il, intL...)
	}
	return il
}

func (il IntList) Filter(fn predFunc) IntList {
	if il.Length() == 0 {
		return []int{}
	}

	var rls IntList
	for _, reg := range il {
		if fn(reg) {
			rls = append(rls, reg)
		}
	}
	return rls
}

func (il IntList) Foldl(fn binFunc, in int) int {
	for _, r := range il {
		in = fn(in, r)
	}
	return in
}

func (il IntList) Foldr(fn binFunc, in int) int {
	for _, r := range il.Reverse() {
		in = fn(r, in)
	}
	return in
}

func (il IntList) Length() int {
	c := 0
	for _, r := range il {
		if r != 0 {
			c++
		}
	}
	return c
}

func (il IntList) Map(fn unaryFunc) IntList {
	if il.Length() == 0 {
		return []int{}
	}

	var rls IntList
	for _, reg := range il {
		rls = append(rls, fn(reg))
	}
	return rls
}

func (il IntList) Reverse() IntList {
	for left, right := 0, len(il)-1; left < right; left, right = left+1, right-1 {
		il[left], il[right] = il[right], il[left]
	}
	return il
}
