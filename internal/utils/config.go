package utils

import (
	"geniot.com/geniot/digger/internal/model"
	"github.com/magiconair/properties"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type Config interface {
	Get(key string) int32
}
type config struct {
	props *properties.Properties
}

func NewConfig() Config {
	props := properties.NewProperties()
	props.Set(model.WINDOW_XPOS_KEY, strconv.FormatInt(int64(sdl.WINDOWPOS_UNDEFINED), 10))
	props.Set(model.WINDOW_YPOS_KEY, strconv.FormatInt(int64(sdl.WINDOWPOS_UNDEFINED), 10))
	displayMode, err := sdl.GetCurrentDisplayMode(0)
	if err != nil {
		println(err.Error())
	}
	props.Set(model.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(displayMode.W/2), 10))
	props.Set(model.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(displayMode.H/2), 10))

	loadedProps, _ := properties.LoadFile(model.PATH_TO_CONFIG, properties.UTF8)
	if loadedProps != nil {
		props = loadedProps
	}
	return config{props}
}

func (cfg config) Get(key string) int32 {
	valStr, _ := cfg.props.Get(key)
	valInt, _ := strconv.ParseInt(valStr, 10, 0)
	return int32(valInt)
}
