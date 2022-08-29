package main

import (
	"fmt"
	"sync"
)

func main(){
	wg:= &sync.WaitGroup{}
	mut:= &sync.Mutex{}
	var scores = []int{0}
	wg.Add(3)
   go func(wg *sync.WaitGroup, m *sync.Mutex){
	   mut.Lock()
	   scores = append(scores,100)
	   mut.Unlock()
	   fmt.Println("Rohini1' score") 
	   wg.Done()
   }(wg,mut)
   go func(wg *sync.WaitGroup,m *sync.Mutex){
	  
	fmt.Println("Rohini2' score") 
	mut.Lock()
	scores = append(scores,200)
	mut.Unlock()
	wg.Done()
}(wg,mut)
go func(wg *sync.WaitGroup,m *sync.Mutex){

	fmt.Println("Rohini3' score") 
	mut.Lock()
	scores = append(scores,300)
	mut.Unlock()
	wg.Done()
}(wg,mut)
wg.Wait()
fmt.Println(scores)
}
