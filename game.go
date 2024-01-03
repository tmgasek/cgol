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
		g.makeCellAndNeighborsAlive(x, y)
	}
}

func (g *Game) makeCellAndNeighborsAlive(mouseX, mouseY int) {
	cellWidth, cellHeight := g.cellDimensions()

	// Convert mouse coordinates to cell position.
	cellX := mouseX / cellWidth
	cellY := mouseY / cellHeight

	// Update the clicked cell and its neighbors.
	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			nx, ny := cellX+dx, cellY+dy
			// Check bounds.
			if nx >= 0 && nx < g.world.width && ny >= 0 && ny < g.world.height {
				g.world.area[ny*g.world.width+nx] = true
			}
		}
	}
}

func (g *Game) cellDimensions() (int, int) {
	cellWidth := screenWidth / g.world.width
	cellHeight := screenHeight / g.world.height
	return cellWidth, cellHeight
}
