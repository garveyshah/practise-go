package main

import "fmt"

func main() {
	str1 := "abdnej vdjc andvenflqf"
	str2 := "dsvvnsjafhjefncllzsjcadafanfdlaf"
	str3 := "12 38765"
	str4 := "1278"

	fmt.Println(WeAreUnique(str1, str2))
	fmt.Println(WeAreUnique(str3, str4))

}

func WeAreUnique(str1, str2 string) int{
	mapCount := make(map[rune]int)

	for _, chr := range str1 {
		mapCount[chr]++
	}
	for _, char := range str2 {
		mapCount[char]++
	}
	unique := 0
	dups := []rune{}
	for val, char := range mapCount{
		if char == 1 {
			unique++
			dups = append(dups, val)
		}
	}
	fmt.Println(string(dups))
	return unique
}