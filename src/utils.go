package main

import (
	gui "github.com/gen2brain/raylib-go/raygui"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	ZERO_VECTOR2 = rl.Vector2{}
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

func IfInt(cond bool, vTrue int32, vFalse int32) int32 {
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
	setTextStyle(40, 10, gui.TEXT_ALIGN_LEFT, 20)
}

func setTextStyle(ts gui.PropertyValue, sp gui.PropertyValue, ta gui.PropertyValue, pp gui.PropertyValue) {
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SIZE, ts)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_SPACING, sp)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT, ta)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_PADDING, pp)
	gui.SetStyle(gui.DEFAULT, gui.TEXT_ALIGNMENT_VERTICAL, gui.TEXT_ALIGN_CENTER)
	gui.SetStyle(gui.TEXTBOX, gui.TEXT_ALIGNMENT, gui.TEXT_ALIGN_LEFT)
}

func GetNextSpritePointerAndInc(currentPointer int, spritePointerInc int, spritesLen int) (int, int) {
	nextSpritePointerInc := spritePointerInc
	nextSpritePointer := currentPointer + spritePointerInc
	if nextSpritePointer >= spritesLen {
		nextSpritePointer = spritesLen - 1
	}
	if nextSpritePointer < 0 {
		nextSpritePointer = 0
	}
	if nextSpritePointer == spritesLen-1 || nextSpritePointer == 0 {
		nextSpritePointerInc = -spritePointerInc
	}
	return nextSpritePointer, nextSpritePointerInc
}
