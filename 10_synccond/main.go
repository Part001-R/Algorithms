package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	pokemonList = []string{"Pikachu", "Charmander", "Squirtle", "Bulbasaur", "Jigglypuff"}
	cond        sync.Cond
	pokemon     string
	stop        bool
	wg          sync.WaitGroup
)

func main() {
	cond.L = &sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		for !stop {
			for pokemon != "Pikachu" {
				cond.Wait()
			}
			fmt.Println("found Pikachu!")
			stop = true
			cond.L.Unlock()
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 100)
			cond.L.Lock()
			pokemon = pokemonList[rand.Intn(len(pokemonList))]
			cond.L.Unlock()
			cond.Signal()
		}
		stop = true
		//cond.Broadcast()
	}()

	wg.Wait()
}
