package main /** When Go sees that a package is named main it knows the package
should be considered a binary, and should be compiled into an executable file,
instead of a library designed to be used in another program. */

import (
	"fmt"
	"simpleModule/shapes"
)

func main() {
	fmt.Println("Hello Modules")
	shapes.PrintShape()
}
