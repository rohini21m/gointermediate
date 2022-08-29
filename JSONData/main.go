package main

import (
	"fmt"
	"encoding/json"
	
)
type Info struct {
	State string `json:"state`
	City  string `json:"city"`
	Timezone string `json:time_zone`
	Religion faith `json:"religion"`
}

type Country struct {
	State string `json:"state`
	City  string `json:"city"`
	District string `json:district`
	
}
type faith struct{
	Name string `json:"name"`
} 
func main(){

	fmt.Println("we will look at converting json data  to struct and marshal n unmarshal method.")
	religion := faith{
		Name : "Christianity",
	}

infomation:= Info{State:"LOUISIANA",City: "Lake Charles", Timezone:"CST", Religion:religion}

polls := `{"state":"Texas","city": "Dallas", "district":"krishna"}`
data ,err:= json.MarshalIndent(infomation,"","  ")
if err!=nil{
	fmt.Println(err)
	}
var result Country
 json.Unmarshal([]byte(polls),&result)

fmt.Println(string(data))
fmt.Printf("%s\n",result)

}


