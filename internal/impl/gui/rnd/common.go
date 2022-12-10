package rnd

import (
	"geniot.com/geniot/digger/internal/ctx"
	"github.com/veandco/go-sdl2/sdl"
)

func getTextWidth(text string) int32 {
	width, _, _ := ctx.FontIns.SizeUTF8(text)
	return int32(width)
}

func drawText(txt string, x int32, y int32, color sdl.Color) int32 {
	textSurface, _ := ctx.FontIns.RenderUTF8Blended(txt, color)
	defer textSurface.Free()
	textTexture, _ := ctx.RendererIns.CreateTextureFromSurface(textSurface)
	ctx.RendererIns.Copy(textTexture, nil,
		&sdl.Rect{X: x, Y: y, W: textSurface.W, H: textSurface.H})
	defer textTexture.Destroy()
	return textSurface.W
}
