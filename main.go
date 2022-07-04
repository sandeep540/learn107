package main

import (
	"fmt"
	"unicode"
)

func main() {

	arr := [3]string{"gopher123", "alpha99beta", "1cita2del3"}
	channel_string := make(chan string)

	// using for loop
	for _, element := range arr {
		go cleanString(element, channel_string)
		retString := <-channel_string
		fmt.Printf("Cleaned String for %s is %s \n", element, retString)
	}
}

func cleanString(str string, channel_string chan string) {
	var result string
	for _, char := range str {
		if unicode.IsLetter(char) {
			result = result + string(char)
		}
	}
	channel_string <- result
}
