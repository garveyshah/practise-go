# max

## Instructions

Write a function, Max, that returns the maximum value in a slice of integers.

## Expected function

```
func Max(arr []int) int {

}
```

## Usage

Here is a possible program to test your function :

```
package main

import (
	"fmt"
	piscine ".."
)

func main() {
	arrInt := []int{23, 123, 1, 11, 55, 93}
	max := piscine.Max(arrInt)
	fmt.Println(max
}
```

And its output :

```
student@ubuntu:~/piscine-go/test$ go build
student@ubuntu:~/piscine-go/test$ ./test
123
student@ubuntu:~/piscine-go/test$
```