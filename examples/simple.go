package main

import (
	"ghostlang.org/engine/engine"
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

func keyboardIsDown(scancode int) {
	switch scancode {
	case 79:
		player.x += player.speed
	case 81:
		player.y += player.speed
	case 80:
		player.x -= player.speed
	case 82:
		player.y -= player.speed
	}
}

func main() {
	engine := engine.NewEngine("Simple Example")
	engine.SetWindow(800, 600)
	engine.SetFPS(120)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)
	engine.SetKeyboardIsDownFunction(keyboardIsDown)

	engine.Run()
}
