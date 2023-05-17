package helper

import (
	"fmt"
	"strings"
)

// To run do: go run . (once you are in booking directory)

// crtl c to exit application

// This is to showcase the package functionality of Go, it is similar to Java (public, private etc..)
// to export a function instead of func functionName it would be func FunctionName that is all. To use packageName.functionName

func Greeting(eventName string, remainingTickets uint){
	fmt.Printf("Welcome to this fake booking sytem, here you can buy tickets for %v \n", eventName)
	fmt.Printf("There are currently %v tickets remaining \n", remainingTickets)
}

func TicketMessage(){
	fmt.Println("Please enter how many tickets you would like to purchase")
}

func PrintFirstNames(userSlice []string) []string {
	// for loop, In Go for replaces all loops including while, for can support all types of loops, while, for, for each, while do
	// _ is a blank identifer it is for a variable you don't need to use 
		// but gives you an error as not in use, so for normally is for index, element but index not used giving error since no need for index
			// can use _ , need to make unused variables explicit with _ in Go
			firstNames := []string{}
			for _, element := range userSlice {
				names := strings.Fields(element) // strings.Fields is equiv to .split in JS however splits by white space only
				firstName := names[0]
				firstNames = append(firstNames, firstName)
			}
			return firstNames
}