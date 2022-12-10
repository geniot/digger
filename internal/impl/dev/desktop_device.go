package dev

import (
	"geniot.com/geniot/digger/internal/ctx"
	"geniot.com/geniot/digger/internal/glb"
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
	if ctx.PressedKeysCodesSetIns.Contains(sdl.K_q) {
		ctx.LoopIns.Stop()
	}
}

func (desktopDevice DesktopDeviceImpl) GetWindowPosAndSize() (int32, int32, int32, int32) {
	return int32(ctx.ConfigIns.Get(glb.WINDOW_XPOS_KEY)),
		int32(ctx.ConfigIns.Get(glb.WINDOW_YPOS_KEY)),
		640, 480
	//int32(ctx.ConfigIns.Get(glb.WINDOW_WIDTH_KEY)),
	//int32(ctx.ConfigIns.Get(glb.WINDOW_HEIGHT_KEY))
}

func (desktopDevice DesktopDeviceImpl) GetWindowState() uint32 {
	//return ctx.ConfigIns.Get(glb.WINDOW_STATE_KEY)
	return sdl.WINDOW_SHOWN
}

func NewDesktopDevice() DesktopDeviceImpl {
	device := DesktopDeviceImpl{}
	device.init()
	return device
}

func (desktopDevice DesktopDeviceImpl) init() {
	initCommon()
}
