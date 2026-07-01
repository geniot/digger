package main

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type TextureImage struct {
	image   *rl.Image
	texture rl.Texture2D
	width   float32
	height  float32
}

func NewTextureImage(fileName string, degrees int32, flipHorizontal bool, flipVertical bool) *TextureImage {
	textureImage := &TextureImage{}
	imgBytes := orPanicRes(resList.ReadFile("res/" + fileName))
	textureImage.image = rl.LoadImageFromMemory(".png", imgBytes, int32(len(imgBytes)))
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
