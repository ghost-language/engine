package main

import (
	"ghostlang.org/engine/engine"
)

func load() {
	//
}

func update(dt uint32) {
	//
}

func draw() {
	//
}

func main() {
	engine := engine.NewEngine("Simple Example")
	engine.SetWindow(800, 600)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)

	engine.Run()
}
