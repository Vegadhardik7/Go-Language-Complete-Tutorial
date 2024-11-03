package main

import (
	"fmt"

	"github.com/pkg/errors"
)

func main() {
	a, b := 5, 0
	sum := add(a, b)
	fmt.Printf("Additon of %v + %v = %v\n", a, b, sum)

	n1, n2, n3 := returnThreeNumbers()

	fmt.Println("Return 3 Numbers:", n1, n2, n3)

	fmt.Printf("Multiplication of %v * %v = %v\n", a, b, multiply(a, b))

	// Error Handling in Go
	res, err := divide(a, b)
	if err != nil {
		fmt.Printf("result : %v\n", err)
	}
	fmt.Printf("Division of %v / %v = %v\n", a, b, res)

	res, err = divide2(a, b)
	if err != nil {
		fmt.Printf("result : %v\n", err)
	}
	fmt.Printf("Division of %v / %v = %v\n", a, b, res)

	fmt.Println("Sum of all the numbers non variadic is", addAllNumbers([]int{1, 2, 3, 4, 4, 5, 6}))

	fmt.Println("Sum of all the even numbers is", addAllEvenNumbers([]int{1, 2, 3, 4, 4, 5, 6}, isEven))

	// --- variadic function ---
	val := addAll(5, 6, 7, 8)
	fmt.Printf("Sum of all the numbers is %v\n", val)

	// --- anonimus function ---
	var anonimusFun func(int) int
	anonimusFun = func(x int) int {
		return x + 1
	}
	fmt.Println("Anonimus Function:", anonimusFun(5))

	// --- anonimus function direct call (Function literal) ---
	anonimusFun2 := func(x int) int {
		return x + 1
	}(7)
	fmt.Println("Anonimus2 Function:", anonimusFun2)

	// --- closure function ---
	closurenum := 5
	func() {
		fmt.Println("closure num =", closurenum)
	}()

	// --- callback function call ---
	fmt.Printf("%v is even %v\n", a, isEven(a))
	fmt.Printf("%v is odd %v\n", a, isOdd(a))

	// --- Create a function that multiplies by 2 ---
	double := createMultiplier(2)

	// --- Create a function that multiplies by 3 ---
	triple := createMultiplier(5)

	fmt.Println("Double of 5:", double(5))
	fmt.Println("Triple of 5:", triple(5))

	// --- calling closure function ---
	counter := newCounter()
	fmt.Printf("Counter call 1st time %v\n", counter())
	fmt.Printf("Counter call 2nd time %v\n", counter())
	fmt.Printf("Counter call 3rd time %v\n", counter())
	fmt.Printf("Counter call 4th time %v\n", counter())
	fmt.Printf("Counter call 5th time %v\n", counter())

	// --- panic example ---
	accessEle := accessElements([]int{1, 2, 4, 5, 6}, 8)
	fmt.Println("Access Element:", accessEle)

}

func add(a int, b int) int {
	return a + b
}

func returnThreeNumbers() (int, int, int) {
	a, b, c := 10, 20, 30
	return a, b, c
}

// Named return variable
func multiply(a int, b int) (mul int) {
	mul = a * b
	return
}

func divide(a, b int) (int, error) {
	var result int
	if b == 0 {
		return result, errors.New("zero division error, denominator can't be 0")
	}
	result = a / b
	return result, nil
}

// Other way to write above divide function
func divide2(a, b int) (res int, err error) {
	if b == 0 {
		return
	}
	res = a / b
	return
}

func addAllNumbers(num []int) int {
	sum := 0
	for _, n := range num {
		sum += n
	}
	return sum
}

func addAllEvenNumbers(num []int, evenfilter func(int) bool) int {
	sum := 0
	for _, n := range num {
		if evenfilter(n) {
			sum += n
		}
	}
	return sum
}

// Variadic functions
func addAll(nums ...int) int {
	res := 0
	for _, n := range nums {
		res += n
	}
	return res
}

// --- Callback function useage ---
func isEven(x int) bool {
	return x%2 == 0
}

func isOdd(x int) bool {
	return x%2 != 0
}

// createMultiplier take int as a value and return function which takes int as a value and return int
// example: 1st int 2 bcz we want to double, taken by createMultiplier...
// make it double func(5)... which return 5*2 = 10 i.e int
func createMultiplier(factor int) func(int) int {
	return func(i int) int {
		return i * factor
	}
}

// closure function example
func newCounter() func() int {
	cnt := 0
	return func() int {
		cnt += 1
		return cnt
	}
}

// Use of panic function
func accessElements(values []int, index int) int {
	if index >= len(values) {
		panic("*** index cannot be greater than length of values ***")
	}
	return values[index]
}
