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
		gameRunning = false
		break
	}
}

func HandleKeyboardEvents() {
	// internal bool
	// ENGINE_getKeyState(ENGINE* engine, char* keyName) {
	// 	    SDL_Keycode keycode =  SDL_GetKeyFromName(keyName);
	// 	    SDL_Scancode scancode = SDL_GetScancodeFromKey(keycode);
	// 	    uint8_t* state = SDL_GetKeyboardState(NULL);
	// 	    return state[scancode];
	// }
	if gameEngine.keyboardIsDownFunction != nil {
		gameEngine.keyboardIsDownFunction(sdl.GetKeyboardState())
	}
}
