package project01

func UltimateDivMod(a *int, b *int) {
	A := *a
	*a = *a / *b
	*b = A % *b
}
