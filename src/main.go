//go:build !wasm
// +build !wasm

package main

func main() {
	application := NewApplication()
	for !application.ShouldExit() {
		application.Update()
	}
	application.Exit()
}
