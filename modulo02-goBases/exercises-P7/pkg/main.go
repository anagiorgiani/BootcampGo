package main

import (
	"fmt"
	"os"
)

func main() {
	status, err := mycustomerErrorsTest(15000)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("status: %d\n", status)

}
