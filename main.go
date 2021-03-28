package main

import (
	"fmt"
	"sync"
	"time"
)

type ChopSticks struct {
	sno   int
	inuse sync.Mutex
}

type Philo struct {
	sno             int
	left, right     *ChopSticks
	times           int
	responseChannel chan bool
}

type Host struct {
	activeEaterCount int
	philoDoneEating  int
}

func (p Philo) eat() {
	fmt.Printf("\nstarting to eat %d", p.sno)
	fmt.Printf("\nfinishing eating %d", p.sno)
}

const noOfPhilo = 5
const noOfTimes = 3

func PhiloRoutine(wg *sync.WaitGroup, p *Philo, requestChannel chan Philo, releaseChannel chan Philo) {
	defer wg.Done()
	for {
		select {
		case allowedToEat := <-p.responseChannel:
			if allowedToEat {
				p.eat()
				p.times++
				releaseChannel <- *p
				if p.times == noOfTimes {
					fmt.Printf("\n ### Philosopher  %d [ate 3 times] ", p.sno)
					return
				}
			} else {
				fmt.Printf("\nEat request not granted %d", p.sno)
			}
		case requestChannel <- *p:
			//fmt.Println("\nEat request submitted ", p.sno)
		default:
			time.Sleep(20 * time.Millisecond)
		}

	}
}

func hostRequestRoutine(wg *sync.WaitGroup, host *Host, requestChannel chan Philo) {
	wg.Done()
	for {
		p := <-requestChannel
		if host.activeEaterCount < 2 {
			p.left.inuse.Lock()
			p.right.inuse.Lock()
			p.responseChannel <- true
			host.activeEaterCount++
		} else {
			fmt.Println("*** 2 eaters already active *** ")
			p.responseChannel <- false
		}
	}
}

func hostResponseRoutine(wg *sync.WaitGroup, host *Host, requestChannel chan Philo, releaseChannel chan Philo) {
	wg.Done()
	for {
		p := <-releaseChannel
		p.left.inuse.Unlock()
		p.right.inuse.Unlock()
		host.activeEaterCount--
		if p.times == noOfTimes {
			host.philoDoneEating++
			if host.philoDoneEating == noOfPhilo {
				close(releaseChannel)
				close(requestChannel)
				return
			}
		}
	}
}

func main() {
	wg := &sync.WaitGroup{}
	requestChannel := make(chan Philo)
	releaseChannel := make(chan Philo)
	h := new(Host)

	wg.Add(1)
	go hostRequestRoutine(wg, h, requestChannel)

	wg.Add(1)
	go hostResponseRoutine(wg, h, requestChannel, releaseChannel)

	var chopStix = [5]ChopSticks{}
	for i := 0; i < 5; i++ {
		chopStix[i] = ChopSticks{sno: i}
	}

	for i := 0; i < 5; i++ {
		p := new(Philo)
		p.responseChannel = make(chan bool)
		p.sno = i
		p.left = &chopStix[i]
		p.right = &chopStix[(i+1)%5]
		p.times = 0
		wg.Add(1)
		go PhiloRoutine(wg, p, requestChannel, releaseChannel)
	}
	wg.Wait()
}
