package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
)

func orPanic(err interface{}) {
	switch v := err.(type) {
	case error:
		if v != nil {
			panic(err)
		}
	case bool:
		if !v {
			panic("condition failed: != true")
		}
	}
}

func If[T any](cond bool, vTrue, vFalse T) T {
	if cond {
		return vTrue
	}
	return vFalse
}

func IfInt(cond bool, vTrue int, vFalse int) int {
	if cond {
		return vTrue
	}
	return vFalse
}

func orPanicRes[T any](res T, err interface{}) T {
	orPanic(err)
	return res
}

func setDefaultTextStyle() {
	setTextStyle(40, 10, int64(gui.TEXT_ALIGN_LEFT), 20)
}

func setTextStyle(ts int64, sp int64, ta int64, pp int64) {
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, ts)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SPACING, sp)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT, ta)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_PADDING, pp)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT_VERTICAL, int64(gui.TEXT_ALIGN_CENTER))
	gui.SetStyle(gui.TEXTBOX, gui.TEXT_ALIGNMENT, int64(gui.TEXT_ALIGN_LEFT))
}
