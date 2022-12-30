package dev

import (
	"github.com/geniot/digger/internal/ctx"
	"github.com/geniot/digger/internal/glb"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type HandheldDeviceImpl struct {
	haptic            *sdl.Haptic
	joystick          *sdl.Joystick
	isRumbleSupported bool
}

func (handhelpDevice HandheldDeviceImpl) GetJoystickAxis(axis int) int16 {
	return handhelpDevice.joystick.Axis(axis)
}

func (handhelpDevice HandheldDeviceImpl) Stop() {
	handhelpDevice.joystick.Close()
	handhelpDevice.haptic.Close()
	closeCommon()
}

func (handhelpDevice HandheldDeviceImpl) ProcessKeyActions() {
	_, ok1 := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_L1]
	_, ok2 := ctx.PressedKeysCodesSetIns[glb.GCW_BUTTON_START]
	if ok1 && ok2 {
		ctx.LoopIns.Stop()
	}
}

func (handhelpDevice HandheldDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return 0, 0, 320, 240
}

func (handhelpDevice HandheldDeviceImpl) GetWindowState() uint32 {
	return sdl.WINDOW_SHOWN | sdl.WINDOW_BORDERLESS
}

func NewHandheldDevice() HandheldDeviceImpl {
	device := HandheldDeviceImpl{}
	device.init()
	return device
}

func (handhelpDevice *HandheldDeviceImpl) init() {
	initCommon()
	numHaptics, err := sdl.NumHaptics()
	if err != nil {
		panic(err)
	}
	if numHaptics > 0 {
		println("Haptics: " + strconv.Itoa(numHaptics))
		println(sdl.HapticName(0))
		handhelpDevice.haptic, err = sdl.HapticOpen(0)
		if err != nil {
			panic(err)
		}
		err = handhelpDevice.haptic.RumbleInit()
		if err != nil {
			panic(err)
		}
		handhelpDevice.isRumbleSupported, _ = handhelpDevice.haptic.RumbleSupported()
	}
	numJoysticks := sdl.NumJoysticks()
	if numJoysticks > 0 {
		println("Joysticks: " + strconv.Itoa(numJoysticks))
		println(sdl.JoystickNameForIndex(0))
		handhelpDevice.joystick = sdl.JoystickOpen(0)
	}
}
