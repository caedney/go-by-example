package main

import "os"

func main() {
	// panic("a problem")

	_, err := os.Create("panic/panic.txt")
	if err != nil {
		panic(err)
	}
}
