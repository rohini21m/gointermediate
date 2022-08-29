package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

func main(){
	a := asChan(1,2,3,4)
	b:= asChan(5,6,7,8)
	c:= merge(a,b)
	for  v := range c{
		fmt.Println(v)
	
	}
	
}
func merge(a,b<- chan int) <- chan int{
	c:=make(chan int)
	go func(){
		defer close(c) // making sure c is also closed after
		// receiving from a,b otherwise it will result in deadlock.
		//adone,bdone:=false,false
for a!= nil || b!= nil {//!adone||!bdone{
	select {
	case v,ok := <-a:
		if !ok{
			a=nil//adone=true
			log.Println("a is done")
			continue
		}
		c<-v
	
	case v, ok :=<-b:
		if !ok{
			b=nil//bdone=true this is true block "a" forever in select 
			log.Println("b is done")
			continue
		}
		c<-v
	
	}
}
	
	}()
	return c  
}
func asChan(vs ... int) <- chan int {//<- chan int: this indicates it can only receive values 
    c := make(chan int)
	go func(){
		for _,v := range vs{
			c <-v
			time.Sleep(time.Duration(rand.Intn(1000))*time.Millisecond)
		}
		close(c)
	}()
	return c 
	}