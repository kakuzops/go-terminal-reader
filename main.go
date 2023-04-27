package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Printf("%#v\n", os.Args)
	fmt.Println("path:", os.Args[0])
	fmt.Println("first:", os.Args[1])
	fmt.Println("second:", os.Args[2])
	fmt.Println("second:", os.Args[3])

	fmt.Println("number of items inside os.,args", len(os.Args))
}
