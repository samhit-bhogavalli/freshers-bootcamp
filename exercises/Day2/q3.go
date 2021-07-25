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
	balance int
	mutex   sync.Mutex
}

//function for depositing the amount (this function will work without wait group)
func (ba *bankAccount) deposit(amount int) {
	ba.mutex.Lock()
	ba.balance += amount
	fmt.Println(strconv.Itoa(amount) + " is deposited total is " + strconv.Itoa(ba.balance))
	ba.mutex.Unlock()
}

//function for withdrawing the amount (this function will work without wait group)
func (ba *bankAccount) withdraw(id int, amount int) error {
	ba.mutex.Lock()
	if ba.balance >= amount {
		ba.balance -= amount
		fmt.Println(strconv.Itoa(amount) + " is withdrew total is " + strconv.Itoa(ba.balance))
		ba.mutex.Unlock()
		return nil
	} else {
		bal := ba.balance
		ba.mutex.Unlock()
		return errors.New("withdraw " + strconv.Itoa(id) + " failed, balance: " + strconv.Itoa(bal) + " withdraw requested: " + strconv.Itoa(amount))
	}
}

//function for displaying the error
func displayError(errors chan error, done chan bool) {
	for error := range errors {
		fmt.Println(error)
	}
	done <- true
}

//wrapper for using waitGroup
func (ba *bankAccount) userDeposit(amount int, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	ba.deposit(amount)
}

//wrapper for using waitGroup
func (ba *bankAccount) userWithdraw(id int, amount int, er chan error, waitGroup *sync.WaitGroup) {
	defer waitGroup.Done()
	if err := ba.withdraw(id, amount); err != nil {
		er <- err
	}
}

func main() {
	//testing
	var waitGroup sync.WaitGroup
	ba := bankAccount{500, sync.Mutex{}}
	er := make(chan error, 5)
	fmt.Println(ba.balance)
	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go ba.userDeposit(rand.Intn(1000), &waitGroup)
	}

	for i := 0; i < 5; i++ {
		waitGroup.Add(1)
		go ba.userWithdraw(i, rand.Intn(1000), er, &waitGroup)
	}
	done := make(chan bool)
	go displayError(er, done)
	waitGroup.Wait()
	close(er)
	<-done

}
