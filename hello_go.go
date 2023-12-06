package main

import (
	"container/list"
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Hello Go")
	fmt.Println("Show me what you got!!") // An inline comment
	/*
		Block
		commets
	*/

	// Variables
	var myString string = "I am a static type string variable"
	fmt.Println(myString) // Output: "I am a static type string variable"

	myString = "58"
	fmt.Println(myString) // Output: 58 (as string)

	myString = "I can change but I need to be a string"
	fmt.Println(myString) //I can change but I need to be a string

	var myInt int = 7
	myInt = myInt - 4
	fmt.Println(myInt - 1) // Output: 2

	fmt.Printf("%s %d", myString, myInt) // Output: I can change but I need to be a string 3

	fmt.Println(reflect.TypeOf(myInt)) // Output: int

	var myFloat float64 = 6.5
	fmt.Println(myFloat)                  // Output: 6.5
	fmt.Println(reflect.TypeOf(myFloat))  // Output: float64
	fmt.Println(myFloat + float64(myInt)) // Output: 9.5

	var myBool bool = false
	myBool = true
	fmt.Println(myBool) // Output: true

	/*
		Another way to declare variables where the type is inferred
		from the value assigned to it => :=
		Example:
	*/

	stringWithInferredType := "I'll be a string"
	fmt.Println(stringWithInferredType)

	// Constants:

	const myConst = "I am a constant"

	// Conditionals:

	myCash := 4
	ticketPrice := 5
	haveDiscount := true
	discountPrice := 2

	if myCash >= ticketPrice {
		fmt.Println("I can watch the show")
	} else if haveDiscount == true || myCash >= discountPrice {
		fmt.Println("I can watch the show using the discount")
	} else {
		fmt.Println("I'll go home")
	}

	// Array

	var myArray [3]int   // Create an array with 3 positions initialized in 0
	fmt.Println(myArray) // Output: [0 0 0]
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray) // Output: [1 2 3]

	// Map

	myMap := make(map[string]int)
	myMap["Rick"] = 70
	myMap["Morty"] = 14
	myMap["Squanchy"] = 100

	fmt.Println(myMap)          // Output: map[Morty:14 Rick:70 Squanchy:100]
	fmt.Println(myMap["Morty"]) // Output: 14

	myMap2 := map[string]int{"Beth": 45, "Jerry": 45, "Summer": 17} // Declare with keys-values
	fmt.Println(myMap2)                                             // Output: map[Beth:45 Jerry:45 Summer:17]

	// List

	myList := list.New()
	myList.PushBack(1) // We add to the back
	myList.PushBack(1)
	myList.PushBack(3)

	fmt.Println(myList.Front().Value) // Output: 1

	// Loop

	scifiFilms := [3]string{"Blade Runner", "The Matrix", "Back to the Future"} // or => var scifiFilms [3]string = [3]string{"Blade Runner", "The Matrix", "Back to the Future"}

	fmt.Println(scifiFilms) // Output: [Blade Runner The Matrix Back to the Future]

	for i := 0; i < len(scifiFilms); i++ {
		println(scifiFilms[i])
	}

	for key, value := range myMap2 {
		fmt.Println(key, value) // Output: Beth 45 Jerry 45 Summer 17
	}
	for index := range scifiFilms {
		fmt.Println(scifiFilms[index])
	}
	for index, value := range scifiFilms {
		fmt.Println(index, value)
	}

	// Function

	result := myFunction("Result")
	println(result)

	// Estructura

	type MyStruct struct {
		name  string
		power int
	}

	myStruct := MyStruct{"Rick", 100}
	fmt.Println(myStruct)      // Output: {Rick 100}
	fmt.Println(myStruct.name) // Output: {Rick}

}

func myFunction(str string) string {
	return str
}
