package game

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Game -
type Game struct {
	ch       chan string
	wg       *sync.WaitGroup
	scores   map[string]int
	winscore int
}

// New - creates new game with winscore
func New(winscore int) *Game {
	return &Game{
		ch:       make(chan string),
		scores:   make(map[string]int),
		winscore: winscore,
		wg:       new(sync.WaitGroup),
	}
}

// Start - starts the game
func (g *Game) Start() {
	fmt.Println("Ping-pong starts")

	rand.Seed(time.Now().Unix())

	g.wg.Add(2)

	go g.player("Alice")
	go g.player("Bob")

	g.ch <- "begin"

	g.wg.Wait()

	fmt.Println("Scores", g.scores)
}

func (g *Game) player(name string) {
	defer g.wg.Done()

	for val := range g.ch {
		if g.haveWinner() {
			close(g.ch)
			return
		}

		var turn string
		switch val {
		case "begin", "stop", "pong":
			turn = "ping"
		case "ping":
			turn = "pong"
		}

		fmt.Println(name, turn)

		if isLucky() {
			g.scores[name]++
			fmt.Println(name, "scored!", g.scores)
			g.ch <- "stop"
		} else {
			g.ch <- turn
		}
	}
}

func (g *Game) haveWinner() bool {
	for _, score := range g.scores {
		if score == g.winscore {
			return true
		}
	}

	return false
}

func isLucky() bool {
	return rand.Intn(5) == 0
}
