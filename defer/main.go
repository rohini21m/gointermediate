package main


import (
	"fmt"
   // "http"
)
type Engineer struct {
	Field string
}

func (t *Engineer) GetField(name string) {
	t.Field = name
}
func UpdateName(t *Engineer){
	name := "civil-engineer"
defer t.GetField(name)
name = "software-engineer"
fmt.Println("stand By Me !!")
// url:= "www.facebook.com"
// resp,err:= http.Get(url)
// 	if err!=nil{
//     fmt.Println(err)
// 	}
//defer resp.Body.Close()
}
func main(){
	
	fmt.Println("Working with defer in GO !!") // %+n will print field name and value 
	Laila:= &Engineer{Field : "Electrical-Engineer",}
	fmt.Printf("%+v\n",Laila)
	UpdateName(Laila)
	fmt.Printf("%+v\n",Laila)
	
}