package engine

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/ttf"
)

func (engine *Engine) initializeFont() {
	fmt.Printf("initializing font\n")

	if err := ttf.Init(); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize TTF: %s\n", err)
		os.Exit(1)
	}

	if engine.font, err = ttf.OpenFont("../assets/silver.ttf", 32); err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open font: %s\n", err)
		os.Exit(4)
	}

	defer engine.font.Close()
}
