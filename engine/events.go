package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// HandleEvents loops through and handles events polled from SDL.
func HandleEvents(_event sdl.Event) {
	switch t := _event.(type) {
	case *sdl.QuitEvent:
		fmt.Println("Quitting Engine...")
		gameRunning = false
		break
	case *sdl.KeyboardEvent:
		HandleKeyboardEvent(t)
		break
	}
}

func HandleKeyboardEvent(_event *sdl.KeyboardEvent) {
	switch _event.Type {
	case sdl.KEYDOWN:
		if gameEngine.keyboardIsDownFunction != nil {
			gameEngine.keyboardIsDownFunction(int(_event.Keysym.Scancode))
		}
	}
}
