package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	title        = "Engine"
	screenWidth  = 800
	screenHeight = 600
)

var (
	err           error
	window        *sdl.Window
	renderer      *sdl.Renderer
	event         sdl.Event
	running       bool
	ticker        *time.Ticker
	screenSurface *sdl.Surface
)

func load() {
	//
}

func update() {
	//
}

func draw() {
	renderer.SetDrawColor(247, 250, 252, 255)
	renderer.Clear()

	renderer.Present()
}

func main() {
	if !initialize() {
		os.Exit(1)
	}

	running = true
	ticker = time.NewTicker(time.Second / 30)

	load()

	for running {
		handleEvents()
		update()
		draw()
		<-ticker.C
	}

	shutdown()
}

func initialize() bool {
	err := sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		fmt.Println("initializing SDL:", err)
		return false
	}

	w, err := sdl.CreateWindow(
		title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		screenWidth,
		screenHeight,
		sdl.WINDOW_OPENGL)

	if err != nil {
		fmt.Println("initializing window:", err)
		return false
	}

	window = w

	r, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		fmt.Println("initializing renderer:", err)
		return false
	}

	renderer = r

	return true
}

func shutdown() {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

func handleEvents() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			println("Quitting Engine...")
			running = false
			break
		}
	}
}
