package piscine

func Swap(a *int, b *int) {
	n := *a
	m := *b
	*a = m
	*b = n
}
