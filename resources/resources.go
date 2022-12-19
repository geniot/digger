package resources

import (
	"bytes"
	"embed"
	"github.com/geniot/digger/internal/ctx"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

var (
	//go:embed media/*
	mediaList embed.FS
	//go:embed levels/*
	levelsList embed.FS
)

func GetResource(fileName string) *sdl.RWops {
	file, _ := mediaList.Open("media/" + fileName)
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
	surface, _ := img.LoadRW(GetResource(fileName), true)
	defer surface.Free()
	txt, err := ctx.RendererIns.CreateTextureFromSurface(surface)
	if err != nil {
		println(err.Error())
	}
	return txt
}
