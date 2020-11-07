package engine

import (
	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
	"github.com/shopspring/decimal"
	"github.com/veandco/go-sdl2/sdl"
)

// WindowBorderedFunction registers the Window.bordered function with Ghost.
func (engine *Engine) WindowBorderedFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=1", len(args))
	}

	engine.window.SetBordered(args[0].(*object.Boolean).Value)

	return ghost.NULL
}

// WindowFullscreenFunction registers the Window.fullscreen function with Ghost.
func (engine *Engine) WindowFullscreenFunction(env *object.Environment, args ...object.Object) object.Object {
	engine.window.SetFullscreen(sdl.WINDOW_FULLSCREEN)

	return ghost.NULL
}

// WindowWidthFunction registers the Window.width function with Ghost.
func (engine *Engine) WindowWidthFunction(env *object.Environment, args ...object.Object) object.Object {
	w, _ := engine.window.GetSize()

	return &object.Number{Value: decimal.NewFromInt(int64(w))}
}

// WindowHeightFunction registers the Window.height function with Ghost.
func (engine *Engine) WindowHeightFunction(env *object.Environment, args ...object.Object) object.Object {
	_, h := engine.window.GetSize()

	return &object.Number{Value: decimal.NewFromInt(int64(h))}
}

// WindowTitleFunction registers the Window.title function with Ghost.
func (engine *Engine) WindowTitleFunction(env *object.Environment, args ...object.Object) object.Object {
	if len(args) != 1 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=1", len(args))
	}

	engine.window.SetTitle(args[0].Inspect())

	return ghost.NULL
}
