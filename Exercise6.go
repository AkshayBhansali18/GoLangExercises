package main

import "fmt"

func main(){
	t:=make(chan int)
	go sum(t)
	fmt.Println(<-t)
}
func sum(chanell chan int) {
	s := 0
	for i := 1; i <= 50; i++ {
		if i%3 == 0 || i%5 == 0 {
			s += 1

		}


	}
	chanell<-s
}

