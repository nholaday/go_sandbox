package main

import "fmt"

type student struct {
	name string
	age uint8
	major string
	// can add a struct to the type -- student.schoolInfo.name
	schoolInfo school
	// or can add the subfields directly the type -- student.year, student.paidOff
	studentStatus
}

type school struct {
	name string
}

type studentStatus struct {
	year uint8
	paidOff bool
}

// methods on structs (like private functions on a class)
func (s student) yearsLeft() uint8 {
	return 4 - s.year
}

type teacher struct {
	name string
}

func (t teacher) yearsLeft() uint8 {
	return 100
}

type person interface {
	yearsLeft() uint8
}

func getMonthsLeft(p person) uint8 {
	return p.yearsLeft() * 12
}

func main(){
	// uninitialized will have default type values
	// "", 0, ""
	var joe student
	fmt.Println(joe.name, joe.age, joe.major)
	
	var tuBerlin school = school{"TU Berlin"}
	// A lot like instantiating a class and calling methods on it
	var mary student = student{"Mary", 39, "Japanese", tuBerlin, studentStatus{3, true}}
	fmt.Println(mary.name, mary.age, mary.major, mary.schoolInfo.name, mary.year)
	fmt.Println(mary.yearsLeft())
	
	// anonymous struct (not reusable)
	var myEngine = struct{
		mpg uint8
		gallons uint8
	}{25,15}
	fmt.Println(myEngine.mpg, myEngine.gallons)
	
	// Interfaces can be like an abstract base class
	// For this interface to work all subclasses just need a yearsLeft method
	var castle teacher = teacher{"Ms. Castle"}
	fmt.Println(getMonthsLeft(mary))
	fmt.Println(getMonthsLeft(castle))
}