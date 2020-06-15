package engine

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// HandleEvents loops through and handles events polled from SDL.
func HandleEvents(_event sdl.Event) {
	switch _event.(type) {
	case *sdl.QuitEvent:
		fmt.Println("Quitting Engine...")
		running = false
	}
}
