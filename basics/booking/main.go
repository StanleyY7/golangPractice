package main

import (
	"fmt"
	"os"
	"strings"
)

// To run do: go run main.go (once you are in booking directory)

// crtl c to exit application

func main(){

	/* Global Variables */

var firstName string
var lastName string
var userTickets uint // uint is positive only numbers, as you can't order negative tickets
var remainingTickets uint = 10
eventName := "Random Event"


// Arrays in Go, fixed size like in Java (so not dynamic), only the same data type can be stored as per the declared/assigned arrayType

userArr := [10]string{} // structure is: var arrName = [size of array]arrayType{"value in array" || leave blank if empty} 

// Instead of Arrays we can use Slices (unqiue to Golang) which is it's own data structure (not method in Golang), 
// Slices are dynamic and are the Go equivalent of an ArrayList from Java

userSlice := []string{} // structure is: var sliceName = []sliceType{"value in slice" || leave blank if empty}

for {
	
	var valid bool = false

	if (remainingTickets == 0){
		fmt.Print("No more tickets to purchase, application exited")
		os.Exit(0)
	}

// Greeting/Intro

fmt.Printf("Welcome to this fake booking sytem, here you can buy tickets for %v \n", eventName)
fmt.Printf("There are currently %v tickets remaining \n", remainingTickets)

// ask user for their name

fmt.Println("Please enter your First Name:")
fmt.Scan(&firstName) // & is for memory address so that it will stop instead of running without asking for user input as it references to memory instead of passing through value

fmt.Println("Please enter your Last Name:")
fmt.Scan(&lastName)

// ask for amount of tickets to purchase

if (!valid){

	fmt.Println("Please enter how many tickets you would like to purchase")
	fmt.Scan(&userTickets)

	for (userTickets > remainingTickets){
		fmt.Printf("Error can only purchase tickets less than %v amount \n", remainingTickets)
		fmt.Println("Please enter how many tickets you would like to purchase")
		fmt.Scan(&userTickets)
	} // for is Go syntax for while loop

}

	valid = true

	remainingTickets = remainingTickets - userTickets
	userArr[0] = firstName + " " + lastName
	userSlice = append(userSlice, firstName + " " + lastName) // this better than arrays as will have to do for each index 
	// whilst with append method for slice don't need to do index

	// Array Properties

	fmt.Printf("------------------------------------------------------------ \n \n")
	fmt.Println("                     Array Properties:                         ")
	fmt.Printf("\n Array values are %v \n", userArr)
	fmt.Printf("First value in array is %v \n", userArr[0])
	fmt.Printf("Array type of Array is %v \n", userArr)
	fmt.Printf("Array length of Array is %v \n \n", len(userArr))
	fmt.Printf("------------------------------------------------------------ \n")

	// Slice Properties, same syntax for Arrapy Properties

		fmt.Printf("------------------------------------------------------------ \n \n")
		fmt.Println("                     Slice Properties:                         ");
		fmt.Printf("\n Slice values are %v \n", userSlice)
		fmt.Printf("First value in slice is %v \n", userSlice[0])
		fmt.Printf("slice type of slice is %v \n", userSlice)
		fmt.Printf("Slice length of slice is %v \n \n", len(userSlice))
		fmt.Printf("------------------------------------------------------------ \n")

	firstNames := []string{}

	// for loop, In Go for replaces all loops including while, for can support all types of loops, while, for, for each, while do
		// _ is a blank identifer it is for a variable you don't need to use 
			// but gives you an error as not in use, so for normally is for index, element but index not used giving error since no need for index
				// can use _ , need to make unused variables explicit with _ in Go

		for _, element := range userSlice {
			names := strings.Fields(element) // strings.Fields is equiv to .split in JS however splits by white space only
			firstName := names[0]
			firstNames = append(firstNames, firstName)
		}


		fmt.Printf("\n User: %v %v has purchased: %v tickets \n Only %v tickets remaining for %v \n \n", firstName, lastName, userTickets, remainingTickets, eventName)
		fmt.Printf("First names of bookings are %v \n", firstNames)
	}

}