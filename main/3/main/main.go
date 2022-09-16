package main

import (
	"fmt"
	"strings"
)

func main() {
	str := "abcdef"
	slice := make([]string, len(str)+1)
	fmt.Println(fmt.Sprintf(str+strings.Join(slice, "%v"), "ghijk"))
}
