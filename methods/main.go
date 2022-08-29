package main

import "fmt"

type Student struct{
  Name string 
  Language string 
  studentLoan uint64
  Profession Profession
}

type Profession struct{
	Job string 
	Salary int64
	Speciality []string
}

func (s Student) GetInfo(){
	fmt.Printf("Student's name is : %s\n", s.Name)
	fmt.Printf("Student's  Language is : %s\n", s.Language)
	fmt.Printf("StudentLoan is : %d\n", s.studentLoan)
	fmt.Printf("current Profession details  : %s in : %s\n with salary : %d\n",s.Profession.Job,s.Profession.Speciality,s.Profession.Salary)
}

func (s *Student) UpdateJob(){
  s.Profession.Job = "Doctor"
}

func (s *Student) GetSalary() int64{
	return s.Profession.Salary
}
func main(){
	fmt.Println("Lets work with methods in Go")
	sara:= Student{
		Name : "Sarah",
		Language : "Italian",
		studentLoan : 80000, 
		Profession : Profession{
			Job : "Product Manager",
			Salary : 120000,
			Speciality : []string{"Phisician",",","Radiologist"},
		},
	}

	
	sara.GetInfo()
	sara.UpdateJob()
	sara.GetInfo()
	fmt.Println(sara.GetSalary())
}