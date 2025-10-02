package main

import (
	"fmt"
	"errors"
	"strings"
)

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
	fmt.Printf("result: %v, remainder: %v\n", result, remainder)
	
	arrayPlayground()
	slicePlayground()
	mapPlayground()
	stringPlayground()
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
	fmt.Println("############# Arrays ###############")
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
	fmt.Println("############# Slices ###############")
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
	
	// LOOPS
	// "for in" syntax
	for idx, val := range intSlice{
		fmt.Printf("idx: %v, val: %v\n", idx, val)
	}
	
	// "while" loops
	var i int = 0
	for i < 10 {
		i++;
	}
	
	// goes until break is hit
	j := 0;
	for {
		j++
		if j == 10 {
			break
		}
	}
	
	// C++ Syntax: initialization, condition, post
	for i:=0; i<10; i++ {
		fmt.Printf("%v, ", i)
	}
	fmt.Println("")
}

func mapPlayground() {
	fmt.Println("############# Maps ###############")
	
	var myMap map[string]uint8
	fmt.Println(myMap)
	// or initialize like this
	var myMap2 = map[string]uint8{"Jocca": 19, "Smocca": 33}
	fmt.Println(myMap2["Jocca"])
	// if not found will not raise a KeyError, instead will return default value
	// use the second return param to get if key found or not
	fmt.Println(myMap2["Samantha"])
	var age, keyFound = myMap2["Psyduck"]
	fmt.Printf("age: %v, keyFound: %v\n", age, keyFound)
	
	delete(myMap2, "Jocca")
	myMap2["Jerry"] = 56
	
	// LOOPS
	// "for in" syntax
	for name := range myMap2 {
		fmt.Printf("Name: %v\n", name)
	}
	
	for key, val := range myMap2 {
		fmt.Printf("Name: %v, Age: %v\n", key, val)
	}
}

func stringPlayground() {
	// Go uses utf-8 encoding for characters
	// utf-8 has variable byte length representing characters unlike ASCII
	var myString = "Résumé"
	// value will be the utf-8 position of the character
	var indexed = myString[1]
	fmt.Printf("%v, %T, length: %v\n", indexed, indexed, len(myString))

	// é is length 2 bytes so length of the total array will be 8 instead of 6
	// to avoid this hassle cast to an array of runes
	var myString2 = []rune("Résumé")
	var myRune = myString2[1]
	fmt.Printf("%v, %T, length: %v\n", myRune, myRune, len(myString2))
	
	// strings are immutable in Go, can't modify
	// to deal with this use the strings package
	var strSlice = []string{"n", "a", "m", "e"}
	var strBuilder strings.Builder
	for i := range strSlice{
		strBuilder.WriteString(strSlice[i])
	}
	// waits until operations are finished before creating a real string
	var catStr = strBuilder.String()
	fmt.Printf("%v\n", catStr)
}