package main

import "fmt"

func main() {

	// Strings
	// Explicit declaring String Variables with the var keyword
	var myName string = "Jesse init" // Explicit variable declaration and assignment
	var anotherName = "Bigman Thing" // Implicit variable declaration and assignment
	var someOtherName string         // Variable instantiation without assignment

	someOtherName = "A cool string value" // Variable assignment operation

	// Implicit declaration with := notation - the shorthand version
	// This shorthand method cannot be used outside of a function
	coolVariable := "The Man Them"

	fmt.Println(myName, anotherName, someOtherName, coolVariable)

	//Ints
	var myAge int = 29
	var elderBroAge = 30
	elderSisAge := 31
	fmt.Println("Ints>>>", myAge, elderSisAge, elderBroAge)

	// Bits and Memory
	// Integers can be of various bit sizes and that detemines the range of numbers the can hold
	// Integers can be set to be signed(allows both positive and negative numbers) or unsigned(only positives numbers)
	var absoluteScore int8 = 30
	fmt.Println("Ints>>>", absoluteScore)

	//Floats
	var accountBalanceOne float32 = 0.453
	var accountBalanceTwo float64 = 354654746746846857865.453
	accountBalanceThree := 45.78 // Float64 is inferred with the shorthand notation(this has got to have an official name)
	fmt.Println("Float>>>", accountBalanceOne, accountBalanceTwo, accountBalanceThree)

}
