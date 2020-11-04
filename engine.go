package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"

	"ghostlang.org/x/engine/engine"
	"ghostlang.org/x/ghost/ghost"
	"ghostlang.org/x/ghost/object"
)

var (
	flagVersion bool
	flagHelp    bool
)

type Game struct {
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

	if flagHelp || len(args) == 0 {
		showHelp()
		os.Exit(2)
	}

	game := Game{args}

	f, err := os.Open(game.args[0])

	if err != nil {
		log.Fatalf("could not open source file %s: %s", game.args[0], err)
	}

	b, err := ioutil.ReadAll(f)

	if err != nil {
		fmt.Fprintf(os.Stderr, "error reading source file: %s", err)
		return
	}

	engine := engine.NewEngine("Simple Ghost Example")
	engine.SetWindow(800, 600)

	engine.SetLoadFunction(load)
	engine.SetUpdateFunction(update)
	engine.SetDrawFunction(draw)

	ghost.RegisterFunction("Graphics.draw", engine.GraphicsDrawFunction)
	ghost.RegisterFunction("Keyboard.isDown", engine.KeyboardIsDownFunction)

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
