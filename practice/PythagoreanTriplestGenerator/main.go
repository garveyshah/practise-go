package main

import (
	"fmt"
)

func main() {
	fmt.Println(PythagoreanTriplets(50))
}

func PythagoreanTriplets(limit int) (int, int, int, int) {
	a, b, c := 0, 0, 0
	count := 0
	for i := 1 ; i <= limit; i++ {
		for j := 2; j <= limit; j++ {
			for k := 1 ; k <= limit; k++ {
				if i*i + j*j == k*k {
					a = i
					b = j
					c = k
					count ++
					
				}
			}
		}
	}
	return count, b, a, c

}

// func PythagoreanTriplets(d, e int) (int, int, int) {
// 	a, b, c := 0, 0, 0
// 	//d, e, f := 0, 0, 0

// 	m, n, o := 0, 0, 0

// 	for i := 1 ; i <= d; i++ {
// 		for j := 1; j <= e; j++ {
// 					a = (i*i) - (j*j)
// 					b = (2*i*j)
// 					c = (i*i) + (j*j)
// 					m, n, o = a, b, c
// 			if a*a + b*b == c*c {
// 						a = m
// 						b = n
// 						c = o
// 					}
// 		}
// 	}
// 	return a, b, c
// }
