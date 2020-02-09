package main

import (
	"errors"
	"fmt"
	"math"
)

func main() {
	//oh hai
	fmt.Println("Hello, World!")

	//error handling
	result, err := sqrt(-16)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	//structs
	p := person{name: "Robbie", age: 28}
	fmt.Println(p)

	/*--------
	 POINTERS
	---------*/

	//initiates an integer, i=7
	//passes the memory address of i's value to inc
	i := 7
	inc(&i)
	fmt.Println(i)
}

func sum(x int, y int) int {
	return x + y
}

//sqrt outputs result and error,
//set result set to 0 if error occurs
//error set to nil if sqrt works
func sqrt(x float64) (float64, error) {
	if x < 0 {
		return 0, errors.New("Undefined for negative numebrs")
	}

	return math.Sqrt(x), nil
}

//pretty structs
type person struct {
	name string
	age  int
}

//increment function that accepts a pointer to an int
func inc(x *int) {
	//de-references the pointer
	*x++
}
