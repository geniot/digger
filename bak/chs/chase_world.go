package chs

import (
	"fmt"
	"github.com/beefsack/go-astar"
	"strings"
)

type ChaseWorld map[int]map[int]*ChaseTile

// Tile ChaseTile gets the tile at the given coordinates in the world.
func (w ChaseWorld) Tile(x, y int) *ChaseTile {
	if w[x] == nil {
		return nil
	}
	return w[x][y]
}

// SetTile sets a tile at the given coordinates in the world.
func (w ChaseWorld) SetTile(t *ChaseTile, x, y int) {
	if w[x] == nil {
		w[x] = map[int]*ChaseTile{}
	}
	w[x][y] = t
	t.X = x
	t.Y = y
	t.W = w
}

// FirstOfKind gets the first tile on the board of a kind, used to get the from
// and to tiles as there should only be one of each.
func (w ChaseWorld) FirstOfKind(kind int) *ChaseTile {
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
func (w ChaseWorld) From() *ChaseTile {
	return w.FirstOfKind(KindFrom)
}

// To gets the to tile from the world.
func (w ChaseWorld) To() *ChaseTile {
	return w.FirstOfKind(KindTo)
}

// RenderPath renders a path on top of a world.
func (w ChaseWorld) RenderPath(path []astar.Pather) string {
	width := len(w)
	if width == 0 {
		return ""
	}
	height := len(w[0])
	pathLocs := map[string]bool{}
	for _, p := range path {
		pT := p.(*ChaseTile)
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

func (w ChaseWorld) Render() string {
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
func ParseWorld(input string) ChaseWorld {
	w := ChaseWorld{}
	for y, row := range strings.Split(strings.TrimSpace(input), "\n") {
		for x, raw := range row {
			kind, ok := RuneKinds[raw]
			if !ok {
				kind = KindField
			}
			w.SetTile(&ChaseTile{
				Kind: kind,
			}, x, y)
		}
	}
	return w
}
