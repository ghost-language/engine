package engine

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

// Image defines a new image resource
type Image struct {
	width     int32
	height    int32
	blendmode sdl.BlendMode
	surface   *sdl.Surface
	texture   *sdl.Texture
}

// NewImage loads a new image resource into memory.
func NewImage(_file string) *Image {
	image := new(Image)

	image.blendmode = 1

	addResource(image)

	image.surface, err = img.Load(_file)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(4)
	}

	image.texture, err = gameEngine.renderer.CreateTextureFromSurface(image.surface)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(5)
	}

	image.width = image.surface.W
	image.height = image.surface.H

	return image
}

func (image *Image) release() {
	if image.surface != nil {
		image.surface.Free()
	}

	if image.texture != nil {
		image.texture.Destroy()
	}
}

func (image *Image) buffer(_src *sdl.Rect, _dst *sdl.Rect) {
	image.texture.SetBlendMode(image.blendmode)
	gameEngine.renderer.Copy(image.texture, _src, _dst)
}

func (image *Image) Draw(_x int, _y int) {
	src := &sdl.Rect{0, 0, int32(image.width), int32(image.height)}
	dst := &sdl.Rect{int32(_x), int32(_y), int32(image.width), int32(image.height)}

	image.buffer(src, dst)
}
