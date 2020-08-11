package main

import (
	"fmt"
	"strings"
)

func main() {
	// 当需要频繁追加字符串时，使用strings.Builder
	b := strings.Builder{}
	b.WriteString("123 ")
	b.Write([]byte{76, 77, 78})
	fmt.Println(b.String())
}
