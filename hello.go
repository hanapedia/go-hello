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

	aShift := aSlice[1:]                         //shift, mutates the original slice
	aUnshift := aSlice[:len(aSlice)-1]           //unshift, mutates the original slice
	aSplice := append(aSlice[:1], aSlice[2:]...) //splice, mutates the original slice
	fmt.Printf("Shift: %v\n", aShift)
	fmt.Printf("unShift: %v\n", aUnshift)
	fmt.Printf("Splice: %v\n", aSplice)

	/**
	Maps and Structs
	return order of the map is not guaranteed
	*/
	// statePopulations := make(map[string]int)
	statePopulations := map[string]int{ //map[type of key]type of value
		"California": 39250017,
		"Texas":      27862596,
	}
	delete(statePopulations, "Texas")
	pop, ok := statePopulations["Ohio"] //if the key is not defined the second return value will be false, pop holds 0
	fmt.Println(statePopulations)
	fmt.Println(pop)
	fmt.Println(ok)

	//struct created with type declaration
	type Doctor struct { //caitalize the field names so that it is exported
		number     int
		actorName  string
		companions []string
	}
	//when you pass around the struct, it only copies the data. Assign with pointer to point at the same data

	aDoctor := Doctor{
		number:    3,
		actorName: "john",
		companions: []string{
			"Joe",
			"Steph",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.companions[0])

	//inheritance like implementation (Embedding)("Has" reltaionship and not "is" )
	type Animal struct {
		Name   string //`required max:"100"` //Tag which can be accesed via reflect package (validations are handled in other libraries)
		Origin string
	}
	type Bird struct {
		Animal
		SpeedKPH float32
		CanFly   bool
	}
	bird := Bird{
		Animal:   Animal{Name: "Dodo", Origin: "Australia"},
		SpeedKPH: 200,
		CanFly:   false,
	}
	fmt.Println(bird.Name)

	// t := reflect.TypeOf(Animal{})
	// field, _ := t.FieldByName("Name")
	// fmt.Println(field.Tag) //accessing tags

	/**
	if, switch statements
		conditions give in boolean
		== > < <= >=
		|| && !  Go lazily evaluates or and eg. if the first value in the or condition returns true, it shortcircuits and does not compute the rest of the condition
		when using floating number in comparison, generate an error number to check if they are same
	*/
	if true { //if condition {executions}
		fmt.Println("Simplest if statement")
	}

	if pop2, ok2 := statePopulations["California"]; ok2 { //initializer syntax
		fmt.Println(pop2)
	}

	//switch
	switch 2 { //switch key {cases}
	case 1:
		fmt.Println("One") //break keywords are unneeded
	case 2, 3, 5:
		fmt.Println("Two, Three, or Five")
	default:
		fmt.Println("Another number")
	}

	switch i2 := 2 + 1; i2 { //initialization syntax
	case 1:
		fmt.Println("One")
	case 2, 3, 5:
		fmt.Println("Two, Three, or Five")
	default:
		fmt.Println("Another number")
	}

	//type switches
	var i3 interface{} = 1
	switch i3.(type) {
	case int:
		fmt.Println("Intger")
	case string:
		fmt.Println("String")
	case [3]int:
		fmt.Println("[3]int")
	default:
		fmt.Println("another type")
	}

	/**
	for loop
	no paranthesis
	for initializer; test; operation{}
	for test{}
	for{}
	*/
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
	for j, k := 0, 0; j < 5; j, k = j+1, k+1 {
		fmt.Println(j, k)
	}

	//while loop no while key word
	inc := 0
	for inc < 5 {
		fmt.Println(inc)
		inc++
	}

	for {
		inc++
		if inc == 5 {
			fmt.Println("Continueing")
			continue //break out of the current iteration
		}
		fmt.Println(inc)
		if inc == 7 {
			break //break out of unconditioned for loop
		}
	}

	//break statement
Label: //label syntax
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			fmt.Println(i, j)
			if j == 3 {
				fmt.Println("Breaking to labeled loop")
				break Label // you can specify which loop to break out of. without any label, break will only break out of the closest loop
			}
		}
	}

	forSlice := []int{1, 2, 3, 4}
	for k, v := range forSlice { //range keyword is used to loop through corrections
		fmt.Println("Looping slice:", k, v)
	}
	for _, v := range statePopulations { //use underline if you don't need key but you want values
		fmt.Println("Looping maps:", v)
	}

	/**
	Defer, Panic, and Recover
		Defer:
			defered functions are excuted after the main function but before the main function returns
			last function defered will be the first to be executed
			defer is usually used to associate opening of resource and closing of resource close to each other
			res, err := http.Get(.......)
			defer res.Body.Close()
			**defered function takes the argument of when the function is defered. if the argument variable is changed after defering, defered function will use the value of that variable prior to the change
		Panic: similar to exeptions in the other languages
			Panic kills the application with message
			use built in panic function to generate panic manually
			regular errors are not considered a panic, so we have to decide whether that error is a problem or not
			**panics are excuted after defered functions are excuted and main function returns
		Recover:
			defer func() {
				if err := recover(); err != nil {
					log.Println("Error:", err)
				}
			}()
		This is like a catch statement. recover function retrieves any panic and allows you to handle the panic
	*/

}
