package main

import (
	"github.com/geniot/digger/internal/impl/gui"
)

//import "github.com/pkg/profile"

func main() {
	//defer profile.Start(profile.CPUProfile).Stop()
	gui.NewApplication().Start()
}
