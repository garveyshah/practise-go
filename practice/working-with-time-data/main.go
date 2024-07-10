package main

import (
	"fmt"
	"time"
)

func main() {
	// Tell the time
	t := time.Now()
	fmt.Println(t)
	fmt.Println()
	// output : 2024-07-10 20:58:52.36530469 +0300 EAT m=+0.000225517
	/*
			 2024-07-10: Year, month, day.

		    20:58:52.36530469 : Hour, minutes, seconds

		   +0300 : Time difference from GMT.

		    EAT: The current time zone you are in.

		   m=+0.000225517: Monotonic clock reading.

	*/

	// Better way to format the ouput.
	//T := time.Now()
	// fmt.Println(T.Year())
	// fmt.Println(T.Month())
	// fmt.Println(T.Day())
	// fmt.Println(T.Date())
	// fmt.Println(T.Hour())
	// fmt.Println(T.Minute())
	// fmt.Println(T.Second())

	// output :
	/* 	2024
		July
		10
		2024 July 10
		21
		7
		14
	*/
	// How can we print this in a prettier format?
	//fmt.Printf("%d %d %d\n", T.Year(), T.Month(), T.Day())
	// Output: 2024 7 10

	// There is a better way to to format time, using the time.Format function.
	// fmt.Println(T.Format("Mon jan 2 15:04:05 2006 MST"))
	// fmt.Println(T.Format("Mon Jan 2 15:04:05"))
	// fmt.Println(T.Format("2006/01/02"))
	// fmt.Println(T.Format("3:01PM"))
	// fmt.Println(T.Format("15:04PM"))

	// OUTPUT :
	
	// Wed jan 10 21:24:04 2024 EAT
	// Wed Jul 10 21:24:04
	// 2024/07/10
	// 9:07PM
	// 21:24PM

	// You could also use a predefined formart as well, like so:
	// fmt.Println(time.UnixDate)
	// fmt.Println(time.RFC3339)

	//Different Time Zones
	nt := time.Now()
	// lt := time.Now().Local()
	// ut := time.Now().UTC()
	// fmt.Println(nt)
	// fmt.Println(lt)
	// fmt.Println(ut)

	// Local() gets the local time Zone, which is same as what time.Now() would automatically detect.
	// Calling UTC() will convert the time zone to UTC.

	i, _ := time.LoadLocation("UTC")
	fmt.Printf("%s\n", nt.In(i))

	i, _ = time.LoadLocation("Europe/London")
	fmt.Printf("%s\n", nt.In(i))

	i, _ =time.LoadLocation("America/Los_Angeles")
	fmt.Printf("%s\n", nt.In(i))

	i, _ = time.LoadLocation("Asia/Seoul")
	fmt.Printf("%s\n",nt.In(i))

	// 	OUTPUT:
	// 	2024-07-10 20:06:53.630249738 +0000 UTC
	// 2024-07-10 21:06:53.630249738 +0100 BST
	// 2024-07-10 13:06:53.630249738 -0700 PDT
	// 2024-07-11 05:06:53.630249738 +0900 KST

	// TIME.loadLocation will load a locale of your choice. You can use this result to cobert time by passing in to 'yourTime.In'.
	

}
