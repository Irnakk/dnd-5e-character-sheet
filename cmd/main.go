package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Printf("Starting func main() at:\t%v\n", time.Now())
	defer fmt.Printf("Finishing func main() at:\t%v\n", time.Now())
}
