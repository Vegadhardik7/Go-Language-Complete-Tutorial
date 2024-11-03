package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

var waitgroup sync.WaitGroup // Define the WaitGroup

func main() {
	// No Concerrency
	// -------------------------------------------

	printAlphanets()
	printNumber()
	fmt.Println("Done")

	s := time.Now()
	for i := 0; i < 10; i++ {
		usr := getUsersFromSite(strconv.Itoa(i))
		fmt.Println(usr)
	}

	e := time.Now()
	difftime0 := e.Sub(s)
	fmt.Printf("Total Time Taken Without Concerrency: %v\n----------\n", difftime0)

	// -------------------------------------------
	// Concerrency
	// -------------------------------------------
	start := time.Now()

	for i := 0; i < 10; i++ {
		waitgroup.Add(1)
		go func() {
			users := getUsersFromSite(strconv.Itoa(i))
			fmt.Println(users)
			waitgroup.Done()
		}()
	}
	waitgroup.Wait()

	end := time.Now()

	// Calculate the total time taken
	diffTime := end.Sub(start)

	// Print the total time taken
	fmt.Printf("Total Time Taken: %v\n", diffTime)

	// -------------------------------------------

	// Run printAlphanets() & printNumber() paraller
	go printAlphanets()
	go printNumber()
	time.Sleep(5 * time.Second)
	fmt.Println("Done")
}

func getUsersFromSite(site string) []string {
	// Simulating 1 second delay to fetch users
	time.Sleep(1 * time.Second)
	return []string{"site" + site, "user1", "user2", "user3"} // Simulated users
}

func printAlphanets() {
	for i := 'a'; i <= 'e'; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Aphabet: %c\n", i)
	}
}

func printNumber() {
	for i := 1; i <= 5; i++ {
		time.Sleep(time.Second)
		fmt.Printf("Number: %d\n", i)
	}
}
