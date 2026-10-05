package main

import (
	"embed"
	"strconv"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed res/*
	resList embed.FS
)

type MaskedTexture struct {
	fileName string
	texture  rl.Texture2D
	mask     [][]bool
	width    int32
	height   int32
}

func NewMaskedTexture(fn string, degrees int32, flipHorizontal bool, flipVertical bool, shouldMask bool) *MaskedTexture {
	maskedTexture := &MaskedTexture{fileName: fn}
	imgBytes := orPanicRes(resList.ReadFile("res/" + fn))
	image := rl.LoadImageFromMemory(".png", imgBytes, int32(len(imgBytes)))
	if shouldMask {
		rl.ImageColorTint(image, rl.Black)
	}
	rl.ImageRotate(image, degrees)
	if flipHorizontal {
		rl.ImageFlipHorizontal(image)
	}
	if flipVertical {
		rl.ImageFlipVertical(image)
	}
	maskedTexture.texture = rl.LoadTextureFromImage(image)
	maskedTexture.width = image.Width
	maskedTexture.height = image.Height
	//mask
	maskedTexture.mask = make([][]bool, image.Width)
	for x := range image.Width {
		maskedTexture.mask[x] = make([]bool, image.Height)
	}
	for y := range image.Height {
		for x := range image.Width {
			if IsPixelBlack(x, y, image) {
				maskedTexture.mask[x][y] = true
			}
		}
	}
	return maskedTexture
}

func initMaskedTextures(size int, prefix string, degrees int32, flipHorizontal bool, flipVertical bool) []*MaskedTexture {
	sprites := make([]*MaskedTexture, size)
	for i := range size {
		sprites[i] = NewMaskedTexture(prefix+strconv.Itoa(i+1)+".png", degrees, flipHorizontal, flipVertical, false)
	}
	return sprites
}
