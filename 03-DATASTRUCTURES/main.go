// Array | Slice | Map

package main

import (
	"fmt"
	"strings"
)

func main() {

	// ------ Array ------

	var arr [5]int
	fmt.Println("Array:", arr)

	arr[4] = 100
	fmt.Println("Set:", arr)
	fmt.Println("Get:", arr[4])
	fmt.Println("Length of Array:", len(arr))

	arr2 := [6]int{21, 12, 33, 14, 56, 63}
	fmt.Println("Custom Array2:", arr2)

	arr3 := [...]int{5, 3, 2, 43, 54, 565, 76, 787, 345, 345, 23, 41, 241, 423, 6, 476, 57, 68, 76}
	fmt.Println("N Array3:", arr3)

	var arr4 [5]int = [5]int{1, 2, 3}
	fmt.Println("Partial Array:", arr4)

	var arr5 [5]int = [5]int{1: 53, 3: 99}
	fmt.Println("Partial Array With Custom Index:", arr5)

	animals := [...]string{"cat", "dog", "rino", "hippo", "lion"}
	fmt.Println("First 3 animals:", animals[0:3])
	fmt.Println("All animals except 1:", animals[1:])
	fmt.Println("All animals except last:", animals[:len(animals)-1])

	for _, ani := range animals {
		if strings.Contains(ani, "i") {
			fmt.Println("Animal with i", ani)
		} else {
			fmt.Println("Animal with no i", ani)
		}
	}

	for i := 0; i < len(animals); i++ {
		fmt.Println("Animal", i, animals[i])
	}

	n := 0
	for range animals {
		fmt.Println(n, "->", animals[n])
		n++
	}

	// Multi-dimentional Array
	words := [2][3]string{{"hello", "world", "sky"}, {"fish", "cat", "onion"}}
	for i := 0; i < len(words); i++ {
		for j := 0; j < 3; j++ {
			fmt.Println(words[i][j])
		}
	}

	// ------ Slice ------

	s1 := make([]int, 3, 4) // (type, length, capacity)
	fmt.Println("Empty Slice:", s1)

	val := make([]int, 3, 5)
	val[0] = 12
	val[1] = 22
	val[2] = 32
	fmt.Println("Slice1:", val)               // slice value
	fmt.Println("Slice1 Length:", len(val))   // length of the slice
	fmt.Println("Slice1 Capacity:", cap(val)) // capacity of the slice

	val = append(val, 5)
	val = append(val, 6)
	val = append(val, 7, 8, 9)

	fmt.Println("Slice1:", val)               // slice value
	fmt.Println("Slice1 Length:", len(val))   // length of the slice
	fmt.Println("Slice1 Capacity:", cap(val)) // capacity of the slice

	val2 := []int{1, 2, 3}
	fmt.Println("val2:", val2)
	val2 = append(val2, 2, 3, 4)
	fmt.Println("val2 new:", val2)

	var xval []int = make([]int, 0, 5)
	fmt.Println("Empty Slice:", xval)

	for i := 0; i < 10; i++ {
		xval = append(xval, i)
	}
	fmt.Println("final xval:", xval)

	// --------------

	var p []string
	s2 := []string{"aa", "bb", "cc", "dd"}
	fmt.Println("s2:", s2)

	// Create a new slice and assign it to n
	p = s2 // never make a copy of the slice like this
	p[2] = "kk"

	// both of them are some because slice is passed by value and not by reference
	fmt.Println("s2-->", s2)
	fmt.Println("p--->", p)

	p = append(p, "pp")
	fmt.Println("s2-->", s2)
	fmt.Println("p--->", p) // append new value and deteched it from the original array and create new one

	// proper way to make copy of a slice
	slice1 := []int{12, 34, 54, 67}
	slice2 := make([]int, len(slice1))

	copy(slice2, slice1)
	fmt.Println("slice1:", slice1)
	fmt.Println("slice2:", slice2)

	// ------ Maps ------

	marks := map[string]int{"Science": 89, "History": 78, "Maths": 99} // map[key]value
	fmt.Println("marks:", marks)

	rollno := make(map[int]string)

	rollno[101] = "Ram"
	rollno[102] = "Kisan"
	rollno[103] = "Sia"

	// Print the keys
	for key := range rollno {
		fmt.Println(key)
	}

	// Print the values
	for value := range rollno {
		fmt.Println(rollno[value])
	}

	// >>> The above code will give you unordered keys, so to fix that we can use the below one <<<

	index := 0
	for k, v := range rollno {
		fmt.Printf("Index: %v, Key: %v, Value: %v\n", index, k, v)
		index++
	}

	// delete any thing from the map
	rollno[104] = "Jay"
	delete(rollno, 103)

	fmt.Println("Latest Map:", rollno)

	// Issue with map is that it return 0 when no key is present
	r106 := rollno[106]
	fmt.Println(r106)

	// Fix
	newval, present := rollno[106]
	fmt.Println("present", present)
	if present {
		fmt.Println("Roll No:", newval)
	}

	if val, fnd := rollno[106]; fnd {
		fmt.Println("Roll No:", val)
	}

	if val, fnd := rollno[106]; fnd {
		fmt.Println("Roll No:", val)
		delete(rollno, 106)
	}

}
