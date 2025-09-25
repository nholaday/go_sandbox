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

	// inferred type
	var myInfString = "inferring a string"
	myColonString := "inferred with colon shorthand"
	fmt.Println(myInfString, myColonString)

	var var1, var2 int32 = 1, 2
	var3, var4 := 3, 4
	fmt.Println(var1, var2, var3, var4)

	const myConst string = "this is constant"
	fmt.Println(myConst)

	simplePrinter(myInfString)
	var divideResult int = divider(5, 3)
	fmt.Println(divideResult)
	var result, remainder int = divideWithRemainder(5, 3)
	fmt.Printf("result: %v, remainder: %v", result, remainder)
}

func simplePrinter(printValue string){
	fmt.Println(printValue)
}

func divider(numerator int, denominator int) int {
	var result int = numerator / denominator
	return result
}

func divideWithRemainder(numerator int, denominator int) (int, int) {
	var result int = numerator / denominator
	var remainder int = numerator % denominator
	return result, remainder
}