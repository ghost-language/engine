package main

import (
	"fmt"

	"ghostlang.org/engine/engine"
)

// Entity defines a new object that interacts with the game world.
type Entity struct {
	sprite *engine.Image
	x      int
	y      int
}

var (
	player *Entity
)

func load() {
	player = &Entity{}
	player.sprite = engine.NewImage("player.png")
}

func update(dt uint32) {
	player.x++
	player.y++
	fmt.Printf("x: %d, y: %d, dt: %d\n", player.x, player.y, dt)
}

func draw() {
	player.sprite.Draw(player.x, player.y)
}

func main() {
	engine := engine.NewEngine("Simple Example")
	engine.SetWindow(800, 600)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)

	engine.Run()
}
