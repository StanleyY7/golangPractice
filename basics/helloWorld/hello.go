// Golang Practice: Hello World

/* What is Golang? Golang (aka Go) is a programming language developed by Google
that was designed to provide a balance between simplicity and performance.
It has characteristics of both high-level programming languages like Python (i.e. simpler/more readable syntax) mixed with
characteristics of low-level programming languages like C++ (pointers) & Java (static typing, efficient execution, more OOP orientated).
*/

package main // similar to Java, groups of related entities (i.e. classes, interfaces, objects) are organised within Packages.

/* Since with Go everything is organised within packages for certain methods like Print, this is also in a package.
Will need to search official docs to find methods from different packages that need to be used. */
import "fmt"

// To run do: go run hello.go (once you are in helloWorld directory)

func main(){ // This is the entry point of a Go application, can only have one main per application. 

	const name = "John" // Like JS can have const or 
	var name2 = "Adam"	// var, var value can change, const value remains constant otherwise gives error

// Similar syntax here to Java

fmt.Println("Hello World!") 
fmt.Println("Hello", name) 
fmt.Printf("Hello %v \n", name2) 

// Similar to JS we can use "syntax sugar" like with JS ()=> arrow functions.

name3 := "Trevor"
fmt.Printf("Hello %v", name3)

}

