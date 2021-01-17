package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url:= "http://www.google.com/robots.txt"
	channel:= make(chan string)
	go read(channel, url)
	fmt.Println(<-channel)
}
func read(c chan string,url string) {
	res, err := http.Get(url)

	if err != nil {
		log.Fatal(err)
	}
	robots, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
	c<-string(robots)

}
