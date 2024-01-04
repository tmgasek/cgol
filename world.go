package main

import "math/rand"

// Game state.
type World struct {
	area   [][]bool
	width  int
	height int
}

// Create a new world.
func NewWorld(width, height, maxInitLiveCells int) *World {
	w := &World{
		area:   make([][]bool, height),
		width:  width,
		height: height,
	}
	for i := range w.area {
		w.area[i] = make([]bool, width)
	}
	w.init(maxInitLiveCells)
	return w
}

// Initialize the world with a random state.
func (w *World) init(maxLiveCells int) {
	for i := 0; i < maxLiveCells; i++ {
		x := rand.Intn(w.width)
		y := rand.Intn(w.height)
		w.area[y][x] = true
	}
}

// Update game state by one tick.

// Update game state by one tick.
func (w *World) Update() {
	next := make([][]bool, w.height)
	for i := range next {
		next[i] = make([]bool, w.width)
	}

	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			pop := neighbourCount(w.area, x, y)
			switch {
			case pop < 2:
				next[y][x] = false
			case (pop == 2 || pop == 3) && w.area[y][x]:
				next[y][x] = true
			case pop > 3:
				next[y][x] = false
			case pop == 3:
				next[y][x] = true
			}
		}
	}
	w.area = next
}

// Paint current game state.
func (w *World) Draw(pix []byte) {
	for y := 0; y < w.height; y++ {
		for x := 0; x < w.width; x++ {
			i := y*w.width + x
			if w.area[y][x] {
				pix[4*i] = 0xff
				pix[4*i+1] = 0xff
				pix[4*i+2] = 0xff
				pix[4*i+3] = 0xff
			} else {
				pix[4*i] = 0
				pix[4*i+1] = 0
				pix[4*i+2] = 0
				pix[4*i+3] = 0
			}
		}
	}
}

func neighbourCount(a [][]bool, x, y int) int {
	c := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == 0 && j == 0 {
				continue
			}
			x2 := x + i
			y2 := y + j
			if x2 < 0 || y2 < 0 || y2 >= len(a) || x2 >= len(a[y2]) {
				continue
			}
			if a[y2][x2] {
				c++
			}
		}
	}
	return c
}

func (w *World) makeCellAndNeighborsAlive(mouseX, mouseY int) {
	cellWidth, cellHeight := w.cellDimensions()
	cellX := mouseX / cellWidth
	cellY := mouseY / cellHeight

	for dy := -1; dy <= 1; dy++ {
		for dx := -1; dx <= 1; dx++ {
			nx, ny := cellX+dx, cellY+dy
			if ny >= 0 && ny < w.height && nx >= 0 && nx < w.width {
				w.area[ny][nx] = true
			}
		}
	}
}

func (w *World) cellDimensions() (int, int) {
	cellWidth := screenWidth / w.width
	cellHeight := screenHeight / w.height
	return cellWidth, cellHeight
}
