package main

import "fmt"

func main() {
	basics()
}

func basics() {
	i := 42
	fmt.Printf("initial value: %d\n", i)
	setValue(i)
	fmt.Printf("after setValue: %d\n", i)
	setValueInPointer(&i)
	fmt.Printf("after setValueInPointer: %d\n", i)
}

// INFO: A function always copies the value of the argument passed to it.
func setValueInPointer(i *int) {
	*i = 21
}

func setValue(i int) {
	i = 21
}
