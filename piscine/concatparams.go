package piscine

func ConcatParams(arg []string) string {
	result := ""
	for i := 0; i < len(arg); i++ {
		result += arg[i]
		if i < len(arg)-1 {
			result += "\n"
		}
	}
	return result
}
