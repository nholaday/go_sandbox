package main

import "fmt"
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
}