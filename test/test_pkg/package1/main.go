package main

import (
	"fmt"
	"package2"
	"package2/hello"
)
func main() {
	package2.New()
	hello.New2()
	fmt.Println("package1.main")
}
