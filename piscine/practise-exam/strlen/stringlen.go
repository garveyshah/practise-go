package main 

import "fmt"

func StrLen(s string) int{
	count := 0
	for range s {
		count++
	}
	return count
}
func main() {
	s := ("string hellow world!")
	fmt.Println(StrLen(s))
}