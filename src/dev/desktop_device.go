package dev

import (
	"github.com/geniot/digger/src/ctx"
	. "github.com/geniot/digger/src/glb"
	"github.com/veandco/go-sdl2/sdl"
)

type DesktopDeviceImpl struct {
}

func (desktopDevice DesktopDeviceImpl) GetJoystickAxis(axis int) int16 {
	return 0 //no joystick on desktop
}

func (desktopDevice DesktopDeviceImpl) Stop() {
	closeCommon()
}

func (desktopDevice DesktopDeviceImpl) ProcessKeyActions() {
	if _, ok := ctx.PressedKeysCodesSetIns[sdl.K_q]; ok {
		ctx.LoopIns.Stop()
	}
}

func (desktopDevice DesktopDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	//return 0, 0, 320, 240
	return int32(ctx.ConfigIns.Get(WINDOW_XPOS_KEY)),
		int32(ctx.ConfigIns.Get(WINDOW_YPOS_KEY)),
		//SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT
		int32(ctx.ConfigIns.Get(WINDOW_WIDTH_KEY)),
		int32(ctx.ConfigIns.Get(WINDOW_HEIGHT_KEY))
}

func (desktopDevice DesktopDeviceImpl) GetWindowState() uint32 {
	return ctx.ConfigIns.Get(WINDOW_STATE_KEY)
	//return sdl.WINDOW_HIDDEN
}

func NewDesktopDevice() DesktopDeviceImpl {
	device := DesktopDeviceImpl{}
	device.init()
	return device
}

func (desktopDevice DesktopDeviceImpl) init() {
	initCommon()
}
