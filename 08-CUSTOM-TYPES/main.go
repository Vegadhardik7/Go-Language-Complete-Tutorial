package main

import "fmt"

// Type definitions
type Celsius float64
type myfunc func(int, int) int

// Type alias declaration
type myStringAlias = string

func main() {
	type hardikint int
	var num hardikint
	num = 10 // type hardikint created from type int
	fmt.Println("num:", num)

	lst := []hardikint{1, 2, 3, 4}
	fmt.Println("Printing the list of values with custom type:", lst)

	temp := Celsius(25)
	fmt.Println("Temperature in Celsius:", temp)
	fmt.Println("Temperature in Fahrenheit:", temp.ToFahrenheit())

	// ---
	fmt.Println("Normal Add:", add(2, 3))
	fmt.Println("calculate0:", calculate0(add))
	fmt.Println("calculate1:", calculate1(add))
	fmt.Println("calculate2:", calculate2(add))

	// --- Type Alias Example ---
	var name myStringAlias = "John"
	var original string = name
	fmt.Printf("Name: %v | Location: %v\n", name, &name)
	fmt.Printf("Original Name: %v | Original Location: %v\n", original, &original)

}

func (c Celsius) ToFahrenheit() float64 {
	return float64(c*9/5 + 32)
}

func (c Celsius) String() string {
	return fmt.Sprintf("%.2f°C", c)
}

func add(a, b int) int {
	return a + b
}

func calculate0(f func(int, int) int) int {
	n1 := 1
	n2 := 2
	return f(n1, n2)
}

func calculate1(f myfunc) int {
	n1 := 5
	n2 := 6
	return f(n1, n2)
}

func calculate2(f myfunc) int {
	n1 := 7
	n2 := 8
	return f(n1, n2)
}
