# About

![Digger](screenshots/main.png)

[Digger](https://www.digger.org/) is a classic game. This clone is a complete rewrite of it,
although I used the original graphics and sounds.

From the player's perspective the main difference with the original game is that this Digger has pixel-perfect
precision:

- all animations are smooth (except sprites - which I borrowed from the original 1983 Windmill game)
- all object collisions are pixel-based and calculated within a [Resolv](github.com/solarlune/resolv) Space
- collision with the field is based on [SDL_GetRGBA](https://wiki.libsdl.org/SDL2/SDL_GetRGBA)

# Motivation

I started this project to learn [Go](https://go.dev/), [SDL2](https://www.libsdl.org/), some nuances of embedded
programming and game development.

# Build

You need go-sdl2 bindings: https://github.com/veandco/go-sdl2

You can find detailed instructions in the [README.md](https://github.com/veandco/go-sdl2/blob/master/README.md) of
go-sdl2.

In general on Windows you would need to install Go and MinGW.
Download SDL2 packages with DLLs and header files.
Put DLLs and header files in MinGW folders.

On Linux you would need to install sdl2-dev packages.

On Windows to build the exe file with static linking I use:

`go build -ldflags "-s -w -H=windowsgui" -tags static -o bin\digger.exe github.com/geniot/digger/src/cmd/digger`

I use [GoLand](https://www.jetbrains.com/go/) for development.

To build the OPK file for my [PocketGo2](https://wagnerstechtalk.com/pg2/) I use make:

`make opk`

See the Makefile. It can only be done on Linux.

# Testing

I only tested the game manually in the following environments:

- Windows 10 (desktop)
- LXLE Linux (desktop)
- arm64 (embedded Linux on my handheld console RK2020)
- mipsel (embedded Linux on my handheld console PocketGo2)

# Links

1. OpenDingux Software list: https://github.com/retrogamehandheld/OpenDingux
2. Go SDL2 bindings: https://github.com/veandco/go-sdl2
3. PocketGo Firmware that I use: https://github.com/Ninoh-FOX/POCKETGO2_ROGUE_CFW
4. Toolchain for it: https://github.com/Ninoh-FOX/toolchain

# Profile CPU

1. Add `defer profile.Start(profile.CPUProfile).Stop()` as the first line in main
2. Add import for it: `import "github.com/pkg/profile"`
3. Run and close the app, `cpu.pprof` will be created in the temporary directory
4. Install `go install github.com/google/pprof@latest` if it is not yet installed
5. Run on the command line: `pprof -top bin/digger_debug.exe <absolute_path_to>/cpu.pprof`
6. You will see a report with method calls that cause most CPU consumption