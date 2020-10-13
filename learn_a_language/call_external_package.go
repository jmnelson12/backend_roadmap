package main

import 	"fmt"
import "rsc.io/quote"

func main() {
	fmt.Println(quote.Go())
}

/*

1. published modules here: pkg.go.dev
2. add code
3. Put your own code in a module for tracking dependencies.
	a. create go.mod file using `$ go mod init hello`
	b. $ go run call_external_package.go
*/