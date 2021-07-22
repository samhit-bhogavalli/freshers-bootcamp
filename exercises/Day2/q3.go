package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

//structure for bank account
type bankAccount struct {
	balance   int
	mutex     sync.Mutex
	waitGroup sync.WaitGroup
}

//function for depositing the amount
func (ba *bankAccount) deposit(amount int) {
	ba.mutex.Lock()
	ba.balance += amount
	fmt.Println(strconv.Itoa(amount) + " is deposited total is " + strconv.Itoa(ba.balance))
	ba.mutex.Unlock()
	ba.waitGroup.Done()
}

//function for withdrawing the amount
func (ba *bankAccount) withdraw(id int, amount int, er chan error) {
	ba.mutex.Lock()
	if ba.balance >= amount {
		ba.balance -= amount
		fmt.Println(strconv.Itoa(amount) + " is withdrew total is " + strconv.Itoa(ba.balance))
	} else {
		er <- errors.New("withdraw " + strconv.Itoa(id) + " failed, balance: " + strconv.Itoa(ba.balance) + " withdraw requested: " + strconv.Itoa(amount))
	}
	ba.mutex.Unlock()
	ba.waitGroup.Done()
}

//function for displaying the error
func displayError(errors chan error, done chan bool) {
	for error := range errors {
		fmt.Println(error)
	}
	done <- true
}

//function for initialising deposit routine
func (ba *bankAccount) userDeposit(amount int) {
	ba.waitGroup.Add(1)
	go ba.deposit(amount)
}

//function for initialising withdraw routine
func (ba *bankAccount) userWithdraw(id int, amount int, er chan error) {
	ba.waitGroup.Add(1)
	go ba.withdraw(id, amount, er)
}

func main() {
	//testing
	ba := bankAccount{500, sync.Mutex{}, sync.WaitGroup{}}
	er := make(chan error, 5)
	fmt.Println(ba.balance)
	for i := 0; i < 5; i++ {
		ba.userDeposit(rand.Intn(1000))
	}

	for i := 0; i < 5; i++ {
		ba.userWithdraw(i, rand.Intn(1000), er)
	}
	done := make(chan bool)
	go displayError(er, done)
	ba.waitGroup.Wait()
	close(er)
	<-done

}
