package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	health int
}

func NewPlayer() *Player {
	return &Player{health: 100}
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("player health: %d\n", p.health)
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		p.health -= rand.Intn(30)
		if p.health <= 0 {
			fmt.Printf("GAME OVER")
			break
		}
		<-ticker.C
	}
}

// main函数，不方便检测 data races 溢出 go test ./... --race
func main() {
}
