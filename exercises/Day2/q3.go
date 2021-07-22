package main

import (
	"errors"
	"fmt"
	"math/rand"
	"strconv"
	"sync"
)

var balance = 200

//function for depositing the amount
func deposit(amount int, mutex *sync.Mutex, waitGroup *sync.WaitGroup) {
	mutex.Lock()
	balance += amount
	fmt.Println(strconv.Itoa(amount) + " is deposited")
	mutex.Unlock()
	waitGroup.Done()
}

//function for withdrawing the amount
func withdraw(id int, amount int, er chan error, mutex *sync.Mutex, waitGroup *sync.WaitGroup) {
	mutex.Lock()
	if balance >= amount {
		balance -= amount
		fmt.Println(strconv.Itoa(amount) + " is withdrew")
	} else {
		er <- errors.New("withdraw " + strconv.Itoa(id) + " failed, balance: " + strconv.Itoa(balance) + " withdraw requested: " + strconv.Itoa(amount))
	}
	mutex.Unlock()
	waitGroup.Done()
}

//function for displaying the error
func displayError(errors chan error, done chan bool) {
	for error := range errors {
		fmt.Println(error)
	}
	done <- true
}

func main() {
	er := make(chan error, 5)
	var waitGroup sync.WaitGroup
	var mutex sync.Mutex
	fmt.Println(balance)
	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go deposit(rand.Intn(1000), &mutex, &waitGroup)
	}

	for i := 0; i < 100; i++ {
		waitGroup.Add(1)
		go withdraw(i, rand.Intn(1000), er, &mutex, &waitGroup)
	}
	done := make(chan bool)
	go displayError(er, done)
	waitGroup.Wait()
	close(er)
	<-done

}
