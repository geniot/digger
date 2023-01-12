package rnd

import (
	"github.com/geniot/digger/src/res"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Media struct {
	bagTexture                   *sdl.Texture
	bagTextureFall               *sdl.Texture
	bagSpritesShake              []*sdl.Texture
	bagSpritesGold               []*sdl.Texture
	bagSpritesShakeFrameSequence []int
	bagSpritesGoldFrameSequence  []int

	diggerSprites                   []*sdl.Texture
	diggerDieTexture                *sdl.Texture
	diggerSpritesGraveFrameSequence []int
	diggerSpritesGrave              []*sdl.Texture

	emeraldTexture     *sdl.Texture
	emeraldTextureMask *sdl.Surface

	fireSprites          []*sdl.Texture
	fireSpritesExplosion []*sdl.Texture

	fieldHorizontalBlob *sdl.Surface
	fieldVerticalBlob   *sdl.Surface
	fieldEndLeftBlob    *sdl.Surface
	fieldEndRightBlob   *sdl.Surface
	fieldEndUpBlob      *sdl.Surface
	fieldEndDownBlob    *sdl.Surface

	monsterSpritesNobbin []*sdl.Texture
	monsterSpritesHobbin []*sdl.Texture

	soundDiggerTune *mix.Chunk
	soundEatEmerald [8]*mix.Chunk
	soundEatGold    *mix.Chunk
	soundFire       *mix.Chunk
	soundExplode    *mix.Chunk
	soundWobble     *mix.Chunk
	soundFall       *mix.Chunk
	soundBagToGold  *mix.Chunk
}

func NewMedia() *Media {
	md := &Media{}
	md.bagTexture = res.LoadTexture("csbag.png")
	md.bagTextureFall = res.LoadTexture("cfbag.png")
	md.bagSpritesShake = []*sdl.Texture{res.LoadTexture("csbag.png"), res.LoadTexture("clbag.png"), res.LoadTexture("crbag.png")}
	md.bagSpritesShakeFrameSequence = []int{0, 1, 2, 1, 2}
	md.bagSpritesGold = []*sdl.Texture{res.LoadTexture("cgold1.png"), res.LoadTexture("cgold2.png"), res.LoadTexture("cgold3.png")}
	md.bagSpritesGoldFrameSequence = []int{0, 1, 2}

	md.diggerSprites = []*sdl.Texture{
		res.LoadTexture("cldig1.png"),
		res.LoadTexture("cldig2.png"),
		res.LoadTexture("cldig3.png")}

	md.diggerDieTexture = res.LoadTexture("cddie.png")
	md.diggerSpritesGrave = []*sdl.Texture{
		res.LoadTexture("cgrave1.png"),
		res.LoadTexture("cgrave2.png"),
		res.LoadTexture("cgrave3.png"),
		res.LoadTexture("cgrave4.png"),
		res.LoadTexture("cgrave5.png"),
	}
	md.diggerSpritesGraveFrameSequence = []int{0, 1, 2, 3, 4, 4, 4, 4, 4, 4, 4, 4} //making a pause at the end

	md.emeraldTexture = res.LoadTexture("emerald.png")
	md.emeraldTextureMask, _ = img.LoadRW(res.GetImage("emerald_mask.png"), true)

	md.fieldHorizontalBlob, _ = img.LoadRW(res.GetImage("blob1.png"), true)
	md.fieldVerticalBlob, _ = img.LoadRW(res.GetImage("blob2.png"), true)
	md.fieldEndLeftBlob, _ = img.LoadRW(res.GetImage("blob3.png"), true)
	md.fieldEndRightBlob, _ = img.LoadRW(res.GetImage("blob4.png"), true)
	md.fieldEndUpBlob, _ = img.LoadRW(res.GetImage("blob5.png"), true)
	md.fieldEndDownBlob, _ = img.LoadRW(res.GetImage("blob6.png"), true)

	md.fireSprites = []*sdl.Texture{res.LoadTexture("cfire1.png"), res.LoadTexture("cfire2.png"), res.LoadTexture("cfire3.png")}
	md.fireSpritesExplosion = []*sdl.Texture{res.LoadTexture("cexp1.png"), res.LoadTexture("cexp2.png"), res.LoadTexture("cexp3.png")}

	md.monsterSpritesNobbin = []*sdl.Texture{
		res.LoadTexture("cnob1.png"),
		res.LoadTexture("cnob2.png"),
		res.LoadTexture("cnob3.png")}

	md.monsterSpritesHobbin = []*sdl.Texture{
		res.LoadTexture("clhob1.png"),
		res.LoadTexture("clhob2.png"),
		res.LoadTexture("clhob3.png")}

	for i := 0; i <= 7; i++ {
		md.soundEatEmerald[i], _ = mix.LoadWAVRW(res.GetAudio("emerald"+strconv.FormatInt(int64(i), 10)+".wav"), true)
	}
	md.soundDiggerTune, _ = mix.LoadWAVRW(res.GetAudio("digger.wav"), true)
	md.soundEatGold, _ = mix.LoadWAVRW(res.GetAudio("gold.wav"), true)
	md.soundFire, _ = mix.LoadWAVRW(res.GetAudio("fire.wav"), true)
	md.soundExplode, _ = mix.LoadWAVRW(res.GetAudio("explode.wav"), true)
	md.soundWobble, _ = mix.LoadWAVRW(res.GetAudio("wobble.wav"), true)
	md.soundFall, _ = mix.LoadWAVRW(res.GetAudio("fall.wav"), true)
	md.soundBagToGold, _ = mix.LoadWAVRW(res.GetAudio("bag2gold.wav"), true)

	return md
}
