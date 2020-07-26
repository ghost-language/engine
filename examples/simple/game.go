package main

import (
	"fmt"

	"ghostlang.org/x/engine/engine"
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

func keypressed(state []uint8) {
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

	if state[sdl.SCANCODE_SPACE] == 1 {
		fmt.Println("Fire!")
	}
}

func main() {
	engine := engine.NewEngine("Simple Example")
	engine.SetWindow(800, 600)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)
	engine.SetKeyPressedFunction(keypressed)

	engine.Run()
}
