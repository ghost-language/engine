package engine

import (
	"strconv"

	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
	"github.com/veandco/go-sdl2/sdl"
)

// GraphicsClearFunction registers the Graphics.clear function with Ghost.
func (engine *Engine) GraphicsClearFunction(env *object.Environment, args ...object.Object) object.Object {
	engine.renderer.Clear()

	return ghost.NULL
}

// GraphicsDrawFunction registers the Graphics.draw function with Ghost.
func (engine *Engine) GraphicsDrawFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 3 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=3", len(args))
	}

	file := args[0].Inspect()
	x, _ := strconv.ParseInt(args[1].Inspect(), 10, 64)
	y, _ := strconv.ParseInt(args[2].Inspect(), 10, 64)

	image := NewImage(file)

	image.Draw(int(x), int(y))

	return ghost.NULL
}

// GraphicsFilledRectangleFunction registers the Graphics.filledRectangle function with Ghost.
func (engine *Engine) GraphicsFilledRectangleFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 4 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=4", len(args))
	}

	x := int32(args[0].(*object.Number).Value.IntPart())
	y := int32(args[1].(*object.Number).Value.IntPart())
	w := int32(args[2].(*object.Number).Value.IntPart())
	h := int32(args[3].(*object.Number).Value.IntPart())

	rectangle := sdl.Rect{X: x, Y: y, W: w, H: h}

	engine.renderer.FillRect(&rectangle)

	return ghost.NULL
}

// GraphicsLineFunction registers the Graphics.line function with Ghost.
func (engine *Engine) GraphicsLineFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 4 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=4", len(args))
	}

	x1 := int32(args[0].(*object.Number).Value.IntPart())
	y1 := int32(args[1].(*object.Number).Value.IntPart())
	x2 := int32(args[2].(*object.Number).Value.IntPart())
	y2 := int32(args[3].(*object.Number).Value.IntPart())
	// width := 2

	engine.renderer.DrawLine(x1, y1, x2, y2)

	return ghost.NULL
}

// GraphicsPixelFunction registers the Graphics.pixel function with Ghost.
func (engine *Engine) GraphicsPixelFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 2 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=2", len(args))
	}

	engine.renderer.DrawPoint(400, 300)

	return ghost.NULL
}

// GraphicsPrintFunction registers the Graphics.print function with Ghost.
func (engine *Engine) GraphicsPrintFunction(env *object.Environment, args ...object.Object) object.Object {
	color := sdl.Color{R: 0, G: 0, B: 0, A: 255}
	solid, err := engine.font.RenderUTF8Solid("Hello, universe!", color)

	if err != nil {
		return ghost.NewError("failed to render text: %s", err)
	}

	defer solid.Free()

	surface, err := engine.window.GetSurface()

	if err != nil {
		return ghost.NewError("failed to get window surface: %s", err)
	}

	err = solid.Blit(nil, surface, nil)

	if err != nil {
		return ghost.NewError("failed to put text on window surface: %s", err)
	}

	return ghost.NULL
}

// GraphicsRectangleFunction registers the Graphics.rectangle function with Ghost.
func (engine *Engine) GraphicsRectangleFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 4 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=4", len(args))
	}

	x := int32(args[0].(*object.Number).Value.IntPart())
	y := int32(args[1].(*object.Number).Value.IntPart())
	w := int32(args[2].(*object.Number).Value.IntPart())
	h := int32(args[3].(*object.Number).Value.IntPart())

	rectangle := sdl.Rect{X: x, Y: y, W: w, H: h}

	engine.renderer.DrawRect(&rectangle)

	return ghost.NULL
}

func (engine *Engine) GraphicsSetColorFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 3 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=3", len(args))
	}

	red := uint8(args[0].(*object.Number).Value.IntPart())
	green := uint8(args[1].(*object.Number).Value.IntPart())
	blue := uint8(args[2].(*object.Number).Value.IntPart())
	alpha := uint8(255)

	if len(args) == 4 {
		alpha = uint8(args[3].(*object.Number).Value.IntPart())
	}

	engine.renderer.SetDrawColor(red, green, blue, alpha)

	return ghost.NULL
}
