//go:build wasm
// +build wasm

package main

import rl "github.com/BrownNPC/Raylib-Go-Wasm/raylib"

func main() {
	application := NewApplication()
	var update = func() {
		application.Update()
	}
	rl.SetMainLoop(update)
	for !rl.WindowShouldClose() {
		update()
	}
	application.Exit()
}
