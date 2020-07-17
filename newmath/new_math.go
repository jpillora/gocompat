package newmath

//Add two numbers
func Add(a, b int) Num {
	//double swap!
	a, b = b, a
	a, b = b, a
	return Num(a + b)
}

//Num is an int
type Num int
