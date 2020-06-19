package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

var (
	gameEngine  *Engine
	gameRunning bool
	err         error
	frameDelay  uint32
)

// Engine holds the bindings for SDL and all callback functions to
// be called during the main game loop.
type Engine struct {
	loadFunction       func()
	updateFunction     func()
	drawFunction       func()
	keyPressedFunction func(state []uint8)

	title  string
	width  int32
	height int32
	fps    uint32

	window   *sdl.Window
	renderer *sdl.Renderer
}

// NewEngine initializes a new engine instance.
func NewEngine(_title string) (engine *Engine) {
	engine = new(Engine)

	engine.title = _title
	engine.fps = 60

	frameDelay = 1000 / engine.fps
	gameEngine = engine

	return engine
}

// SetWindow defines the windows width and height.
func (engine *Engine) SetWindow(_width int32, _height int32) {
	engine.width = _width
	engine.height = _height
}

// SetTitle defines the title of the window.
func (engine *Engine) SetTitle(_title string) {
	engine.title = _title
}

// SetFPS defines the desired frames per second threshold.
func (engine *Engine) SetFPS(_fps uint32) {
	engine.fps = _fps
	frameDelay = 1000 / engine.fps
}

// SetLoadFunction defines the load function to be used by Engine.
func (engine *Engine) SetLoadFunction(_load func()) {
	engine.loadFunction = _load
}

// SetUpdateFunction defines the update function to be used by Engine.
func (engine *Engine) SetUpdateFunction(_update func()) {
	engine.updateFunction = _update
}

// SetDrawFunction defines the draw function to be used by Engine.
func (engine *Engine) SetDrawFunction(_draw func()) {
	engine.drawFunction = _draw
}

// SetKeyPressedFunction defines the keydown function to be used by Engine.
func (engine *Engine) SetKeyPressedFunction(_keyPressed func(state []uint8)) {
	engine.keyPressedFunction = _keyPressed
}

func (engine *Engine) initialize() {
	err = sdl.Init(sdl.INIT_EVERYTHING)

	if err != nil {
		panic(fmt.Sprintf("Engine error: Could not initialize SDL - %s", err))
	}

	engine.window, err = sdl.CreateWindow(
		engine.title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		engine.width,
		engine.height,
		sdl.WINDOW_SHOWN)

	if err != nil {
		panic(fmt.Sprintf("Engine error: Could not create window - %s", err))
	}

	engine.renderer, err = sdl.CreateRenderer(engine.window, -1, sdl.RENDERER_ACCELERATED|sdl.RENDERER_PRESENTVSYNC)

	if err != nil {
		panic(fmt.Sprintf("Engine error: Could not create renderer - %s", err))
	}

	img.Init(img.INIT_JPG | img.INIT_PNG)

	if engine.loadFunction != nil {
		engine.loadFunction()
	}

	gameRunning = true
}

func (engine *Engine) update() {
	if engine.updateFunction != nil {
		engine.updateFunction()
	}
}

func (engine *Engine) draw() {
	engine.renderer.SetDrawColor(26, 32, 44, 255)
	engine.renderer.Clear()

	if engine.drawFunction != nil {
		engine.drawFunction()
	}

	engine.renderer.Present()
}

// Run the main game loop.
func (engine *Engine) Run() {
	defer engine.Exit()

	if engine.loadFunction == nil {
		fmt.Println("Engine warning: No load function present.")
	}

	if engine.updateFunction == nil {
		fmt.Println("Engine warning: No update function present.")
	}

	if engine.drawFunction == nil {
		fmt.Println("Engine warning: No draw function present.")
	}

	engine.initialize()

	for gameRunning {
		frameStart := sdl.GetTicks()

		// Check for events and handle them
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			HandleEvents(event)
		}

		HandleKeyboardEvents()

		engine.update()
		engine.draw()

		frameTime := sdl.GetTicks() - frameStart

		// Give the CPU some time to run calculations
		if frameDelay > frameTime {
			sdl.Delay(frameDelay - frameTime)
		}
	}
}

// Exit frees all resources used by Engine for a clean exit.
func (engine *Engine) Exit() {
	freeResources()

	engine.renderer.Destroy()

	engine.window.Destroy()

	sdl.Quit()
}
