package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type ChopStick struct {
	sync.Mutex
}

type Philosopher struct {
	lChopStick *ChopStick
	rChopStick *ChopStick
	idx        int
}

func (p Philosopher) eat(wg *sync.WaitGroup, eating chan bool) {
	defer wg.Done()

	for i := 0; i < 3; i++ {
		/* Waiting if 2 philosophers are
		   already eating at the moment */

		eating <- true

		/* The philosopher takes two
		   sticks and starts eating */

		p.lChopStick.Lock()
		p.rChopStick.Lock()

		/* The philosopher eats at random time */

		time.Sleep(time.Millisecond * time.Duration(rand.Intn(2000)))
		fmt.Printf("%d philosopher starting to eat\n", p.idx)

		/* The philosopher puts down two
		   chopsticks and finishes eating */

		p.rChopStick.Unlock()
		p.lChopStick.Unlock()

		fmt.Printf("%d philosopher finishing eating\n", p.idx)

		/* The philosopher notifies the queue
		   channel that he has finished eating */

		<-eating
	}
}

func main() {
	var wg sync.WaitGroup

	/* The queue channel for philosophers.
	   Only 2 philosophers can eat at the same time	*/

	var eatingChannel = make(chan bool, 2)

	/* Generate 5 chop sticks for philosophers */

	sticks := make([]*ChopStick, 5)
	for i := 0; i < 5; i++ {
		sticks[i] = new(ChopStick)
	}

	fmt.Printf("Philosophers began to eat!\n\n")

	/* Generate 5 philosophers who start eating */

	philosophers := make([]*Philosopher, 5)
	for i := 0; i < 5; i++ {
		/* Creation of a philosopher */

		philosophers[i] = &Philosopher{
			lChopStick: sticks[i],
			rChopStick: sticks[(i+1)%5],
			idx:        i,
		}

		/* The philosopher begins to eat */

		wg.Add(1)
		go philosophers[i].eat(&wg, eatingChannel)
	}

	wg.Wait()
	fmt.Printf("\nAll the philosophers have eaten!\n")
}
