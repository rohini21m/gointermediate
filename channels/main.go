package main

import (
	"fmt"
	"sync"
)


// myCh:= make(chan int) way to initialise a channel and passing value into it.
// myCh <- 10
// fmt.Println(<- myCh)
func main(){
	fmt.Println("lets look at channels in Golang")
	myCh:= make(chan int, 2)
	wg:= &sync.WaitGroup{}
	wg.Add(2)
	go func( ch chan<- int,wg *sync.WaitGroup){ // send only channel
		//close(myCh) // this gives the signal value 0 as we are trying to csend values to a closed channel.
		//myCh <- 10 // inorder to differenciate valuse zero being sent to channel and between signal zero we need use IsChanOpen
		myCh<-100
	  myCh<-200 
	  close(myCh) //these gives values true 100 , 200 , 0 
		wg.Done()
	}(myCh,wg)
	
	go func(ch <-chan int,wg *sync.WaitGroup){ // receive only channel  this doesnot alow to close the channel.
		val, IsChanOpen:= <- myCh 
		fmt.Println(IsChanOpen,val) // gives flase 0 , 0 
		fmt.Println(<- myCh)
		fmt.Println(<- myCh)
		wg.Done()
	}(myCh, wg)
	
	wg.Wait()
}
