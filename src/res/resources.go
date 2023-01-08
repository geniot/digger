package res

import (
	"bytes"
	"embed"
	"github.com/geniot/digger/src/ctx"
	"github.com/geniot/digger/src/glb"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"io/fs"
	"strconv"
)

var (
	//go:embed audio/*
	audioList embed.FS
	//go:embed media/*
	mediaList embed.FS
	//go:embed levels/*
	levelsList embed.FS
)

func GetImage(fileName string) *sdl.RWops {
	file, _ := mediaList.Open("media/" + fileName)
	return GetResource(file)
}

func GetAudio(fileName string) *sdl.RWops {
	file, _ := audioList.Open("audio/" + fileName)
	return GetResource(file)
}

func GetResource(file fs.File) *sdl.RWops {
	stat, _ := file.Stat()
	size := stat.Size()
	buf := make([]byte, size)
	file.Read(buf)
	rwOps, _ := sdl.RWFromMem(buf)
	return rwOps
}

func GetLevel(level int) string {
	file, _ := levelsList.Open("levels/" + strconv.FormatInt(int64(level), 10) + ".txt")
	buf := new(bytes.Buffer)
	buf.ReadFrom(file)
	return buf.String()
}

func LoadTexture(fileName string) *sdl.Texture {
	return LoadSurfTexture(fileName).T
}

func LoadSurfTexture(fileName string) *glb.SurfTexture {
	surface, err := img.LoadRW(GetImage(fileName), true)
	if err != nil {
		println(err.Error())
	}
	defer surface.Free()
	txt, err := ctx.RendererIns.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return &glb.SurfTexture{T: txt, W: surface.W, H: surface.H}
}
