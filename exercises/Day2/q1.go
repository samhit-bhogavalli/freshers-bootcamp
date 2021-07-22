package main

import (
	"fmt"
	"sync"
)

//function for counting alphabets in a job
func countAlphabets(jobs chan string, done chan bool, count map[string]int, mutex *sync.Mutex) {
	for job := range jobs {
		for _, rune := range job {
			mutex.Lock()
			count[string(rune)]++
			mutex.Unlock()
		}
	}
	done <- true
}

func main() {
	//testing
	input := []string{"quick", "quick", "quick", "quick", "quick", "quick"}
	jobs := make(chan string, 5)
	done := make(chan bool)
	count := make(map[string]int)
	var mutex sync.Mutex
	go countAlphabets(jobs, done, count, &mutex)
	for _, val := range input {
		jobs <- val
	}
	close(jobs)
	<-done
	fmt.Println(count)
}
