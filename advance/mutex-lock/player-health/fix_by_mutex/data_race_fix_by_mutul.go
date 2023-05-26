package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type Player struct {
	mutex3 sync.RWMutex
	health int
}

func NewPlayer() *Player {
	return &Player{health: 100}
}

// 提取方法 重构， do it in a clean way
func (p *Player) getHealth() int {
	p.mutex3.RLock()
	defer p.mutex3.RUnlock()
	return p.health
}

// 提取方法 重构， do it in a clean way
func (p *Player) takeDamage(value int) {
	p.mutex3.Lock()
	defer p.mutex3.Unlock()
	p.health -= value
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
