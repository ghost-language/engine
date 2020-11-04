package engine

import (
	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
	"ghostlang.org/x/ghost/value"
	"github.com/veandco/go-sdl2/sdl"
)

// KeyboardIsDownFunction registers the Keyboard.isDown function with Ghost.
func (engine *Engine) KeyboardIsDownFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=1", len(args))
	}

	scancode := sdl.GetScancodeFromName(args[0].Inspect())

	if engine.KeyState[scancode] != 0 {
		return value.TRUE
	}

	return value.FALSE
}
