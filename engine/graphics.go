package engine

import (
	"strconv"

	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
)

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
