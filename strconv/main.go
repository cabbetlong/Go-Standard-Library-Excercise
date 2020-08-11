package main

import (
	"fmt"
	"strconv"
)

func main() {
	i := 362
	iStr := strconv.Itoa(i)
	fmt.Println(iStr)

	s := "333"
	sInt, _ := strconv.Atoi(s)
	fmt.Printf("%d\n", sInt)
}
