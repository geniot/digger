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

type TextureImage struct {
	texture rl.Texture2D
	mask    [][]bool
	width   int32
	height  int32
}

func NewTextureImage(fileName string, degrees int32, flipHorizontal bool, flipVertical bool, shouldMask bool) *TextureImage {
	textureImage := &TextureImage{}
	imgBytes := orPanicRes(resList.ReadFile("res/" + fileName))
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
	textureImage.texture = rl.LoadTextureFromImage(image)
	textureImage.width = image.Width
	textureImage.height = image.Height
	//mask
	textureImage.mask = make([][]bool, image.Width)
	for x := range image.Width {
		textureImage.mask[x] = make([]bool, image.Height)
	}
	for y := range image.Height {
		for x := range image.Width {
			if IsPixelColored(x, y, image) {
				textureImage.mask[x][y] = true
			}
		}
	}
	return textureImage
}

func initTextureImages(size int, prefix string, degrees int32, flipHorizontal bool, flipVertical bool) []*TextureImage {
	sprites := make([]*TextureImage, size)
	for i := range size {
		sprites[i] = NewTextureImage(prefix+strconv.Itoa(i+1)+".png", degrees, flipHorizontal, flipVertical, false)
	}
	return sprites
}
