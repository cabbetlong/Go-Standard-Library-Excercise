package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	bufrd := bufio.NewReader(strings.NewReader("Hello, world!\nThis is second line."))
	// b, _ := bufrd.ReadByte()
	// fmt.Printf("%c\n", b) // Output: H

	line1, _ := bufrd.ReadBytes('\n') // 会将界定符一起读出
	fmt.Printf("%s", line1)           // Output: Hello, world! 界定符为\n，因此line1自带换行
	line2, _ := bufrd.ReadBytes('\n')
	fmt.Printf("%s\n", line2) // Output: This is second line.
}
