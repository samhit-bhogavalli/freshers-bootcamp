package main

import "fmt"

//interface for all types of Employee
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
func totalExpense(employees ...Employee) int {
	total := 0
	for _ , employee := range employees {
		total += employee.calculateSalary()
	}
	return total
}

func main(){
	//testing
	emp1 := fullTime{500}
	emp2 := contractor{100}
	emp3 := freelancer{10,20}
	var employees = []Employee{emp1, emp2, emp3}
	fmt.Println(totalExpense(employees[0], employees[1], employees[2]))

}
