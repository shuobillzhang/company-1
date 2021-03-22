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
func add(n1 int, n2 int) {
	var result int = n1 + n2
	fmt.Println(result)
}
func main() {

	add(3, 5)

}
