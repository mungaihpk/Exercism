package strain

type Ints []int
type Lists [][]int
type Strings []string

func (i Ints) Keep(filter func(int) bool) Ints {
	//panic("Please implement the Keep function")
	var integers Ints
	for _, num := range i {
		if filter(num) {
			integers = append(integers, num)
		}
	}
	return integers
}

func (i Ints) Discard(filter func(int) bool) Ints {
	//panic("Please implement the Discard function")
	var integers Ints
	for _, num := range i {
		if !filter(num) {
			integers = append(integers, num)
		}
	}
	return integers
}

func (l Lists) Keep(filter func([]int) bool) Lists {
	//panic("Please implement the Keep function")
	var lists Lists
	for _, list := range l {
		if filter(list) {
			lists = append(lists, list)
		}
	}
	return lists
}

func (s Strings) Keep(filter func(string) bool) Strings {
	//panic("Please implement the Keep function")
	var strings Strings
	for _, str := range s {
		if filter(str) {
			strings = append(strings, str)
		}
	}
	return strings
}
