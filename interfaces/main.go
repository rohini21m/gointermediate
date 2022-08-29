package main 

import (
	"fmt"
	"strconv"
	//"strings"
)

type Company interface{
	GetNames() string

}

type Employee struct{
  Name string
  Age int
}
type Manager struct{
	Name string
}
func (e *Employee) GetNames() string{
return "Name of Employee is : " + e.Name +", " + "Age is : " + strconv.Itoa(e.Age)
	
}

func (m *Manager) GetNames() string{
	return "manager is : " + m.Name
}


func PrintDetails(e Company){
fmt.Println(e.GetNames())

}
	func main(){
		name:= &Employee{Name : "Rohini", Age: 25}
		manager:= &Manager{Name : "Brian King"}
		PrintDetails(name)
		PrintDetails(manager)
		

	}