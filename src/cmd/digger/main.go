package main

import (
	"github.com/geniot/digger/src/gui"
	"runtime/debug"
)

//import "github.com/pkg/profile"

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	debug.SetGCPercent(-1)
	gui.NewApplication().Start()
}