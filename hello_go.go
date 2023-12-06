package main

import (
	"container/list"
	"errors"
	"fmt"
	"reflect"
	"unicode/utf8"
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

	myString = "I can change but... \nI need to be a string" //Line breaker
	fmt.Println(myString)                                    //I can change but I need to be a string

	myString = `Another way
to
change
the line`

	fmt.Println(len("Test"))
	fmt.Println(utf8.RuneCountInString("Test"))

	fmt.Println(myString)

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

	// If we declare a variable with no value, the default value will be "" (string), 0 (int, float and rune)or false (Booleans)
	var defaultString string
	var defaultInt int8
	var defaultBool bool

	// You can also initialize multiple variables at once
	var var1, var2 = 1, 2 // Short version: var1, var2 := 1, 2 || Long version: var var1, var2 int8 = 1, 2

	fmt.Println(var1 + var2)

	fmt.Println(defaultString)
	fmt.Println(defaultInt)
	fmt.Println(defaultBool)

	// Constants:

	const myConst = "Hello there! I am a constant" // It must init with a value
	const pi float32 = 3.1415                      // Constant is a good choice for pi since it won't change

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

	var myArray [3]int   // Create an array with 3 positions initialized in default values for int: 0
	fmt.Println(myArray) // Output: [0 0 0]
	myArray[0] = 1
	myArray[1] = 2
	myArray[2] = 3
	fmt.Println(myArray[1:3]) // Access to index 1 and 2. Output: [2,3]
	fmt.Println(myArray)      // Output: [1 2 3]

	fmt.Println(&myArray[0]) // Output the memory location byt adding & before the variable

	/* 	Other ways to do it:

	var intArr [3]int32 = [3]int32{1,2,3}
	intArr := [3]int32{1,2,3}
	intArr :=[...]int32{1,2,3}
	*/

	// Slices:

	var intSlice []int32 = []int32{4, 5, 6}

	intSlice = append(intSlice, 7) // Create a new Array with 6 elements capacity but the length will be 4

	var intSlice2 []int32 = []int32{8, 9}
	intSlice = append(intSlice, intSlice2...) // We can use the spread operator

	var intSlice3 []int32 = make([]int32, 3, 8) // With the function we can create an array and specify the length (3) and the capacity (8). If not, the capacity will be the same as the length.
	fmt.Println(intSlice3)
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

	var stringToPrint string = myFunction("Hello go") //This version just return the parameter in the argument as a string, it is a good practise to write the return type, not using := shortcut
	println(stringToPrint)

	// Estructura

	type MyStruct struct {
		name  string
		power int
	}

	myStruct := MyStruct{"Rick", 100}
	fmt.Println(myStruct)      // Output: {Rick 100}
	fmt.Println(myStruct.name) // Output: {Rick}

	// Handle errors:

	var result, remainder, err = intDivision(11, 2)
	if err != nil { // err is initialized with nil value as default
		fmt.Println(err.Error())
	} else if remainder == 0 {
		fmt.Printf("The result of the integer division si %v with remainder", result)
	} else {
		fmt.Printf("The result of the integer division si %v with remainder %v", result, remainder) // Another print function to format the string with values
	}

	/* Switch option:

	switch{
	case err!=nil:
		fmt.Println(err.Error())
	case remainder==0:
		fmt.Printf("The result of the integer division si %v with remainder", result)
	default:
		fmt.Printf("The result of the integer division si %v with remainder %v", result, remainder)
	}

	*/

	switch remainder { //Checking how was the variable result
	case 0:
		fmt.Println("The division was exact")
	case 1, 2:
		fmt.Println("The division was close")
	default:
		fmt.Println("The division was not close")
	}

}

func myFunction(str string) string { // Last type means the return value type
	return str
}

func intDivision(numerator int, denominator int) (int, int, error) { // Three return values, three declared types
	var err error // initialized with nil value by default
	if denominator == 0 {
		err = errors.New("Cannot divide by zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}
