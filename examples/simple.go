package main

import (
	"ghostlang.org/engine/engine"
	"github.com/veandco/go-sdl2/sdl"
)

// Entity defines a new object that interacts with the game world.
type Entity struct {
	sprite *engine.Image
	x      int
	y      int
	speed  int
}

var (
	player *Entity
)

func load() {
	player = &Entity{}
	player.sprite = engine.NewImage("player.png")
	player.speed = 5
}

func update() {
	//
}

func draw() {
	player.sprite.Draw(player.x, player.y)
}

func keyboardIsDown(state []uint8) {
	if state[sdl.SCANCODE_RIGHT] == 1 {
		player.x += player.speed
	}

	if state[sdl.SCANCODE_LEFT] == 1 {
		player.x -= player.speed
	}

	if state[sdl.SCANCODE_UP] == 1 {
		player.y -= player.speed
	}

	if state[sdl.SCANCODE_DOWN] == 1 {
		player.y += player.speed
	}
}

func contains(s []uint8, e uint8) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func main() {
	engine := engine.NewEngine("Simple Example")
	engine.SetWindow(800, 600)
	// engine.SetFPS(120)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)
	engine.SetKeyboardIsDownFunction(keyboardIsDown)

	engine.Run()
}
