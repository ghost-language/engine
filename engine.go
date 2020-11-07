package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"

	"ghostlang.org/x/engine/engine"
	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
)

var (
	flagVersion bool
	flagHelp    bool
)

// Console holds the passed values through the command line.
type Console struct {
	args []string
}

func init() {
	flag.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), "Usage: %s [options] [<filename>]\n", path.Base(os.Args[0]))
		flag.PrintDefaults()
		os.Exit(0)
	}

	flag.BoolVar(&flagHelp, "h", false, "display help information")
	flag.BoolVar(&flagVersion, "v", false, "display version information")
}

func main() {
	flag.Parse()

	args := flag.Args()

	if flagVersion {
		fmt.Printf("%s %s\n", path.Base(os.Args[0]), "dev-nightly")
		os.Exit(0)
	}

	if flagHelp {
		showHelp()
		os.Exit(2)
	}

	console := Console{args}

	var f *os.File
	var err error

	if len(console.args) == 0 {
		// Do we have a main.ghost file present?

		ex, err := os.Executable()

		if err != nil {
			panic(err)
		}

		exPath := filepath.Dir(ex)
		fmt.Println(exPath)

		mainFile, _ := filepath.Abs(exPath + "/main.ghost")
		f, err = os.Open(mainFile)

		if err != nil {
			log.Fatalf("could not find main.ghost: %s", err)
			showHelp()
			os.Exit(2)
		}
	} else {
		f, err = os.Open(console.args[0])

		if err != nil {
			log.Fatalf("could not open source file %s: %s", console.args[0], err)
		}
	}

	b, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading source file: %s", err)
		return
	}

	engine := engine.NewEngine("Engine")
	engine.SetWindow(800, 600)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)

	// Graphics Functions
	ghost.RegisterFunction("Graphics.clear", engine.GraphicsClearFunction)
	ghost.RegisterFunction("Graphics.draw", engine.GraphicsDrawFunction)
	ghost.RegisterFunction("Graphics.filledRectangle", engine.GraphicsFilledRectangleFunction)
	ghost.RegisterFunction("Graphics.line", engine.GraphicsLineFunction)
	ghost.RegisterFunction("Graphics.pixel", engine.GraphicsPixelFunction)
	ghost.RegisterFunction("Graphics.print", engine.GraphicsPrintFunction)
	ghost.RegisterFunction("Graphics.rectangle", engine.GraphicsRectangleFunction)
	ghost.RegisterFunction("Graphics.setColor", engine.GraphicsSetColorFunction)

	// Keyboard Functions
	ghost.RegisterFunction("Keyboard.isDown", engine.KeyboardIsDownFunction)

	// Window Functions
	ghost.RegisterFunction("Window.bordered", engine.WindowBorderedFunction)
	ghost.RegisterFunction("Window.borderless", engine.WindowBorderlessFunction)
	ghost.RegisterFunction("Window.fullscreen", engine.WindowFullscreenFunction)
	ghost.RegisterFunction("Window.width", engine.WindowWidthFunction)
	ghost.RegisterFunction("Window.height", engine.WindowHeightFunction)
	ghost.RegisterFunction("Window.title", engine.WindowTitleFunction)

	ghost.NewScript(string(b))
	env := ghost.Evaluate()

	engine.Run(env)
}

func load(env *object.Environment) {
	ghost.Call(`
		if (load) {
			load()
		}
	`, env)
}

func update(env *object.Environment) {
	ghost.Call(`
		if (update) {
			update()
		}
	`, env)
}

func draw(env *object.Environment) {
	ghost.Call(`
		if (draw) {
			draw()
		}
	`, env)
}

func showHelp() {
	fmt.Println("Usage:")
	fmt.Println()
	fmt.Println("    engine [flags] {file}")
	fmt.Println()
	fmt.Println("Flags:")
	fmt.Println()
	fmt.Println("    -h  show help")
	fmt.Println("    -v  show version")
	fmt.Println()
	fmt.Println("Examples:")
	fmt.Println()
	fmt.Println("    engine game.ghost")
	fmt.Println()
	fmt.Println("            Execute source file (game.ghost)")
	fmt.Println()
	fmt.Println()
}
