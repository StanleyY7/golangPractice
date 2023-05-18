package main

import (
	"booking/helper"
	"fmt"
)

// To run do: go run . (once you are in booking directory)

// crtl c to exit application

// Instead of Arrays we can use Slices (unqiue to Golang) which is it's own data structure (not method in Golang),
// Slices are dynamic and are the Go equivalent of an ArrayList from Java
// structure is: var sliceName = []sliceType{"value in slice" || leave blank if empty}, we are creating slice of maps here so different

// var userSlice = make([]map[string]string, 0)

var userSlice = make([]userData, 0)

// Struct, similar to interface types in TypeScript, can create custom data structure and define what's in it unlike interface types and
// similar to classes can also have it "do something" (so not just a data but also methods)
// "create a pre defined struct/structure by listing all the properties it should have"

type userData struct {
	firstName string
	lastName string
	ticketsOrdered uint
}

func main(){

	/* Global Variables */

var firstName string
var lastName string
var userTickets uint // uint is positive only numbers, as you can't order negative tickets
var remainingTickets uint = 10
eventName := "Random Event"

// Arrays in Go, fixed size like in Java (so not dynamic), only the same data type can be stored as per the declared/assigned arrayType

userArr := [10]string{} // structure is: var arrName = [size of array]arrayType{"value in array" || leave blank if empty} 

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
	
	// Map, a map is a data type/structure in Go. Consisting of key value pairs (as it maps unique keys to values)
	// this way we can store a full block of data per user
		
	/*
	var userMap = make(map[string]string) // structure to create empty map: var mapName = make(map[dataType of key]dataType of value)

	// we can not mix data types so can't be var mapName = make(map[string]uint), specific to Go
	// adding key value pairs to map
	
	userMap["firstName"] = firstName
	userMap["lastName"] = lastName
	userMap["ticketsOrdered"] = strconv.FormatUint(uint64(userTickets), 10)
			*/

			var userMap = userData {
				firstName: firstName,
				lastName: lastName,
				ticketsOrdered: userTickets,
			}


	userArr[0] = firstName + " " + lastName

	userSlice = append(userSlice, userMap) // this better than arrays as will have to do for each index 
	// whilst with append method for slice don't need to do index

	// Array Properties

	fmt.Printf("------------------------------------------------------------ \n \n")
	fmt.Println("                     Array Properties:                         ")
	fmt.Printf("\n Array values are %v \n", userArr)
	fmt.Printf("First value in array is %v \n", userArr[0])
	fmt.Printf("Array type of Array is %v \n", userArr)
	fmt.Printf("Array length of Array is %v \n \n", len(userArr))
	fmt.Printf("------------------------------------------------------------ \n")

	// Slice Properties, same syntax for Array Properties

		fmt.Printf("------------------------------------------------------------ \n \n")
		fmt.Println("                     Slice Properties:                         ");
		fmt.Printf("\n Slice values are %v \n", userSlice)
		fmt.Printf("First value in slice is %v \n", userSlice[0])
		fmt.Printf("slice type of slice is %v \n", userSlice)
		fmt.Printf("Slice length of slice is %v \n \n", len(userSlice))
		fmt.Printf("------------------------------------------------------------ \n")

		PrintFirstNames()


		fmt.Printf("\n User: %v %v has purchased: %v tickets \n Only %v tickets remaining for %v \n \n", firstName, lastName, userTickets, remainingTickets, eventName)
		fmt.Printf("First names of bookings are %v \n", PrintFirstNames())
	} else {
		fmt.Printf("Error: please make sure your first and last name are longer than 2 characters each \n")
	}
}
} 

func PrintFirstNames() []string {
	// for loop, In Go for replaces all loops including while, for can support all types of loops, while, for, for each, while do
	// _ is a blank identifer it is for a variable you don't need to use 
		// but gives you an error as not in use, so for normally is for index, element but index not used giving error since no need for index
			// can use _ , need to make unused variables explicit with _ in Go
			firstNames := []string{}
			for _, element := range userSlice {

				// names := strings.Fields(element) // strings.Fields is equiv to .split in JS however splits by white space only
				// firstName := names[0]
				
				// with map
				// firstNames = append(firstNames, element["firstName"])

				//with struct
				firstNames = append(firstNames, element.firstName)
			}
			return firstNames
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


