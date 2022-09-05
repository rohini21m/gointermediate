package main

import (
	"log"
	"context"
)

const reqID = 45
func Println(ctx context.Context,msg string){
	id,ok:= ctx,Value(reqId).(int64)
	if !ok{
    log.Println("Could not print reqID from conetxt")
	return 
	}
    log.Println("%d","%s",id,msg)

}
	// below code is doing following : receiving the context , adding id value to the context
	//and sending it  back using as wrapper function & decorate.

	func decorate(f http.HandlerFunc) http.HandlerFunc{
		return (w http.ResponseWriter, r *http.Request){
			ctx:= r.Context()
			id:= rand.Int63()
			ctx = context.WithValue(ctx,reqID,id)
			f(w,r.WithContext(ctx))
		}
	}

