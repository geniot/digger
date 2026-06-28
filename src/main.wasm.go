//go:build wasm
// +build wasm

package main

import rl "github.com/BrownNPC/Raylib-Go-Wasm/raylib"

func main() {
	application := NewApplication()
	accumulator := float64(0)
	lastTime := float64(0)
	tick := int64(0)
	var update = func() {
		t := rl.GetTime()
		dt := t - lastTime
		lastTime = t
		application.ProcessInput()
		for accumulator += dt; accumulator > TICK; accumulator -= TICK {
			application.Update(tick)
			tick++
		}
		application.Render()
	}
	rl.SetMainLoop(update)
	for !rl.WindowShouldClose() {
		update()
	}
	application.Exit()
}
