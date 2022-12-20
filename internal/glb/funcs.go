package glb

import (
	"fmt"
	"github.com/geniot/digger/internal/api"
	"github.com/geniot/digger/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"strconv"
)

func Opposite(dir api.Direction) api.Direction {
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

func DrawText(txt string, x int32, y int32, color sdl.Color) (int32, int32) {
	textSurface, _ := ctx.FontIns.RenderUTF8Blended(txt, color)
	defer textSurface.Free()
	textTexture, _ := ctx.RendererIns.CreateTextureFromSurface(textSurface)
	ctx.RendererIns.Copy(textTexture, nil,
		&sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
	defer textTexture.Destroy()
	return textSurface.W, textSurface.H
}

func DrawRect(x int32, y int32, width int32, height int32) {
	ctx.RendererIns.DrawLine(x, y, x+width, y)
	ctx.RendererIns.DrawLine(x+width, y, x+width, y+height)
	ctx.RendererIns.DrawLine(x+width, y+height, x, y+height)
	ctx.RendererIns.DrawLine(x, y+height, x, y)
}
