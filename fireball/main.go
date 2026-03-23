package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type BallFire struct {
	heats      int
	lastPlayer string
}

var id int

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	player := []string{"Player 1", "Player 2", "Player 3", "Player 4", "Player 5"}

	ballfire := make(chan BallFire, len(player))
	done := make(chan BallFire, len(player))

	rand.Seed(time.Now().UnixNano())
	var wg sync.WaitGroup
	var mx sync.Mutex
	ballfire <- BallFire{}
	for i, v := range player {
		wg.Add(1)
		go Play(v, ballfire, done, &wg, i, &mx, ctx)
	}
	if v, ok := <-done; ok {
		cancel()
		fmt.Println("🔥 bola terlalu panas")
		fmt.Printf("%s kalah (heats: %d)\n", v.lastPlayer, v.heats)
	}
	wg.Wait()
}

func Play(name string, ballF, done chan BallFire, wg *sync.WaitGroup, index int, mx *sync.Mutex, ctx context.Context) {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		default:
			mx.Lock()
			if id > 4 {
				id = 0
			}

			if id != index {
				mx.Unlock()
				break
			}

			h := <-ballF

			if h.lastPlayer != "" {
				fmt.Printf("%s mengoper bola ke %s\n", h.lastPlayer, name)
			}

			v := rand.Intn(5) + 1
			h.heats += v
			h.lastPlayer = name
			time.Sleep(time.Second)
			if h.heats > 100 {
				done <- h
				mx.Unlock()
				return
			}

			fmt.Printf("%s menerima bola (heats: %d)\n", h.lastPlayer, h.heats)
			ballF <- h
			id += 1
			mx.Unlock()

		}
	}

}
