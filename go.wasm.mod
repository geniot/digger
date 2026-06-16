module digger

go 1.26.1

require (
	github.com/BrownNPC/Raylib-Go-Wasm/raylib v0.0.0-20260421110350-7c24b2d5e6d3
	github.com/gen2brain/raylib-go/raygui v0.0.0-20250617194346-eddd038123ee
	github.com/gen2brain/raylib-go/raylib v0.0.0-20250504022611-e6017e5fc409
)

require (
	github.com/BrownNPC/Raylib-Go-Wasm/wasm-runtime v0.0.0-20260421110350-7c24b2d5e6d3 // indirect
	github.com/BrownNPC/wasm-ffi-go v1.3.0 // indirect
)

replace (
	github.com/BrownNPC/Raylib-Go-Wasm/wasm-runtime => ./Raylib-Go-Wasm/wasm-runtime
	github.com/gen2brain/raylib-go/raygui => ./Raylib-Go-Wasm/raygui
	github.com/gen2brain/raylib-go/raylib => ./Raylib-Go-Wasm/raylib
)
