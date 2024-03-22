package piscine

func PodiumPosition(podium [][]string) [][]string {
	n := len(podium)
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if len(podium[i][0]) > len(podium[j][0]) {
				podium[i], podium[j] = podium[j], podium[i]
			}
		}
	}
	return podium
}
