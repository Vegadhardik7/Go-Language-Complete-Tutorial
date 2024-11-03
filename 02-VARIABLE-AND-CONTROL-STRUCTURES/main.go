package main

import "fmt"

var globalVar = "I am global"

func main() {

	var name string
	var age int
	var location string

	var address, company string

	var (
		city, country string
		weight        float32
	)

	name = "Hardik"
	age = 25
	location = "abc"
	address = "104-xoxo"
	company = "Cybage"
	city = "Pune"
	country = "INDIA"
	weight = 73.2
	var currency = "INR" // type inference
	car := "BMW"         // short variable declaration

	fmt.Println("Data:", name, address, age, location, city, company, country, weight, currency, car)

	fmt.Printf("My name is %v \nmy age is %v \nmy current location is %v \nmy current city is %v \nI work in %v \nMy country is %v ", name, age, location, city, company, country)

	// Zero value of variables
	var (
		i int
		f float32
		b bool
		s string
	)

	fmt.Printf("%v %v %v %q\n", i, f, b, s) // %q : prints value in ""

	getGlobal()
	fmt.Print(globalVar)

	// for loop

	x := 0
	for i := 0; i < 10; i++ {
		x += i
		fmt.Printf("x%v:%v\n", i, x)
	}

	for x < 10 {
		x += 1
	}
	fmt.Println("x-->", x)

	// continue and break & if condition

	fmt.Println("-------------")
	val := 1
	for {
		val += val

		if val%2 != 0 {
			continue
		}
		fmt.Println(val)
		if val > 100 {
			break
		}
	}

	fmt.Println("---------")
	x = 1
	for x < 10 {
		x += 1
		if x%2 == 0 {
			fmt.Println("EVEN:", x)
		} else if x%3 == 0 {
			fmt.Println("ODD:", x)
		} else {
			fmt.Println("NONE:", x)
		}
	}

	// Switch Case

	i = 2
	fmt.Print("Write ", i, " as ")
	switch i {
	case 0, 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	case 4:
		fmt.Println("Four")
	}

	x = 55
	fmt.Println("Result ", x)
	switch {
	case x > 50:
		fmt.Println("PASS")
	case x < 50:
		fmt.Println("FAIL")
	default:
		fmt.Println("EQUAL")
	}
}
