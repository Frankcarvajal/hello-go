package main
import (
	"fmt"
	"errors"
	"math"
)

// A struct is a collection of fields so that I can group things together to create a more logical type
// use the word type and the name and struct word, each field has to have a name and a type
type person struct {
	name string
	age int
	birthplace string
}

func main() {
	fmt.Println("Hello Franklin...")
	var x int = 5 // declaring variable of integer type
	y := 20 // := shorthand, neater to read, and go infers it's type.
	sum := x + y
	fmt.Println(sum)

	if x > 6 { // If, else if, else conditionals also existant
		fmt.Println("x is greater than 3!")
	} else if x > 2 && x < 6 { // Operator support as typically thought
		fmt.Println("x is greater than 2, but less than 6!")
	}

	var array [5]int // Declaring an array with 5 ints defaulted to 0
	array[2] = 4 // assignment of index 2 of the array to 4
	fmt.Println("First array: ", array)
	fmt.Println("1st array index 2: ", array[2])

	arrayTwo := [4]int{1, 2, 3, 4} // Shorthand specific array declaration
	fmt.Println("Second array: ", arrayTwo)
	fmt.Println("2nd array index 3: ", arrayTwo[3])

	// setting array like this means the length is apart of the type, to overcome this I will
	// need to use slices which is abstraction on top of arrays to make them easier to work with
	arrayThree := []int{5, 8, 123, 99} // slice event by removing element amount
	fmt.Println("Third array: ", arrayThree)
	fmt.Println("3rd array index 2: ", arrayThree[2])
	arrayThree = append(arrayThree, 23) // append does not add to original slice it returns a new one.
	arrayFour := append(arrayThree, 13) // append creates a new array, here I set it equal to new variable
	fmt.Println("Fourth array: ", arrayFour)
	fmt.Println("Third array appended: ", arrayThree)

	// maps hold key value pairs kind of like JS objects or dictionaries in python
	// type def is map and in square bracket type of keys followed by type of values in the map
	// to create a map use built in make function and give it this type
	verticles := make(map[string]int)
	verticles["triangle"] = 3 // we can assign keys similar to object bracket notation.
	verticles["square"] = 4
	verticles["octogon"] = 8
	fmt.Println("verticles: ", verticles)
	// Can use same syntax to get particular value of a key
	fmt.Println("verticles['octogon']: ", verticles["octogon"])
	// Can also use the delete built in function to delete something from the map
	delete(verticles, "triangle")
	fmt.Println("verticles: ", verticles)

	// Loops, the only type of loop in Go is the for loop
	// I declare and initialize a variable using shorthand syntax, then the counter, then the stopping condition, then increment or whatever
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	// The for loop also doubles as a while loop if I drop the variable declaration and the increment
	j := 0
	for j > 5 {
		fmt.Println(j)
		j++
	}

	// Another thing that can be done within a loop is iterate over each element within in an array or slice by using range
	arrayFive := []string{"a", "b", "c", "d"}

	for index, value := range arrayFive{
		fmt.Println("index: ", index, "value: ", value)
	}

	// Let me do the same thing with a map
	mapTwo := make(map[string]string)
	mapTwo["a"] = "alpha"
	mapTwo["b"] = "beta"
	mapTwo["o"] = "omega"

	for key, value := range mapTwo{ // However will get key instead of index
		fmt.Println("key: ", key, "value: ", value)
	}

	// Called function from elsewhere
	result := getsum(2, 3)
	fmt.Println(result)

	// calling sqroot function getting result and err
	result2, err := sqroot(16)
	// if the error is not nil that means something went wrong so outputting that else results
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result2)
	}

	// Creating a struct of that type
	thepersonthatisme := person{name: "Franklin", age: 26, birthplace: "Carmichael, CA"}
	fmt.Println("Me: ", thepersonthatisme)
	// Dot notation to get specific field out of the struct
	fmt.Println("My Birthplace: ", thepersonthatisme.birthplace)

	// Pointers
	pointerVariable := 7
	// If we pass a pointer to the variable then the function is going to be able to look at the value at that memory address
	incrementVar(&pointerVariable) // Send the pointer by using ambersand
	fmt.Println("pointerVariable: ", pointerVariable)
	// I can get the memory address by using the amperstand giving us a pointer to the variable
	fmt.Println("&pointerVariable: ", &pointerVariable)
}

// So far everything is in the main function, but I can easily create a new function
// If i want to accept arguments just have to give the name and a type
// and if I want a return value I need to give the type of that after the brackets
// then invoke in main func
func getsum(x int, y int) int {
	return x + y
}

// in Go I can have multiple return values
func sqroot(x float64)(float64, error) {
	// All built in types and since the square root of negative numbers are complex numbers
	// I can limit this function to positive integers only if x is less than 0
	// use errors package from errors import
	if x < 0 {
		return 0, errors.New("Do not use negative numbers!")
	}

	return math.Sqrt(x), nil
}

func incrementVar(x *int){ // Astric accepts the pointer prefixing type
	*x++ // To modify the variable directly not the memory which would be x without Astrics
}
