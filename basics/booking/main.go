package main

import (
	"booking/helper"
	"fmt"
)

// To run do: go run . (once you are in booking directory)

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
		break
	}

// Greeting/Intro

helper.Greeting(eventName, remainingTickets)

// Ask user for their name

	fmt.Println("Please enter your First Name:")
	fmt.Scan(&firstName) // & is for memory address so that it will stop instead of running without asking for user input as it references to memory instead of passing through value
	
	fmt.Println("Please enter your Last Name:")
	fmt.Scan(&lastName)

	// Input Validation

	var nameValid = len(firstName) >=2 && len(lastName) >= 2

// Ask for amount of tickets to purchase

if (!valid){

	helper.TicketMessage()
	fmt.Scan(&userTickets)

	// Input Validation

		if (userTickets == 0){
			for (userTickets == 0){
				fmt.Println("Error can only purchase tickets greater than amount of 0")
				helper.TicketMessage()
				fmt.Scan(&userTickets)
			}  // for is Go syntax for while loop
		} else if (userTickets > remainingTickets){
			for (userTickets > remainingTickets){
			fmt.Printf("Error can only purchase tickets less than %v amount and greater than 0 \n", remainingTickets)
			helper.TicketMessage()
			fmt.Scan(&userTickets)
			}  
		}

}

	valid = true

	if (nameValid){
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

		helper.PrintFirstNames(userSlice)


		fmt.Printf("\n User: %v %v has purchased: %v tickets \n Only %v tickets remaining for %v \n \n", firstName, lastName, userTickets, remainingTickets, eventName)
		fmt.Printf("First names of bookings are %v \n", helper.PrintFirstNames(userSlice))
	} else {
		fmt.Printf("Error: please make sure your first and last name are longer than 2 characters each \n")
	}
}
} 

// ----------------------------------------------- Further Notes ------------------------------------------------------------------------- //


// In Go a function can return multiple values

/* If we were to do a more generic type of input validation

func validate(firstName string, lastName string, valid bool) (bool, bool) {

	nameValid := firstName >=2 && lastName >=2
	isValid := valid == true

	return nameValid, isValid
}

// In use:

nameValid, isValid := validate(firstName, lastName, valid)

*/

