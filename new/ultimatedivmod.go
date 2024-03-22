package piscine

func UltimateDivMod(a *int, b *int) {
	tempA := *a
	*a = *a / *b
	*b = tempA % *b
}
