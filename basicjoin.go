package piscine

func BasicJoin(elems []string) string {
	str := ""
	for _, y := range elems {
		str = str + y
	}
	return str
}
