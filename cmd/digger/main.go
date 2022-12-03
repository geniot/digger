package main

import (
	"github.com/geniot/digger/gui"
)

func main() {
	gl := gui.GameLoop()
	gl.Start()
	println("Hello digger")
}
