/*
	God program to illustrate how to create an arrya using shorthand

declaration and accessing the elements of the array using for loop
*/
package main

import "fmt"

func main() {
	//  Using var keyword to declare array
	var months [3]string
	months[0] = "April"
	months[1] = "May"
	months[2] = "June"

	var salesByMonth [3]float64
	salesByMonth[0] = 1710.26
	salesByMonth[1] = 2245.97
	salesByMonth[2] = 3032.45

	fmt.Println(months[0], salesByMonth[0])
	fmt.Println(months[1], salesByMonth[1])
	fmt.Println(months[2], salesByMonth[2])
	fmt.Println()

	// Shorthand declaration of array
	arr := [4]string{"geek", "gfg", "Geeks1231", "GeeksforGeeks"}

	// Accessing the elements of the array Using for loop
	fmt.Println("Elements of the array:")

	for i := 0; i <= 3; i++ {
		fmt.Println(arr[i])
	}
	fmt.Println()

	// Illustration of the concept of multi-dimension array
	/* Creaating and initializing 2-dimentional array Using shorthand declaration. Here the (,) Comma is necessary*/
	arr1 := [3][3]string{{"c #", "C", "Python"}, {"Java", "Scala", "Perl"}, {"C++", "Go", "HTML"}}

	// Accessing the values of the array Using for loop
	fmt.Println("Elements of Array 1")
	for x := 0; x < 3; x++ {
		for y := 0; y < 3; y++ {
			fmt.Println(arr1[x][y])
		}
	}
	fmt.Println()

	/* Creating a 2-dimensional array using the var keyword and initialing a multi-dimential arrau using index*/
	var arr2 [2][2]int
	arr2[0][0] = 100
	arr2[0][1] = 200
	arr2[1][0] = 300
	arr2[1][1] = 400

	// Accessing the values of the array
	fmt.Println("Elements of array 2")
	for p := 0; p < 2; p++ {
		for q := 0; q < 2; q++ {
			fmt.Println(arr2[p][q])
		}
	}
	fmt.Println()

	/* Creating an array of int type which store, two elements
	Here, we do not initalize the array so the value of the array is zero*/
	var myarr[2]int
 fmt.Println("Elements of the Array: ", myarr)
 fmt.Println()

 // Illustration of the concept of eliipsis in an array
 /*Creating an array whpse size is determined by the number of elements present in it Using ellipsis*/
 myarray := [...]string{"GFG", "gfg", "geeks", "GeekforGreeks", "GEEK"}

 fmt.Println("Elements of the array: ", myarray)

 /*Length of the array is determined using len() method */
 fmt.Println("Length of the array is:", len(myarray))
 fmt.Println()

 // Illustration of value type of array

 /* Creating an arrray whose size is represented by rhe ellipsis*/
 my_array := [...]int{100, 200, 300, 400, 500}
 fmt.Println("Original array(Before):", my_array)

 // Creating a variable and initializing qith my_array
 new_array := my_array

 fmt.Println("New array(before)",new_array)

 // Change the value at index 0 to 500
new_array[0] = 500

fmt.Println("New array(after):", new_array)

fmt.Println("Original array(After):", my_array)
fmt.Println()

// Arrays
arr3 := [3]int{9,7,6}
arr4 := [...]int{9,7,6}
arr5 := [3]int{9,5,3}

// Comparing arrays using == operator
fmt.Println(arr3 == arr4)
fmt.Println(arr4 == arr5)
fmt.Println(arr3 == arr5)

/*
// this will give an error because the type of arr3 and arr6 is a mismatch
arr6 := [4]int{9,7,6}
fmt.Println(arr3== arr6)
*/
}