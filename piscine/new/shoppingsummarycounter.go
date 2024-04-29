package piscine

func ShoppingSummaryCounter(str string) map[string]int {
	list := make(map[string]int)
	start := 0
	end := 0
	for i := 0; i < len(str); i++ {
		if str[i] == ' ' {
			word := str[start:end]
			list[word]++
			start = i + i
		}
		end++
	}
	word := str[start:end]
	list[word]++
	return list
}
