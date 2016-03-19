package pack

//Add adds two numbers
func Add(n ...int) int {
	var sum int
	for _, i := range n {
		sum += i
	}
	return sum
}
