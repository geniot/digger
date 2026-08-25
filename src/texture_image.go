package main

import (
	"embed"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//go:embed res/*
	resList embed.FS
)

type TextureImage struct {
	image   *rl.Image
	texture rl.Texture2D
	width   float32
	height  float32
}

func NewTextureImage(fileName string, degrees int32, flipHorizontal bool, flipVertical bool, shouldMask bool) *TextureImage {
	textureImage := &TextureImage{}
	imgBytes := orPanicRes(resList.ReadFile("res/" + fileName))
	textureImage.image = rl.LoadImageFromMemory(".png", imgBytes, int32(len(imgBytes)))
	if shouldMask {
		rl.ImageColorTint(textureImage.image, rl.Black)
	}
	rl.ImageRotate(textureImage.image, degrees)
	if flipHorizontal {
		rl.ImageFlipHorizontal(textureImage.image)
	}
	if flipVertical {
		rl.ImageFlipVertical(textureImage.image)
	}
	textureImage.texture = rl.LoadTextureFromImage(textureImage.image)
	textureImage.width = float32(textureImage.image.Width)
	textureImage.height = float32(textureImage.image.Height)
	return textureImage
}

func initSprites(prefix string, degrees int32, flipHorizontal bool, flipVertical bool) []*TextureImage {
	sprites := make([]*TextureImage, 3)
	sprites[0] = NewTextureImage("graphics/digger/"+prefix+"1.png", degrees, flipHorizontal, flipVertical, false)
	sprites[1] = NewTextureImage("graphics/digger/"+prefix+"2.png", degrees, flipHorizontal, flipVertical, false)
	sprites[2] = NewTextureImage("graphics/digger/"+prefix+"3.png", degrees, flipHorizontal, flipVertical, false)
	return sprites
}
