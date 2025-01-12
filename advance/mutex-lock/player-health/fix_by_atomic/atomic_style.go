package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type Player struct {
	health int32
}

func NewPlayer() *Player {
	return &Player{health: 100}
}

func (p *Player) getHealth() int {

	return int(atomic.LoadInt32(&p.health))
}

func (p *Player) takeDamage(value int) {
	health := p.getHealth()
	atomic.StoreInt32(&p.health, int32(health-value))
}

func startUILoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		fmt.Printf("player health: %d\n", p.getHealth())
		<-ticker.C
	}
}

func startGameLoop(p *Player) {
	ticker := time.NewTicker(time.Second)
	for {
		p.takeDamage(rand.Intn(30))
		if p.getHealth() <= 0 {
			fmt.Printf("GAME OVER")
			break
		}
		<-ticker.C
	}
}

func main() {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}
