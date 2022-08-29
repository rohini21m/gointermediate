package main

import (
	"fmt"
	"time"
)

func main(){
	ch:= make(chan int,2)
	tunnel:= make (chan struct{})
	go func(){
	fmt.Println("lets work with buffered channels")
	for i:= 0;i<5;i++{
    fmt.Println("sending Value : ",i)
	ch <- i // sending value "i" to channel "ch"
	fmt.Println("sent value:",i)
	time.Sleep(1*time.Second)
	}
	fmt.Println(time.Now(),"all values are sent")
}()
close(ch)

go func(){
	// XXX: This is overcomplicated because is only channel only, "select"
		// shines when using multiple channels.
	for {
		select{
		case v,open:= <-ch:
			if !open{
				close(tunnel)
				return
			}
			fmt.Println(time.Now(),"received :",v)
		}
	}
}()
fmt.Println(time.Now(),"waiting for everyone to complete")
<- tunnel
fmt.Println(time.Now(), "exiting")
}
