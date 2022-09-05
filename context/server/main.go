package main

import(
	"github.com/rohini21m/GOINTERMEDIATE/context/log"
	"time"
	"fmt"
	"net/http"

)

func main(){
	http.HandleFunc("/",log.Decorate(handler))
	log.Fatal(http.ListenAndServe("127.0.0.1:3000",nil))
}

func handler(w http.ResponseWriter,r *http.Request){
	ctx:= r.Context()
	log.Printf("server started")
	defer log.Printf("server ended")
select {
  case<-time.After(5*time.Second):
	fmt.Fprintln(w,"Hello , What a wonderful World !!")

case<-ctx.Done():
	err:= ctx.Err()
	log.Print(ctx,err.Error())
	http.Error(w,err.Error(), http.StatusInternalServerError)
}
}


