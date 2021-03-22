package main

/*func test(x int) {
	var n int
	var result int = 0
	for n = 1; n <= x; n++ {
		result += n
	}
	fmt.Println(result)
}*/
func add(n1 int, n2 int) int {
	var result int = n1 + n2
	return result

}
func test() (int, string) {
	return 5, "10"
}
func main() {
	/*var x int = add(5, 3)
	fmt.Println(x * 10)*/
	test()
}
