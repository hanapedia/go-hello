package main

import (
	"fmt"
	"strconv"
)

var i int //when declaring variables but don't want to initialize it

func main() {
	//variables
	i = 16 //variable assignment

	var j float32 = 24 //one line declaration and assignment, useful for when specifying the type of the variable like float32

	k := 32 //most common declaration and assigment

	//Go throws an error if the declared variable is never used
	//Go throws an error if a variable is redclared

	fmt.Println(i, j, k)

	/**
	naming convention:
		- lowercase variables are scoped in the package
		- uppercase variables are globally visible (across packages)
		- variables defined in a block is visible only in that block
	Scopes:
	package, block, globe

	length of variables should represent variable lifespan eg. i for iterator
	camel case
	*/

	// var l float32
	l := float32(i) //int to float *float to int throws and compile error
	fmt.Printf("int to float32: %v, %T \n", l, l)

	// s := string(i) //integer to ASCII char
	// fmt.Printf("%v, %T", s, s)

	str := strconv.Itoa(i) // converts integer to string of that integer
	fmt.Printf("Int to string: %v, %T \n", str, str)

	/**
	Primitive Types
	Boolean(bool),
	number(int8~64, uint8~32 float32, float64, complex128, complex64),
		Math operators +-*%/
		bit logical operation & | ^ &^ and or Xor andnot
		bit shift >> << adding to the exponent of 2 eg. 8 << 3 = 64, 8 >> 3 = 1
		i literal can be used to define complex numbers
	Text(string, rune)
		strings can be treated kind of like an array of ASCII bytes
		immutable, concatenate with +,
		runes are type alias fot int32 (tricky) UTF32
	*/
	n := true //Boolean can be assigned with logical operator
	fmt.Printf("Boolean: %v, type: %T\n", n, n)

	i = 8
	fmt.Println(i >> 3)
	fmt.Println(i << 3)

	var c complex64 = 1 + 2i //or complex(1, 2)
	fmt.Printf("complex number %v, %v \n", real(c), imag(c))

	s := "this is a string"
	b := []byte(s) // converts to byte ASCII literal array, used when sending data to other apps
	fmt.Printf("string as bytes: %v \n", b)

	/**
	Constants
		typed, untyped, enumerated, enmeration expressions
		const keyword added at the begining of the regular variable declaration
		named with regular naming convention
		the values of contstants cannot be changed
		constants cannot be set from something that has to be solved in the run time
		operations can be performed between variables and constants if they are the same types
		immutable but can be shadowed
		value must be calculable at compile time

		enumerated constants: usually defined at the package level
		const a = iota
		const (
			a = iota //->0
			b = iota //->1
			c = iota //->2
		)
		const (
			a1 = iota //->0
		)
		iota is scoped to each constant block eg, a,b,c are enumerated together but a1 isn't
	*/

	const ( //usually defined in the package scope
		_  = iota
		KB = 1 << (10 * iota)
		MB
		GB
		TB //compiler assumes the assignment pattern
	)
	fileSize := 400000000.
	fmt.Printf("%.2fGB\n", fileSize/GB)

	const ( //usually defined in the package scope
		isAdmin       = 1 << iota //bit shift 00000001
		isHeadquaters             //bit shift 00000010
		canSeeFinacials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica //10000000
	)

	var roles byte = isAdmin | canSeeFinacials | canSeeEurope //creates a byte with binary attributes. roles have attributes admin, financials, and europe
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is at Headquaters? %v\n", isHeadquaters&roles == isHeadquaters)

	/**
	Arrays and Slices
	declaration
	name := [size int]type{values}
	*/
	grades := [3]int{97, 88, 71} // ... can bes used declare its size if values are initialized in the same line
	fmt.Printf("Grades: %v\n", grades)
	fmt.Printf("Grades length: %v\n", len(grades))
	//when you copy an array, it only copies its values
	// use pointers to point at the same data

	//slices can do pretty match every thing that arrays can do

	a := []int{1, 2, 3} //slices defined without size.
	//slices naturally references underlying data in slices
	aCopy := a[:] //all the element [start index(inclusive):end index(exclusive)]
	fmt.Println(a)
	fmt.Println(aCopy)
	fmt.Printf("length: %v\n", len(a))
	fmt.Printf("capacity: %v\n", cap(a)) //slices have capacity function

	aSlice := make([]int, 3, 100) //type, length, capacity
	fmt.Println(aSlice)
	fmt.Printf("length: %v\n", len(aSlice))
	fmt.Printf("capacity: %v\n", cap(aSlice))
	aSlice = append(aSlice, 1)
	fmt.Println(aSlice)
	fmt.Printf("length: %v\n", len(aSlice))
	fmt.Printf("capacity: %v\n", cap(aSlice))
	aSlice = append(aSlice, a...) //spread operator
	fmt.Println(aSlice)
	fmt.Printf("length: %v\n", len(aSlice))
	fmt.Printf("capacity: %v\n", cap(aSlice))

	aShift := aSlice[1:]               //shift, mutates the original slice
	aUnshift := aSlice[:len(aSlice)-1] //unshift, mutates the original slice
	aSplice := append(aSlice[:1], aSlice[2:]...)
	fmt.Printf("Shift: %v\n", aShift)
	fmt.Printf("unShift: %v\n", aUnshift)
	fmt.Printf("Splice: %v\n", aSplice)
}
