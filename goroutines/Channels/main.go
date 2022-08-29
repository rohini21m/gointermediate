package main

import (
	"fmt"
	"time"
)

func main(){
	ch:= make(chan int,2)
	go func(){
	fmt.Println("lets work with buffered channels")
	for i:= 0;i<5;i++{
    fmt.Println("sending Value : ",i)
	ch <- i // sending value "i" to channel "ch"
	fmt.Println("sent value:",i)
	}
}()

time.Sleep(2* time.Second)
fmt.Println( time.Now(),"waiting for receiver message" )
fmt.Println(time.Now()," message recieved",<-ch)
fmt.Println(time.Now()," message recieved",<-ch)
fmt.Println(time.Now()," message recieved",<-ch)

fmt.Println(time.Now()," exiting")
}