package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// Declare an integer variable with type inference
	x := 10

	// Get the address of x and store it in a pointer variable with type inference
	ptr := &x

	fmt.Println("Value of x:", x) // Print the value of x
	fmt.Println("ptr ", ptr)

	ptrPtr := &ptr
	fmt.Println("ptrPtr ", ptrPtr)

	ptrPtrPtr := &ptrPtr
	fmt.Println("ptrPtrPtr ", ptrPtrPtr)

	valuePtr := *ptr
	fmt.Println("valuePtr ", valuePtr)

	// Print the value of x using the pointer
	fmt.Println("Value of x:  &  ", *&ptr) // Dereferencing the pointer to get the value//x: 0xc0000a4018
	fmt.Println("Value of x:  &  ", &*ptr) // Dereferencing the pointer to get the value//x: 0xc0000a4018

	fmt.Println("Value of x:", ptr)     // Dereferencing the pointer to get the value//x: 0xc0000a4018
	fmt.Println("Value of x: * ", *ptr) // Dereferencing the pointer to get the value//x: 0xc0000a4018

	// Modify the value of x using the pointer
	*ptr = 20

	// Print the modified value of x
	fmt.Println("Modified value of x:", x)

	var unknown *string = nil
	fmt.Println("Value of unknown:", unknown)

	firstName := "John"
	lastName := "Doe"
	fullName := firstName + " " + lastName

	fmt.Println("Full Name:", fullName)

	name := "John"
	age := 30
	message := name + "has" + strconv.Itoa(age) + "years"

	fmt.Println("Message:", message)

	var sb strings.Builder
	sb.WriteString("Hello, ")
	sb.WriteString("World!")

	fmt.Println(sb.String())

	fmt.Printf("Hello, %s!\n", "World")

	y := 10
	fmt.Println("Before incrementByReference:", y)

	incrementByReference(&y)

	incrementByValue(y)

	fmt.Println("After incrementByReference:", y)

	fmt.Printf("------------------------------------------------------------\n")

	shapes := []shape{
		rectangle{width: 10, height: 20},
		rectangle{width: 15, height: 25},
		circle{radius: 5},
	}

	for _, s := range shapes {
		fmt.Printf("Area: %f, Circumference: %f\n", s.area(), s.circumference())
	}

	var array = [5]int{1, 2, 3, 4, 5}
	fmt.Println("Array:", array[0], array[1], array[2], array[3], array[4])
}

func incrementByReference(num *int) {
	*num++

	fmt.Println("Inside incrementByReference:", *num)
}

func incrementByValue(num int) {
	num++

	fmt.Println("Inside incrementByValue:", num)
}

type shape interface {
	area() float64
	circumference() float64
}

type rectangle struct {
	width  float64
	height float64
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return 3.14 * c.radius * c.radius
}

func (c circle) circumference() float64 {
	return 2 * 3.14 * c.radius
}

func (r rectangle) area() float64 {
	return r.width * r.height
}

func (r rectangle) circumference() float64 {
	return 2*r.width + 2*r.height
}
