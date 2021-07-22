package main

import (
	"fmt"
	"sync"
)

type counter struct {
	frequency map[string]int
	mutex     sync.Mutex
	workers   int
	waitGroup sync.WaitGroup
}

//function for counting alphabets in a job
func (cnt *counter) worker(jobs chan string) {
	for job := range jobs {
		for _, rune := range job {
			cnt.mutex.Lock()
			cnt.frequency[string(rune)]++
			cnt.mutex.Unlock()
		}
	}
	defer cnt.waitGroup.Done()
}

//a function which takes slice of strings and calculates frequency
func (cnt *counter) getCount(input []string) {
	jobs := make(chan string, 3)
	for i := 0; i < cnt.workers; i++ {
		cnt.waitGroup.Add(1)
		go cnt.worker(jobs)
	}
	for _, val := range input {
		jobs <- val
	}
	close(jobs)
	cnt.waitGroup.Wait()
}

func main() {
	//testing
	input := []string{"quick", "quick", "quick", "quick", "quick", "quick"}
	frequencyCounter := counter{make(map[string]int), sync.Mutex{}, 3, sync.WaitGroup{}}
	frequencyCounter.getCount(input)
	fmt.Println(frequencyCounter.frequency)
}
