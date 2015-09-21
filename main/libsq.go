// filename: libsq.go
package main

import "C"

//export sq
func sq(num int) int {
	return num * num
}

func main() {
	// We need the main function to make possible
	// CGO compiler to compile the package as C shared library
}