package main

import (
	"fmt"
	"reflect"
	"strings"
)

func main() {
	fmt.Println(reflect.TypeOf(3.14))
	fmt.Println(strings.Title("go developers"))
	fmt.Println("Kim\nInha\t\"\\")
	fmt.Println('A', '가')
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(91))
}
