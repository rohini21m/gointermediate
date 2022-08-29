package main 

import "fmt"

type Manager struct {
	Name string 
	Age int 
}

func (m *Manager) UpdateAge(){ //pointer receiver 
	                           // func (m Manager)  UpdateAge(){}  example of value receiver
	m.Age+= 1 

}
func (m *Manager)UpdateName(){
	m.Name = "dolly"
	fmt.Println(m)
}

func UpdateAge(m *Manager){
	m.Age+= 4
}
func main(){
	var name string 
	name = "patrick"
	fmt.Println(name)
	ptr:= &name
	fmt.Println(*ptr)
	*ptr = "Ramu"
	fmt.Println(name)
	sam := &Manager{Name : "Samantha", Age : 24}
	fmt.Println(sam)
	sam.UpdateAge()
	fmt.Println(sam)
	sam.UpdateName()
	fmt.Println(sam)
	jyo:= &Manager{Name: "Jyo", Age : 22}
	UpdateAge(jyo)
	fmt.Println(jyo)
}