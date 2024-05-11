package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func CheckError(err error) {
	if err != nil {
		log.Println("Error!! :- ", err)
	}
}

func main() {
	fmt.Println("Please input number")

	reader := bufio.NewReader(os.Stdin)
	num, err := reader.ReadString('\n')
	CheckError(err)

	num = strings.TrimSuffix(num, "\n")

	num1, err := strconv.Atoi(num)
	CheckError(err)

	result := LargestPrimeFact(num1)

	fmt.Println(result)
}

// func LargestPrimeFact(num int) int {

// 	if num <= 1 {
// 		return 1
// 	}

// 	num3 := 0
// 	for i := 3; i <= num; i += 2 {
// 		//num2 := 0
// 		if IsPrime(i) {
// 			if num%i == 0 {
// 				num3 = i
// 			}
// 		}
// 	}
// 	return num3
// }

func IsPrime(n int) bool {
	if n <= 1 {
		return false
	}

	if n <= 3 {
		return true
	}

	if n%2 == 0 || n%3 == 0 {
		return false
	}

	i := 5
	for i*i <= n {
		if n%i == 0 || n%(i+2) == 0 {
			return false
		}
		i += 6
	}

	return true
}


func LargestPrimeFact(num int) int {
    if num <= 1 {
        return 1
    }

    largestPrime := 1
    for i := 2; i*i <= num; i++ {
        for num%i == 0 {
            largestPrime = i
            num /= i
        }
    }
    if num > 1 {
        largestPrime = num
    }

    return largestPrime
}









// func IsPrime2(n int) interface{} {
// 	if n <= 1 {
// 		return "not prime"
// 	}

// 	if n <= 3 {
// 		return n
// 	}

// 	if n%2 == 0 || n%3 == 0 {
// 		return "not prime"
// 	}

// 	i := 5
// 	for i*i <= n {
// 		if n%i == 0 || n%(i+2) == 0 {
// 			return "not prime"
// 		}
// 		i += 6
// 	}

// 	return n
// }
