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

func NewTextureImage(fileName string) *TextureImage {
	textureImage := &TextureImage{}
	imgBytes := orPanicRes(resList.ReadFile("res/" + fileName))
	textureImage.image = rl.LoadImageFromMemory(".png", imgBytes, int32(len(imgBytes)))
	textureImage.texture = rl.LoadTextureFromImage(textureImage.image)
	textureImage.width = float32(textureImage.texture.Width)
	textureImage.height = float32(textureImage.texture.Height)
	return textureImage
}
