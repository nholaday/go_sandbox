package main

import "fmt"
import "errors"

func main(){
	var myNum int
	myNum = 15
	// positive ints
	var myUnsigned uint32 = 40000
	var myFloat float32 // defaults to 0
	fmt.Println("Hey first program!")
	fmt.Println(myNum, myUnsigned, myFloat)


	var myString string = "first \nstring"
	var myString2 string = `string using backticks`
	fmt.Println(myString, myString2)

	var myRune rune = 'a' // single quotes for runes
	fmt.Println(myRune)

	var myBool bool = true
	fmt.Println(myBool)

	// inferred type
	var myInfString = "inferring a string"
	myColonString := "inferred with colon shorthand"
	fmt.Println(myInfString, myColonString)

	var var1, var2 int32 = 1, 2
	var3, var4 := 3, 4
	fmt.Println(var1, var2, var3, var4)

	const myConst string = "this is constant"
	fmt.Println(myConst)

	// Functions
	simplePrinter(myInfString)
	var divideResult int = divider(5, 3)
	fmt.Println(divideResult)
	var result, remainder, err = divideWithRemainder(5, 3)
	
	// Error Handling
	if err != nil {
		fmt.Printf(err.Error())
	}
	fmt.Printf("result: %v, remainder: %v", result, remainder)
	
	arrayPlayground()
	slicePlayground()
}

func simplePrinter(printValue string){
	fmt.Println(printValue)
}

func divider(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func divideWithRemainder(numerator int, denominator int) (int, int, error) {
	var err error  // initializes to nil
	if denominator == 0 {
		err = errors.New("Cannot Divide by Zero")
		return 0, 0, err
	}
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder, err
}

func arrayPlayground() {
	var intArr [3]int32 // [0,0,0]
	// memory location
	fmt.Println(&intArr[0])
	var intArr2 = [3]int32{1,2,3}
	// equivalent shorthands
	// intArr := [3]int32{1,2,3}
	// intArr := [...]int32{1,2,3}
	fmt.Println(intArr2)
}

func slicePlayground() {
	// slice are wrappers on arrays to give more functionality
	// with length and capacity
	var intSlice []int32 = []int32{4,5,6}
	fmt.Println(intSlice)
	intSlice = append(intSlice, 7)
	fmt.Println(intSlice)
	fmt.Printf("length: %v, cap: %v\n", len(intSlice), cap(intSlice))
	
	// when going over the limit of the capacity it will create a
	// new array with a larger capacity.  Predefine capacity for better performance
	
	var intSlice2 = make([]int32, 5, 15)
	fmt.Println(intSlice2)
	fmt.Printf("length: %v, cap: %v\n", len(intSlice2), cap(intSlice2))
	
	// appending multiple values with spread operator
	intSlice = append(intSlice, intSlice2...)
	fmt.Println(intSlice)
}