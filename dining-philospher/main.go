package main

import (
	"fmt"
	"sync"
)

type Chops struct{ sync.Mutex }

type Philo struct {
	LeftCs  *Chops
	RightCs *Chops
}

func (p *Philo) eat(philoNums int, wg *sync.WaitGroup) {
	for i := 0; i < 3; i++ {
		p.LeftCs.Lock()
		p.RightCs.Lock()
		fmt.Println("starting to eat", philoNums)
		p.LeftCs.Unlock()
		p.RightCs.Unlock()
		fmt.Println("finishing to eat", philoNums)
	}
	wg.Done()
}

func host(Philos []*Philo, wg *sync.WaitGroup) {
	/*
		there is a session, that two philo start eating at same time, and less waiting.
	*/
	philoEatCounter := 0     // philo eat counter indentify philo eat at same time.
	eatingSession := 0       // count the eating session
	for i := 0; i < 5; i++ { // 5 philos start eating
		wg.Add(1)               // add a wait group to identify running goroutine at same time
		go Philos[i].eat(i, wg) // call eat method from philos
		philoEatCounter++

		// if philos is even then wait the time until all two philos finished eating
		if philoEatCounter%2 == 0 {
			eatingSession++
			fmt.Println("====== Eating session", eatingSession, "======")
			wg.Wait()
		}

		// last philosoper
		if 5-philoEatCounter == 0 {
			eatingSession++
			fmt.Println("====== Eating session", eatingSession, "======")
			wg.Wait()
		}
	}
}

func main() {
	// populate chopsticks
	var wg sync.WaitGroup
	CSticks := make([]*Chops, 5)
	for i := 0; i < 5; i++ {
		CSticks[i] = new(Chops)
	}

	// populate philospers
	Philos := make([]*Philo, 5)
	for i := 0; i < 5; i++ {
		Philos[i] = &Philo{
			CSticks[i],
			CSticks[(i+1)%5],
		}
	}

	// call host
	host(Philos, &wg)
}
