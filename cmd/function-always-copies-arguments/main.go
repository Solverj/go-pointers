package main

import (
	"fmt"
)

func main() {
	i := 42
	fmt.Printf("inside main %p\n", &i)
	argumentsAlwaysCopiesValues(i)
}

func argumentsAlwaysCopiesValues(i int) {
	fmt.Printf("inside printAddrress of value %p\n", &i)
}
