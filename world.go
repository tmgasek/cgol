package main

import "math/rand"

// Game state.
type World struct {
	area   []bool
	width  int
	height int
}

// Create a new world.
func NewWorld(width, height, maxInitLiveCels int) *World {
	w := &World{
		area:   make([]bool, width*height),
		width:  width,
		height: height,
	}
	w.init(maxInitLiveCels)
	return w
}

// Initialize the world with a random state.
func (w *World) init(maxLiveCells int) {
	for i := 0; i < maxLiveCells; i++ {
		x := rand.Intn(w.width)
		y := rand.Intn(w.height)
		w.area[y*w.width+x] = true
	}
}

// Update game state by one tick.
func (w *World) Update() {

	width := w.width
	height := w.height
	next := make([]bool, width*height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {

			pop := neighbourCount(w.area, width, height, x, y)
			switch {
			case pop < 2:
				// Rule 1. Any live cell with fewer than two live neighbours
				// dies by under-population
				next[y*width+x] = false
			case (pop == 2 || pop == 3) && w.area[y*width+x]:
				// Rule 2. Any live cell with two or three live neighbours
				// lives on to the next generation.
				next[y*width+x] = true
			case pop > 3:
				// Rule 3. Any live cell with more than three live neighbours
				// dies by overcrowding.
				next[y*width+x] = false
			case pop == 3:
				// Rule 4. Any dead cell with exactly three live neighbours
				// becomes a live cell by reproduction.
				next[y*width+x] = true
			}
		}
	}
	w.area = next
}

// Paint current game state.
func (w *World) Draw(pix []byte) {
	for i, v := range w.area {
		if v {
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

func neighbourCount(a []bool, width, height, x, y int) int {
	c := 0
	for j := -1; j <= 1; j++ {
		for i := -1; i <= 1; i++ {
			if i == 0 && j == 0 {
				continue
			}
			x2 := x + i
			y2 := y + j
			if x2 < 0 || y2 < 0 || width <= x2 || height <= y2 {
				continue
			}
			if a[y2*width+x2] {
				c++
			}
		}
	}
	return c
}
