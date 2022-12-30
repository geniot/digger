package dev

import (
	"github.com/geniot/digger/internal/api"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
	"os"
	"runtime"
	"strings"
)

func NewDevice() api.IDevice {
	if strings.Index(runtime.GOARCH, "mips") == 0 {
		return NewHandheldDevice()
	} else {
		return NewDesktopDevice()
	}
}

func initCommon() {
	err := ttf.Init()
	if err != nil {
		panic(err)
	}
	err = sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		panic(err)
	}
	err = mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 4096)
	if err != nil {
		println(err.Error())
	}
	sdl.GameControllerEventState(sdl.ENABLE)
	sdl.JoystickEventState(sdl.ENABLE)
	println(sdl.JoystickGetDeviceGUID(0))
	println(sdl.GameControllerAddMapping(os.Getenv("SDL_GAMECONTROLLERCONFIG")))
}

func closeCommon() {
	ttf.Quit()
	sdl.Quit()
	mix.CloseAudio()
}
