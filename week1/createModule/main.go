package main

import (
	"fmt"
	"introToGoLangCourse/week1/moreStrings"
)

func main() {
	var a string
	fmt.Scan(&a)
	fmt.Println(SayHello(a))
	fmt.Println(moreStrings.ReverseStrings(SayHello(a)))

}
