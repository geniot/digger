package glb

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	IS_DEBUG_ON = false

	APP_NAME       = "Digger"
	APP_VERSION    = "0.1"
	CONF_FILE_NAME = ".digger.properties"
	FONT_FILE_NAME = "OpenSans-Regular.ttf"
	ICON_FILE_NAME = "digger.png"
	FONT_SIZE      = 12

	TICK               float64 = 1.0 / 200.0
	SPRITE_UPDATE_RATE         = 18
	DIGGER_SPEED               = 4 //less is faster
	DIGGER_DIE_SPEED           = 8
	DIGGER_GRAVE_SPEED         = SPRITE_UPDATE_RATE * 2
	BAG_PUSH_SPEED             = 8
	FIRE_SPEED                 = 2
	BAG_FALL_SPEED             = 2

	SCREEN_LOGICAL_WIDTH  = 320
	SCREEN_LOGICAL_HEIGHT = 240
	CELLS_HORIZONTAL      = 15
	CELLS_VERTICAL        = 10
	CELL_WIDTH            = 20
	CELL_HEIGHT           = 20
	CELLS_OFFSET          = 10
	FIELD_OFFSET_Y        = 20

	WINDOW_XPOS_KEY   = "WINDOW_XPOS_KEY"
	WINDOW_YPOS_KEY   = "WINDOW_YPOS_KEY"
	WINDOW_WIDTH_KEY  = "WINDOW_WIDTH_KEY"
	WINDOW_HEIGHT_KEY = "WINDOW_HEIGHT_KEY"
	WINDOW_STATE_KEY  = "WINDOW_STATE_KEY"

	DIGGER_COLLISION_TAG  = "digger"
	EMERALD_COLLISION_TAG = "emerald"
	BAG_COLLISION_TAG     = "bag"
	FIRE_COLLISION_TAG    = "fire"

	GCW_BUTTON_UP    = sdl.K_UP
	GCW_BUTTON_DOWN  = sdl.K_DOWN
	GCW_BUTTON_LEFT  = sdl.K_LEFT
	GCW_BUTTON_RIGHT = sdl.K_RIGHT

	GCW_BUTTON_A = sdl.K_LCTRL
	GCW_BUTTON_B = sdl.K_LALT
	GCW_BUTTON_X = sdl.K_SPACE
	GCW_BUTTON_Y = sdl.K_LSHIFT

	GCW_BUTTON_L1 = sdl.K_TAB
	GCW_BUTTON_R1 = sdl.K_BACKSPACE

	//GCW_BUTTON_L2 = sdl.K_RSHIFT
	//GCW_BUTTON_R2 = sdl.K_RALT

	GCW_BUTTON_L2 = sdl.K_PAGEUP
	GCW_BUTTON_R2 = sdl.K_PAGEDOWN

	GCW_BUTTON_SELECT = sdl.K_ESCAPE
	GCW_BUTTON_START  = sdl.K_RETURN
	GCW_BUTTON_MENU   = sdl.K_HOME

	GCW_VOLUMEUP   = sdl.K_VOLUMEUP
	GCW_VOLUMEDOWN = sdl.K_VOLUMEDOWN

	GCW_BUTTON_L3 = sdl.K_KP_DIVIDE
	//GCW_BUTTON_R3    = sdl.K_KP_PERIOD
	//GCW_BUTTON_POWER = sdl.K_HOME
)

var (
	COLOR_RED    = sdl.Color{R: 192, G: 64, B: 64, A: 255}
	COLOR_GREEN  = sdl.Color{R: 64, G: 192, B: 64, A: 255}
	COLOR_GRAY   = sdl.Color{R: 192, G: 192, B: 192, A: 255}
	COLOR_WHITE  = sdl.Color{R: 255, G: 255, B: 255, A: 255}
	COLOR_PURPLE = sdl.Color{R: 255, G: 0, B: 255, A: 255}
	COLOR_YELLOW = sdl.Color{R: 255, G: 255, B: 0, A: 255}
	COLOR_BLUE   = sdl.Color{R: 0, G: 255, B: 255, A: 255}
	COLOR_BLACK  = sdl.Color{R: 0, G: 0, B: 0, A: 255}

	BGR_COLOR = [4]uint8{0, 0, 0, 255} //black

	JoyButtonEventsMap = map[uint8]sdl.Keycode{
		6: GCW_BUTTON_UP,
		7: GCW_BUTTON_DOWN,
		8: GCW_BUTTON_LEFT,
		9: GCW_BUTTON_RIGHT,

		0: GCW_BUTTON_B,
		1: GCW_BUTTON_A,
		2: GCW_BUTTON_X,
		3: GCW_BUTTON_Y,

		10: GCW_BUTTON_SELECT,
		15: GCW_BUTTON_START,
		4:  GCW_BUTTON_L1,
		12: GCW_BUTTON_L2,
		5:  GCW_BUTTON_R1,
		13: GCW_BUTTON_R2,
	}
)
