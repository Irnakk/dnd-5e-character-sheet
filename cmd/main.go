package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Starting func main() at:\t%v\n", time.Now())

	fmt.Println("Performing some activity inside func main()...")

	defer fmt.Printf("Finishing func main() at:\t%v\n", time.Now())
}
