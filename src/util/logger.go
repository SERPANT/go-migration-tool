package util

import "fmt"

func Print(something interface{}) {

	fmt.Println(something)
	fmt.Println("")
}

func Error(err error) {
	fmt.Println("Failed: ", err)
	fmt.Println("")
}
