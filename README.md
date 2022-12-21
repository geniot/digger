# About
[Digger](https://www.digger.org/) is a classic game. This clone is a complete rewrite, 
although I used the original source code for reference.

The main difference with the original game is that this Digger has pixel perfect precision. 

- all animation is smooth (except sprites - which I borrowed)

# Build

You need go-sdl2 bindings: https://github.com/veandco/go-sdl2

You can find detailed instructions there.

In general on Windows you would need to install Go and MinGW. 
Download SDL2 packages with DLLs and header files. 
Put DLLs and header files in MinGW folders. 

On Linux you would install sdl2-dev packages.

On Windows to build the exe file with static linking I use: 

`go build -ldflags "-s -w -H=windowsgui" -tags static -o bin\digger.exe github.com/geniot/digger/cmd/digger`

I use [GoLand](https://www.jetbrains.com/go/) for development.

# Links
1. OpenDingux Software list: https://github.com/retrogamehandheld/OpenDingux
2. Go SDL2 bindings: https://github.com/veandco/go-sdl2

# Roadmap
1. Digger can fire. 
2. Fireball hits the wall.
2. Bag starts to shake. 
3. Bag can fall. Multiple bags can fall.
4. Bag can turn to gold or stay a bag.
3. Digger can collect gold.
4. Digger can move bags (one or many).
4. Digger can be killed by a bag.
5. Nobbin chases digger.
6. Nobbin can move bags.
7. Nobbin can be killed by bag or fire.
8. Nobbin can turn to Hobbin and back.
8. Hobbin chases digger.
9. Hobbin eats everything: bags, emeralds and field.
11. Hobbin can be killed by fire or a bag.
12. 