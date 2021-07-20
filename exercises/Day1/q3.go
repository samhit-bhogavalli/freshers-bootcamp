package main

import "fmt"

//interface for all types of employes
type Employee interface {
	calculateSalary() int
}

//struct for fulltime employee
type fullTime struct {
	basic int
}
//salary for fulltime employee
func (f fullTime) calculateSalary() int {
	return 28*f.basic
}

//struct for contract employee
type contractor struct {
	basic int
}
//salary for contract employee
func (c contractor) calculateSalary() int {
	return 28*c.basic
}

//struct for freelancer
type freelancer struct {
	basic int
	hours int
}
//salary for freelancer
func (f freelancer) calculateSalary() int {
	return f.basic*f.hours
}

//function for calculating total expense for the employer
func totalExpence(employes []Employee) int {
	total := 0
	for i := 0; i < len(employes); i++ {
		total += employes[i].calculateSalary()
	}
	return total
}

func main(){
	//testing
	emp1 := fullTime{500}
	emp2 := contractor{100}
	emp3 := freelancer{10,20}
	var employes = []Employee{emp1, emp2, emp3}
	fmt.Println(totalExpence(employes))

}
