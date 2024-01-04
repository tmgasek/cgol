package main

import "github.com/hajimehoshi/ebiten/v2"

type Game struct {
	world  *World
	pixels []byte
}

func (g *Game) Update() error {
	g.handleMouseInput()
	g.world.Update()
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if g.pixels == nil {
		g.pixels = make([]byte, screenWidth*screenHeight*4)
	}
	g.world.Draw(g.pixels)
	screen.WritePixels(g.pixels)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}

func (g *Game) handleMouseInput() {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		x, y := ebiten.CursorPosition()
		g.world.makeCellAndNeighborsAlive(x, y)
	}
}
