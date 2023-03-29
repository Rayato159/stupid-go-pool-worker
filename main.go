package main

import (
	"fmt"
	"log"
	"time"
)

type number struct {
	a int
	b int
}

type result struct {
	sum int
}

func worker(workId int, jobsCh <-chan number, resultsCh chan<- result, errsCh chan<- error) {
	for job := range jobsCh {
		fmt.Printf("work_id: %d\n", workId+1)
		sumResult, err := sum(job)
		if err != nil {
			errsCh <- err
			return
		}
		// ห้ามลืม!!! ยัด nil ลง errsCh ที่มันว่างทกครั้ง ไม่งั้นมันจะค้างงงง
		errsCh <- nil
		resultsCh <- result{sumResult}
		time.Sleep(2 * time.Second)
	}
}

func sum(num number) (int, error) {
	if num.a+num.b == 7 {
		return -1, fmt.Errorf("error!")
	}
	return num.a + num.b, nil
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
	errsCh := make(chan error, len(nums))

	for _, j := range nums {
		jobsCh <- j
	}
	close(jobsCh)

	numWorkers := 2
	for w := 0; w < numWorkers; w++ {
		go worker(w, jobsCh, resultsCh, errsCh)
	}

	for r := 0; r < len(nums); r++ {
		err := <-errsCh
		if err != nil {
			log.Fatal(err)
		}
		result := <-resultsCh
		fmt.Println(result)
	}
}
