package main

import "fmt"

/*func test(x int) {
	var n int
	var result int = 0
	for n = 1; n <= x; n++ {
		result += n
	}
	fmt.Println(result)
}*/
/*func add(xptr *int) {
	*xptr = *xptr + 10
	//fmt.Println(*xptr)

}
*/
type Point struct {
	name string
	age  int
}

func main() {
	/*	var x int = 10

		add(&x)
		fmt.Println(x)

		var msg string
		fmt.Scanln(&msg)
		fmt.Println(msg)
	*/
	var p1 Point = Point{"BILL", 50}
	var p2 Point = Point{age: 25, name: "Brain"}
	fmt.Println(p1.name, p2.name)

}
