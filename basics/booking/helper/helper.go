package helper

import (
	"fmt"
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

