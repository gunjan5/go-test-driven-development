package pack

//Add adds two numbers
func Add(n ...int) int {
	var sum int

	if len(n) == 0 {
		print("no args passed")
		return 0
	}
	for _, i := range n {
		sum += i
	}
	return sum
}
