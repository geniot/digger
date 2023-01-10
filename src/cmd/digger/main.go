package main

import (
	"github.com/geniot/digger/src/gui"
	"runtime"
	"runtime/debug"
	"strings"
)

//import "github.com/pkg/profile"

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()

	//On PocketGo garbage collection causes frame drops (freezes), so I collect garbage when digger dies
	if strings.Index(runtime.GOARCH, "mips") == 0 {
		debug.SetGCPercent(-1)
	}
	gui.NewApplication().Start()
}
