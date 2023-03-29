package main

import (
	"fmt"
	"time"
)

type number struct {
	a int
	b int
}

type result struct {
	sum int
}

func worker(workId int, jobs <-chan number, results chan<- result) {
	for job := range jobs {
		fmt.Printf("work_id: %d\n", workId+1)
		results <- result{sum(job)}
		time.Sleep(2 * time.Second)
	}
}

func sum(num number) int {
	return num.a + num.b
}

func main() {
	nums := []number{
		{a: 1, b: 2},
		{a: 3, b: 4},
		{a: 5, b: 6},
		{a: 7, b: 8},
		{a: 9, b: 10},
		{a: 11, b: 12},
		{a: 13, b: 14},
		{a: 15, b: 16},
		{a: 17, b: 18},
		{a: 19, b: 20},
	}

	jobsCh := make(chan number, len(nums))
	resultsCh := make(chan result, len(nums))

	for _, j := range nums {
		jobsCh <- j
	}
	close(jobsCh)

	numWorkers := 2
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobsCh, resultsCh)
	}

	for r := 0; r < len(nums); r++ {
		result := <-resultsCh
		fmt.Println(result)
	}
}
