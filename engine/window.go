package engine

import (
	"ghostlang.org/x/ghost/object"
	"github.com/shopspring/decimal"
)

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
