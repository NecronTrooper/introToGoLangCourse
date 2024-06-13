package main

import "fmt"

func SayHello(name string) string {
	massage := fmt.Sprintf("Hello, %s", name)
	return massage
}
