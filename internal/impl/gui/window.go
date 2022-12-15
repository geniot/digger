package gui

import (
	"github.com/geniot/digger/internal/ctx"
	. "github.com/geniot/digger/internal/glb"
	"github.com/geniot/digger/resources"
	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
	"strconv"
)

type WindowImpl struct {
	sdlWindow   *sdl.Window
	iconSurface *sdl.Surface
}

func NewWindow() *WindowImpl {
	w := WindowImpl{}

	xPos, yPos, width, height := ctx.DeviceIns.GetWindowPosAndSize()
	w.sdlWindow, _ = sdl.CreateWindow(
		APP_NAME+" "+APP_VERSION,
		xPos, yPos, width, height,
		ctx.DeviceIns.GetWindowState())

	w.iconSurface, _ = img.LoadRW(resources.GetResource(ICON_FILE_NAME), true)
	w.sdlWindow.SetIcon(w.iconSurface)

	ctx.RendererIns, _ = sdl.CreateRenderer(w.sdlWindow, -1,
		sdl.RENDERER_PRESENTVSYNC|sdl.RENDERER_ACCELERATED)
	ctx.RendererIns.SetLogicalSize(SCREEN_LOGICAL_WIDTH, SCREEN_LOGICAL_HEIGHT)

	sdl.AddEventWatchFunc(w.resizingEventWatcher, nil)

	go w.show()

	return &w
}

func (window WindowImpl) show() {
	window.sdlWindow.Show()
	window.sdlWindow.Raise()
}

func (window WindowImpl) resizingEventWatcher(event sdl.Event, data interface{}) bool {
	switch t := event.(type) {
	case *sdl.WindowEvent:
		if t.Event == sdl.WINDOWEVENT_RESIZED {
			ctx.RenderLoopIns.Run()
		}
		break
	}
	return false
}

func (window WindowImpl) SaveWindowState() {
	width, height := window.sdlWindow.GetSize()
	xPos, yPos := window.sdlWindow.GetPosition()
	windowState := window.sdlWindow.GetFlags()
	ctx.ConfigIns.Set(WINDOW_STATE_KEY, strconv.FormatInt(int64(windowState), 10))

	if windowState&sdl.WINDOW_MAXIMIZED <= 0 {
		ctx.ConfigIns.Set(WINDOW_WIDTH_KEY, strconv.FormatInt(int64(width), 10))
		ctx.ConfigIns.Set(WINDOW_HEIGHT_KEY, strconv.FormatInt(int64(height), 10))
		ctx.ConfigIns.Set(WINDOW_XPOS_KEY, strconv.FormatInt(int64(xPos), 10))
		ctx.ConfigIns.Set(WINDOW_YPOS_KEY, strconv.FormatInt(int64(yPos), 10))
	}

	ctx.ConfigIns.Save()
}
