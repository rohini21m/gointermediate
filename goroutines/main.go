package main

import (
	"fmt"
	"time"
)

func main(){
  go greet("Mehraba  Nazli !!")
  greet("Salem Ferit")
  ch := make(chan string)
	go func() {
		fmt.Println(time.Now(), "Good Evening")
		time.Sleep(4 * time.Second)
		ch <- "Seni seviyorum"

	}()

	fmt.Println(time.Now(), "Hello, Gunayidin")
	v := <-ch
	fmt.Println(time.Now(), "msg sent :", v)
}

func greet(s string){
	for d:=0; d<6; d++{
		time.Sleep(3*time.Millisecond)
		fmt.Println(s)
	}

}
