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
)

// Engine holds the bindings for SDL and all callback functions to
// be called during the main game loop.
type Engine struct {
	loadFunction           func()
	updateFunction         func(dt uint32)
	drawFunction           func()
	keyboardIsDownFunction func(scancode int)

	title  string
	width  int32
	height int32

	window   *sdl.Window
	renderer *sdl.Renderer
}

// NewEngine initializes a new engine instance.
func NewEngine(_title string) (engine *Engine) {
	engine = new(Engine)

	engine.title = _title

	gameEngine = engine

	return engine
}

// SetWindow defines the windows width and height.
func (engine *Engine) SetWindow(_width int32, _height int32) {
	engine.width = _width
	engine.height = _height
}

// SetLoadFunction defines the load function to be used by Engine.
func (engine *Engine) SetLoadFunction(_load func()) {
	engine.loadFunction = _load
}

// SetUpdateFunction defines the update function to be used by Engine.
func (engine *Engine) SetUpdateFunction(_update func(dt uint32)) {
	engine.updateFunction = _update
}

// SetDrawFunction defines the draw function to be used by Engine.
func (engine *Engine) SetDrawFunction(_draw func()) {
	engine.drawFunction = _draw
}

// SetKeyboardIsDownFunction defines the keydown function to be used by Engine.
func (engine *Engine) SetKeyboardIsDownFunction(_keyboardIsDown func(scancode int)) {
	engine.keyboardIsDownFunction = _keyboardIsDown
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

	engine.renderer, err = sdl.CreateRenderer(engine.window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		panic(fmt.Sprintf("Engine error: Could not create renderer - %s", err))
	}

	img.Init(img.INIT_JPG | img.INIT_PNG)

	if engine.loadFunction != nil {
		engine.loadFunction()
	}

	gameRunning = true
}

func (engine *Engine) update(dt uint32) {
	if engine.updateFunction != nil {
		engine.updateFunction(dt)
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

	var deltaTime, oldDeltaTime, newDeltaTime uint32 = 0, 0, 0

	for gameRunning {
		// Check for events and handle them
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			HandleEvents(event)
		}

		// Calculate delta time
		newDeltaTime = sdl.GetTicks()
		deltaTime = newDeltaTime - oldDeltaTime
		oldDeltaTime = newDeltaTime

		engine.update(deltaTime)
		engine.draw()

		// Give the CPU some time to run calculations
		sdl.Delay(1)
	}
}

// Exit frees all resources used by Engine for a clean exit.
func (engine *Engine) Exit() {
	freeResources()

	engine.renderer.Destroy()

	engine.window.Destroy()

	sdl.Quit()
}
