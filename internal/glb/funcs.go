package glb

import (
	"fmt"
	"github.com/geniot/digger/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"strconv"
)

func Opposite(dir Direction) Direction {
	if dir == UP {
		return DOWN
	} else if dir == DOWN {
		return UP
	} else if dir == LEFT {
		return RIGHT
	} else if dir == RIGHT {
		return LEFT
	} else {
		println("Unidentified direction: " + strconv.FormatInt(int64(dir), 10))
		return dir
	}
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func Bytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, sizes)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	//https://emptycharacter.com/
	f := "%.0f%s"
	//if val < 10 {
	//	f = "%.1f%s"
	//}

	return fmt.Sprintf(f, val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}

func DrawText(txt string, color sdl.Color) *SurfTexture {
	textSurface, err := ctx.FontIns.RenderUTF8Blended(txt, color)
	if err != nil {
		println(err.Error())
	}
	defer textSurface.Free()
	textTexture, err := ctx.RendererIns.CreateTextureFromSurface(textSurface)
	if err != nil {
		println(err.Error())
	}
	return &SurfTexture{T: textTexture, W: textSurface.W, H: textSurface.H}
}

// Renderer.DrawRect is not available in some older versions of SDL2.
// This is a workaround to solve hanging on my PocketGo2v2
func DrawRectLines(rect *sdl.Rect) {
	ctx.RendererIns.DrawLine(rect.X, rect.Y, rect.X+rect.W, rect.Y)
	ctx.RendererIns.DrawLine(rect.X+rect.W, rect.Y, rect.X+rect.W, rect.Y+rect.H)
	ctx.RendererIns.DrawLine(rect.X+rect.W, rect.Y+rect.H, rect.X, rect.Y+rect.H)
	ctx.RendererIns.DrawLine(rect.X, rect.Y+rect.H, rect.X, rect.Y)
}

func Collide(rect1 *sdl.Rect, rect2 *sdl.Rect) bool {
	x1 := rect1.X
	y1 := rect1.Y
	x2 := x1 + rect1.W
	y2 := y1 + rect1.H
	x3 := rect2.X
	y3 := rect2.Y
	x4 := x3 + rect2.W
	y4 := y3 + rect2.H
	// If one rectangle is on left side of other
	if x1 > x4 || x3 > x2 {
		return false
	}
	// If one rectangle is above other
	if y2 < y3 || y4 < y1 {
		return false
	}
	return true
}

func GetNextSpritePointerAndInc(currentPointer int, spritePointerInc int, spritesLen int) (int, int) {
	nextSpritePointerInc := spritePointerInc
	nextSpritePointer := currentPointer + spritePointerInc
	if nextSpritePointer >= spritesLen {
		nextSpritePointer = spritesLen - 1
	}
	if nextSpritePointer < 0 {
		nextSpritePointer = 0
	}
	if nextSpritePointer == spritesLen-1 || nextSpritePointer == 0 {
		nextSpritePointerInc = -spritePointerInc
	}
	return nextSpritePointer, nextSpritePointerInc
}
