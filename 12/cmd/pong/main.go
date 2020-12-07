package main

import "github.com/romanserikov/thinknetica-go/12/pkg/game"

func main() {
	pong := game.New(5)
	pong.Start()
}
