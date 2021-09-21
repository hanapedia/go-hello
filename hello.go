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

	const (
		isAdmin = 1 << iota
		isHeadquaters
		canSeeFinacials

		canSeeAfrica
		canSeeAsia
		canSeeEurope
		canSeeNorthAmerica
		canSeeSouthAmerica
	)

	var roles byte = isAdmin | canSeeFinacials | canSeeEurope //creates a byte with binary attributes
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin&roles == isAdmin)
	fmt.Printf("Is at Headquaters? %v\n", isHeadquaters&roles == isHeadquaters)
}
