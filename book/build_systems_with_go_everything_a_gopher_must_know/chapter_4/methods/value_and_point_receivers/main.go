/*
	Receivers are very important because they define the type "receiving" the logic inside the method.
	In the example below we define two methods with the same logix.
	Metho `Enlarge` receives a value of type `Rectangle` and method `EnlargeP` receives a type `*Rectangle`.
	Observe how `Enlarge` does not modify any field of the original variable, while `EnlargeP` does.
	This happens becouse the `EnlargeP` receives the pointer to `rect` whereas, `Enlarge` receivers a copy.

	In the example below, the `EnlargeP` method requires a pointer. However, we invoke the method with rect.ElargeP(2)
	and `rect` is not a pointer. This is possible because the Go interpreter translates this into `(&rect).EnlargeP(2).`

	If a method can have value or pointer receivers, which one should you use?
	Generally, using pointers is more efficient because it reduces the number of copies. However, in some situations you
	may be more comfortable with value receivers. In any case, you should be consistent and do not mix them.
*/

package main

import "fmt"

type Rectangle struct {
	Height int
	Width  int
}

func (r Rectangle) Enlarge(factor int) {
	r.Height = r.Height * factor
	r.Width = r.Width * factor
}

func (r *Rectangle) EnlargeP(factor int) {
	r.Height = r.Height * factor
	r.Width = r.Width * factor
}

func main() {
	rect := Rectangle{2, 2}
	fmt.Println(rect)

	rect.Enlarge(2)
	fmt.Println(rect)

	rect.EnlargeP(2)
	fmt.Println(rect)
}
