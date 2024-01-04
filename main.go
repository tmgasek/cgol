package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

var (
	screenWidth  int
	screenHeight int
	cellDensity  float64
)

func main() {
	width := flag.Int("width", 720, "screen width")
	height := flag.Int("height", 480, "screen height")
	density := flag.Float64("density", 1.0, "cell density")
	flag.Parse()

	screenWidth = *width
	screenHeight = *height
	cellDensity = *density

	g := &Game{
		world: NewWorld(screenWidth, screenHeight, int(float64(screenWidth*screenHeight)/10*cellDensity)),
	}

	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Game of Life")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
