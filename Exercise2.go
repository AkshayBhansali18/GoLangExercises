package main

import "fmt"

func main() {
	
	fmt.Println("Please Enter a number")
	var inputNumber int
	fmt.Scanln(&inputNumber)
	if inputNumber%2 == 0 {
		fmt.Println(inputNumber," is even")
	}else{
		fmt.Println(inputNumber," is odd")
	}
	
			
}

