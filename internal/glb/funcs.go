package glb

import (
	"fmt"
	"github.com/geniot/digger/internal/api"
	"math"
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

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func Bytes(s uint64) string {
	sizes := []string{"B", "kB", "MB", "GB", "TB", "PB", "EB"}
	return humanateBytes(s, 1000, sizes)
}

func humanateBytes(s uint64, base float64, sizes []string) string {
	if s < 10 {
		return fmt.Sprintf("%d B", s)
	}
	e := math.Floor(logn(float64(s), base))
	suffix := sizes[int(e)]
	val := math.Floor(float64(s)/math.Pow(base, e)*10+0.5) / 10
	//https://emptycharacter.com/
	f := "%.0f%s"
	//if val < 10 {
	//	f = "%.1f%s"
	//}

	return fmt.Sprintf(f, val, suffix)
}

func logn(n, b float64) float64 {
	return math.Log(n) / math.Log(b)
}
