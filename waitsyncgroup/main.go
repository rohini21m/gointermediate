package main 

import (
	"fmt"
	"sync"
	"net/http"


)
var streets = []string{"teachers colony"}
var wg sync.WaitGroup//pointer
var mut sync.Mutex //pointer

func main(){
	fmt.Println("lets work with wait sync group")
	
websites := []string {"https://www.google.com","https://www.facebook.com","https://www.gmail.com",} // slice syntax
for _, web:= range websites{
   go StatusCode(web)
   wg.Add(1)

}
wg.Wait()
}

func StatusCode(endpoint string) {
	
	defer wg.Done()
	
	res,err := http.Get(endpoint)

	if err!=nil{
		fmt.Println( "Could not get the code")
	} else{
	mut.Lock()
	streets = append(streets,endpoint)
	mut.Unlock()
	fmt.Printf( "%s %s - %d\n",streets,endpoint, res.StatusCode)

	}
}