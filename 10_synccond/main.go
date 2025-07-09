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
	stopFin     bool
	stopFound   bool
	wg          sync.WaitGroup
)

func main() {
	cond.L = &sync.Mutex{}

	wg.Add(1)
	go func() {
		defer wg.Done()
		cond.L.Lock()
		for pokemon != "Pikachu" && !stopFin {
			cond.Wait()
		}
		stopFound = true
		fmt.Println("found Pikachu!")
		cond.L.Unlock()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			time.Sleep(time.Millisecond * 100)
			cond.L.Lock()
			pokemon = pokemonList[rand.Intn(len(pokemonList))]
			cond.L.Unlock()

			if stopFound {
				return
			}
			cond.Signal()
		}
		stopFin = true
		//cond.Broadcast()
	}()

	wg.Wait()
}
