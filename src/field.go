package main

import (
	"image/color"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Field struct {
	scene            *GameScene
	texture          rl.RenderTexture2D
	image            *rl.Image
	textureSourceRec rl.Rectangle
	imageSourceRec   rl.Rectangle
	destRec          rl.Rectangle
}

func NewField(scene *GameScene) *Field {
	fld := &Field{}
	fld.scene = scene

	fld.textureSourceRec = rl.NewRectangle(0, 0, FIELD_WIDTH, -FIELD_HEIGHT) //see https://github.com/raysan5/raylib/issues/3803
	fld.imageSourceRec = rl.NewRectangle(0, 0, FIELD_WIDTH, FIELD_HEIGHT)
	fld.destRec = rl.NewRectangle(0, 0, FIELD_WIDTH, FIELD_HEIGHT)

	fld.texture = rl.LoadRenderTexture(FIELD_WIDTH, FIELD_HEIGHT)
	fld.image = rl.GenImageColor(FIELD_WIDTH, FIELD_HEIGHT, rl.Black)

	rl.BeginTextureMode(fld.texture)
	rl.ClearBackground(rl.Black)
	l := 1 //level
	for y := 14; y < 200; y += 4 {
		for x := 0; x < 320; x += 20 {
			fld.drawmiscspr(x, y, 93+l, 5, 4)
		}
	}
	rl.EndTextureMode()
	return fld
}

var (
	pixels   = [256 * 256]int32{} //65536
	sprch    = [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprx     = [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	spry     = [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprwid   = [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprhei   = [17]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprbwid  = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprbhei  = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprnch   = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprnwid  = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprnhei  = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprnbwid = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sprnbhei = [16]int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
)

func (field *Field) drawmiscspr(x, y, ch, wid, hei int) {
	sprx[16] = x & -4
	spry[16] = y
	sprch[16] = ch
	sprwid[16] = wid
	sprhei[16] = hei
	field.gputim(sprx[16], spry[16], sprch[16], sprwid[16], sprhei[16])
}

func (field *Field) gputim(x, y, ch, w, h int) {
	width := 320
	//height := 200
	//size := width * height

	spr := cgatable[ch*2]
	msk := cgatable[ch*2+1]

	src := 0
	dest := y*width + (x & 0xfffc)

	for i := 0; i < h; i++ {
		d := dest
		for j := 0; j < w; j++ {
			px := spr[src]
			mx := msk[src]
			src++
			if mx&3 == 0 {
				field.drawPixel(d+3, int32(px&3))
				pixels[d+3] = int32(px & 3)
			}
			px >>= 2
			if (mx & (3 << 2)) == 0 {
				field.drawPixel(d+2, int32(px&3))
				pixels[d+2] = int32(px & 3)
			}
			px >>= 2
			if (mx & (3 << 4)) == 0 {
				field.drawPixel(d+1, int32(px&3))
				pixels[d+1] = int32(px & 3)
			}
			px >>= 2
			if (mx & (3 << 6)) == 0 {
				field.drawPixel(d, int32(px&3))
				pixels[d] = int32(px & 3)
			}
			d += 4
			if src == len(spr) || src == len(msk) {
				return
			}
		}
		dest += width
	}
}

func (field *Field) drawPixel(d int, c int32) {
	xPos := int32(d % 320)
	yPos := int32(d / 320)
	red := If(c&0xFF00 > 0, 255, 0)
	green := If(c&0x00FF00 > 0, 255, 0)
	blue := If(c&0x0000FF > 0, 255, 0)
	//red := ((c & 0xC0) >> 6) * 64
	//green := ((c & 0x30) >> 4) * 64
	//blue := ((c & 0x0C) >> 2) * 64
	col := color.RGBA{R: uint8(red), G: uint8(green), B: uint8(blue), A: 255}
	rl.DrawPixel(xPos, yPos, col)
}

func (field *Field) draw(x float32, y float32, textureImage *TextureImage) {
	sourceRect := rl.NewRectangle(0, 0, textureImage.width, textureImage.height)
	destRect := rl.NewRectangle(x, y, textureImage.width, textureImage.height)
	rl.DrawTexturePro(textureImage.texture, sourceRect, destRect, ZERO_VECTOR2, 0, rl.White)
	rl.ImageDraw(field.image, textureImage.image, sourceRect, destRect, rl.White)
}

func (field *Field) Update(_ int64) {
}

func (field *Field) Render(drawTarget rl.RenderTexture2D) {
	//field.Debug()
	rl.BeginTextureMode(drawTarget)
	//rl.DrawTextureRec(rl.LoadTextureFromImage(field.image), field.imageSourceRec, ZERO_VECTOR2, rl.White)
	rl.DrawTexturePro(field.texture.Texture, field.textureSourceRec, field.destRec, ZERO_VECTOR2, 0, rl.White)
	rl.EndTextureMode()
}

func (field *Field) Debug() {
	clone1 := rl.ImageCopy(field.image)
	rl.ImageFlipVertical(clone1)
	colors1 := rl.LoadImageColors(clone1)
	defer rl.UnloadImageColors(colors1)

	clone2 := rl.LoadImageFromTexture(field.texture.Texture)
	colors2 := rl.LoadImageColors(clone2)
	defer rl.UnloadImageColors(colors2)

	//println(len(colors1))
	if len(colors1) != len(colors2) {
		panic("colors are different")
	}
	for i := 0; i < len(colors1); i++ {
		if colors1[i].R != colors2[i].R || colors1[i].G != colors2[i].G || colors1[i].B != colors2[i].B || colors1[i].A != colors2[i].A {
			println(colors1[i].R, " ", colors1[i].G, " ", colors1[i].B, " ", colors1[i].A, " ")
			println(colors2[i].R, " ", colors2[i].G, " ", colors2[i].B, " ", colors2[i].A, " ")
			panic("colors are different")
		}
	}
}
