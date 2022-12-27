package glb

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"strings"
)

type World map[int]map[int]*Tile

// Tile gets the tile at the given coordinates in the world.
func (w World) Tile(x, y int) *Tile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w World) SetTile(t *Tile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*Tile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

func (w World) SetTiles(kinds [9]int, x, y int) {
	realX := x * 3
	realY := y * 3

	w.SetTile(&Tile{kinds[0], realX, realY, w}, realX, realY)
	w.SetTile(&Tile{kinds[1], realX + 1, realY, w}, realX+1, realY)
	w.SetTile(&Tile{kinds[2], realX + 2, realY, w}, realX+2, realY)

	w.SetTile(&Tile{kinds[3], realX, realY + 1, w}, realX, realY+1)
	w.SetTile(&Tile{kinds[4], realX + 1, realY + 1, w}, realX+1, realY+1)
	w.SetTile(&Tile{kinds[5], realX + 2, realY + 1, w}, realX+2, realY+1)

	w.SetTile(&Tile{kinds[6], realX, realY + 2, w}, realX, realY+2)
	w.SetTile(&Tile{kinds[7], realX + 1, realY + 2, w}, realX+1, realY+2)
	w.SetTile(&Tile{kinds[8], realX + 2, realY + 2, w}, realX+2, realY+2)
}

// FirstOfKind gets the first tile on the board of a kind, used to get the from
// and to tiles as there should only be one of each.
func (w World) FirstOfKind(kind int) *Tile {
	for _, row := range w {
		for _, t := range row {
			if t.Kind == kind {
				return t
			}
		}
	}
	return nil
}

// From gets the from tile from the world.
func (w World) From() *Tile {
	return w.FirstOfKind(KindFrom)
}

// To gets the to tile from the world.
func (w World) To() *Tile {
	return w.FirstOfKind(KindTo)
}

// RenderPath renders a path on top of a world.
func (w World) RenderPath(path []astar.Pather) string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*Tile)
		pathLocs[fmt.Sprintf("%d,%d", pT.X, pT.Y)] = true
	}
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.Tile(x, y)
			r := ' '
			if pathLocs[fmt.Sprintf("%d,%d", x, y)] {
				r = KindRunes[KindPath]
			} else if t != nil {
				r = KindRunes[t.Kind]
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

func (w World) Render() string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	rows := make([]string, height)
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			t := w.Tile(x, y)
			r := ' '
			if t != nil {
				r = KindRunes[t.Kind]
			}
			rows[y] += string(r)
		}
	}
	return strings.Join(rows, "\n")
}

// ParseWorld parses a textual representation of a world into a world map.
func ParseWorld(input string) World {
	w := World{}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			kind, ok := RuneKinds[raw]
			if !ok {
				kind = KindField
			}
			w.SetTile(&Tile{
				Kind: kind,
			}, x, y)
		}
	}
	return w
}
