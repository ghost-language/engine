package engine

import (
	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
	"github.com/veandco/go-sdl2/sdl"
)

// MouseHideCursorFunction registers the Mouse.hideCursor function with Ghost.
func (engine *Engine) MouseHideCursorFunction(env *object.Environment, args ...object.Object) object.Object {
	sdl.ShowCursor(sdl.DISABLE)

	return ghost.NULL
}

// MouseShowCursorFunction registers the Mouse.showCursor function with Ghost.
func (engine *Engine) MouseShowCursorFunction(env *object.Environment, args ...object.Object) object.Object {
	sdl.ShowCursor(sdl.ENABLE)

	return ghost.NULL
}
