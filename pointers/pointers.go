package main

import "fmt"

func main() {

	a := "string"
	// Pass as a copy. Variable will not change...
	testPointer(a)
	fmt.Printf("a: %s\n", a)
	testPointerPassAddress(&a)
	fmt.Printf("&a: %s\n", a)

	b := []string{"string"}
	testSlices(b)
	fmt.Printf("b(slice): %s\n", a)

	c := []string{"string"}
	testAppend(c)
	fmt.Printf("c(append): %s\n", c)

}

func testAppend(b []string) {
	b = append(b, "another string")
}

func testSlices(b []string) {
	b[0] = "another string"
}

func testPointer(a string) {
	a = "another string"
}

func testPointerPassAddress(a *string) {
	*a = "another string"
}
