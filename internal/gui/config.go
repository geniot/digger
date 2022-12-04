package gui

import (
	"geniot.com/geniot/digger/internal/model"
	"github.com/magiconair/properties"
	"github.com/veandco/go-sdl2/sdl"
	"os"
	"strconv"
)

type Config struct {
	application *Application
	props       *properties.Properties
}

func NewConfig(app *Application) *Config {
	return &Config{app, Load()}
}

func (cfg Config) Get(key string) uint32 {
	valStr, _ := cfg.props.Get(key)
	valInt, _ := strconv.ParseInt(valStr, 10, 0)
	return uint32(valInt)
}

func (cfg Config) Set(key string, value string) {
	cfg.props.Set(key, value)
}

func Load() *properties.Properties {
	loadedProps, _ := properties.LoadFile(model.PATH_TO_CONFIG, properties.UTF8)

	if loadedProps == nil {
		loadedProps = properties.NewProperties()
		loadedProps.Set(model.WINDOW_XPOS_KEY, strconv.FormatInt(int64(sdl.WINDOWPOS_UNDEFINED), 10))
		loadedProps.Set(model.WINDOW_YPOS_KEY, strconv.FormatInt(int64(sdl.WINDOWPOS_UNDEFINED), 10))
		displayMode, _ := sdl.GetCurrentDisplayMode(0)
		loadedProps.Set(model.WINDOW_WIDTH_KEY, strconv.FormatInt(int64(displayMode.W/2), 10))
		loadedProps.Set(model.WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(displayMode.H/2), 10))
		loadedProps.Set(model.WINDOW_STATE_KEY, strconv.FormatInt(int64(sdl.WINDOW_SHOWN|sdl.WINDOW_RESIZABLE), 10))
	}

	//patching window state
	windowStateStr, _ := loadedProps.Get(model.WINDOW_STATE_KEY)
	windowState, _ := strconv.ParseInt(windowStateStr, 10, 0)
	windowState |= sdl.WINDOW_SHOWN
	windowState |= sdl.WINDOW_RESIZABLE
	loadedProps.Set(model.WINDOW_STATE_KEY, strconv.FormatInt(windowState, 10))

	return loadedProps
}

func (cfg Config) Save() {
	f, err := os.OpenFile(model.PATH_TO_CONFIG, os.O_WRONLY|os.O_CREATE, 0600)
	if err != nil {
		println(err.Error())
	}
	defer f.Close()
	err = f.Truncate(0)
	_, err = f.Seek(0, 0)
	cfg.props.Write(f, properties.UTF8)
}
