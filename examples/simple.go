package main

import (
	"ghostlang.org/engine/engine"
)

var (
	player *engine.Image
)

func load() {
	player = engine.NewImage("player.png")
}

func update(dt uint32) {
	//
}

func draw() {
	player.Draw(200, 200)
}

func main() {
	engine := engine.NewEngine("Simple Example")
	engine.SetWindow(800, 600)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)

	engine.Run()
}
