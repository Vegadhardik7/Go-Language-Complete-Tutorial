package main

import "fmt"

func main() {
	// pointer
	var pnt *int
	num := 100
	// reference
	pnt = &num
	fmt.Println("Pointer:", pnt)
	fmt.Println("Reference:", *pnt) // dereference a pointer
	fmt.Println("Original num vairable:", num)
	// changing pointer we also change the value where the num variable is refering too
	*pnt = 123
	fmt.Println("Reference:", *pnt) // dereference a pointer
	fmt.Println("Updated num because of pointer update:", num)

	number1 := 100
	fmt.Println("Original Value of Number1:", number1)
	changeValue(&number1)
	fmt.Println("Updated Value of Number1:", number1)

	// Array upgradation required value passed by reference
	var arr = [3]int{1, 2, 3}
	fmt.Println("Original value of arr:", arr)
	changeArrayValue(&arr)
	fmt.Println("Updated value of arr:", arr)

	// Slice upgradation
	slice := []string{"a", "a", "a"}
	fmt.Println("Slice Before:", slice)
	changeSlice(&slice)
	fmt.Println("Slice After:", slice)

	// Map upgradation
	originalMap := make(map[string]int)

	originalMap["Krishna"] = 25
	originalMap["Ram"] = 32

	fmt.Println("Before Map Upgradation:", originalMap)

	modifyMap(originalMap)

	fmt.Println("After Map Upgradation:", originalMap)

}

func changeValue(p *int) {
	*p = 2
}

func changeArrayValue(p *[3]int) {
	p[0] = 7
	p[1] = 8
	p[2] = 9
}

func changeSlice(s *[]string) {
	(*s)[0] = "b"
	(*s)[1] = "b"
	(*s)[2] = "b"

	*s = append(*s, "c")
}

func modifyMap(m map[string]int) {
	m["Ram"] = 30
	m["Sham"] = 20
}
