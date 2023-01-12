# About

[Digger](https://en.wikipedia.org/wiki/Digger_(video_game)) was released back in 1983.
I was two years old and only learned about this game when I was eight.
We didn't have a computer at home but my father had it at work.
So it was the first computer game that I learned to play.

It was a black-and-white IBM 8080 and I didn't know digger could shoot,
so I was just driving away each time trying to collect all the emeralds,
finding the optimal escape route.

40 years later (it's 2023 now) the game is still a lot of fun.
It actually inspired me to create this clone.

![Digger](screenshots/main.png)

This clone is a complete rewrite,
although I used the original graphics and sounds.

From the player's perspective the main difference with the original game is that this Digger has pixel-perfect
precision:

- all animations are smooth (except sprites - which I borrowed from the original 1983 Windmill game)
- all object collisions are pixel-based and calculated within a [Resolv](github.com/solarlune/resolv) Space
- collision with the field is based on [SDL_GetRGBA](https://wiki.libsdl.org/SDL2/SDL_GetRGBA)

There are also some minor functional differences:

- gold can stack on top of each other
- gold doesn't disappear within one session (while digger is alive)

I thought it makes more sense, at least for me.

# Controls

I mostly play it on my handheld consoles, so mouse is not used in the desktop versions.

Only keys:

| Desktop    | PocketGo2v2 | Function             |
|------------|-------------|----------------------|
| Left Ctrl  | A           | fire                 |
| Arrow Keys | ←,↑,→,↓     | move selection frame |  
| q          | L1+Start    | quit                 |  

# Motivation

Apart from the nostalgia that I'm having from time to time I started this project to learn [Go](https://go.dev/), [SDL2](https://www.libsdl.org/), some nuances of embedded
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
5. Digger.org: https://www.digger.org/
6. Wikipedia article: https://en.wikipedia.org/wiki/Digger_(video_game)

# Profile CPU

1. Add `defer profile.Start(profile.CPUProfile).Stop()` as the first line in main
2. Add import for it: `import "github.com/pkg/profile"`
3. Run and close the app, `cpu.pprof` will be created in the temporary directory
4. Install `go install github.com/google/pprof@latest` if it is not yet installed
5. Run on the command line: `pprof -top bin/digger_debug.exe <absolute_path_to>/cpu.pprof`
6. You will see a report with method calls that cause most CPU consumption