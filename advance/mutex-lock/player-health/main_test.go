package main

import "testing"

func TestGame(t *testing.T) {
	player := NewPlayer()
	go startUILoop(player)
	startGameLoop(player)
}

// go test ./... --race  failed here
