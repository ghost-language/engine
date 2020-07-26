package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path"
	"strconv"

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

	engine.SetDrawFunction(draw)

	ghost.RegisterFunction("draw_image", drawImageFunction)
	ghost.NewScript(string(b))

	engine.Run()
}

func draw() {
	ghost.Evaluate()
}

func drawImageFunction(args ...object.Object) object.Object {
	if len(args) != 3 {
		return ghost.NewError("wrong number of arguments. got=%d, expected=3", len(args))
	}

	file := args[0].Inspect()
	x, _ := strconv.ParseInt(args[1].Inspect(), 10, 64)
	y, _ := strconv.ParseInt(args[2].Inspect(), 10, 64)

	image := engine.NewImage(file)

	image.Draw(int(x), int(y))

	return ghost.NULL
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
