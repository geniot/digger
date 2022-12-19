package glb

import (
	"github.com/geniot/digger/internal/api"
	"strconv"
)

func Opposite(dir api.Direction) api.Direction {
	if dir == UP {
		return DOWN
	} else if dir == DOWN {
		return UP
	} else if dir == LEFT {
		return RIGHT
	} else if dir == RIGHT {
		return LEFT
	} else {
		println("Unidentified direction: " + strconv.FormatInt(int64(dir), 10))
		return dir
	}
}
