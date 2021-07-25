package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//structure for teacher
type teacher struct {
	totalRating int
}

//function for asking student for rating
func (t *teacher) askStudent(waitGroup *sync.WaitGroup, ratings chan int) {
	defer waitGroup.Done()
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1000)))
	ratings <- rand.Intn(1000)
}

//function for reading answer from student
func (t *teacher) listenStudent(done chan bool, ratings chan int) {
	for rating := range ratings {
		t.totalRating += rating
	}
	done <- true
}

//function to calculate total rating
func (t *teacher) calculateRating(noOfStudents int) {
	var waitGroup sync.WaitGroup
	ratings := make(chan int)
	for i := 0; i < noOfStudents; i++ {
		waitGroup.Add(1)
		go t.askStudent(&waitGroup, ratings)
	}
	done := make(chan bool)
	go t.listenStudent(done, ratings)
	waitGroup.Wait()
	close(ratings)
	<-done
}

func main() {

	sir := teacher{}
	sir.calculateRating(200)
	fmt.Println(sir.totalRating / 200)
}
