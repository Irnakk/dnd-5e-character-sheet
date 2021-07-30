package main

import (
	"fmt"
)

func main() {
	fmt.Println("Starting func main()...")
	defer fmt.Println("Finishing func main()...")
}
