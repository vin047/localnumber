package main

import (
	"fmt"
	"os"

	"github.com/vin047/localnumber/pkg/localnumber"
)

func main() {
	input := os.Args[1]
	number, err := localnumber.Format(input)
	if err != nil {
		panic(err)
	}
	fmt.Print(number)
}
