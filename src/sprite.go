package main

import "slices"

type Sprite struct {
	sequence []int
	textures []*MaskedTexture
}

func NewSprite(s []int, prefix string, degrees int32, flipHorizontal bool, flipVertical bool) *Sprite {
	sp := &Sprite{}
	sp.sequence = s
	sp.textures = initMaskedTextures(slices.Max(s)+1, prefix, degrees, flipHorizontal, flipVertical)
	return sp
}

func (sp *Sprite) Next(ptr int) int {
	return If(ptr < len(sp.sequence)-1, ptr+1, 0)
}

func (sp *Sprite) GetTexture(ptr int) *MaskedTexture {
	return sp.textures[sp.sequence[ptr]]
}
