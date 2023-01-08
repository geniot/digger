PROGRAM_NAME := digger

all: clean build

clean:
	rm bin/${PROGRAM_NAME}* -f

build:
	go build -o bin/${PROGRAM_NAME} github.com/geniot/${PROGRAM_NAME}/src/cmd/${PROGRAM_NAME}

build-debug:
	go build -v -gcflags="all=-N -l" -o bin/${PROGRAM_NAME} github.com/geniot/${PROGRAM_NAME}/src/cmd/${PROGRAM_NAME}

debug:
	SDL_GAMECONTROLLERCONFIG="190000004b4800000010000000010000,GO-Advance Controller,a:b1,b:b0,back:b10,dpdown:b7,dpleft:b8,dpright:b9,dpup:b6,leftshoulder:b4,lefttrigger:b12,leftx:a0,lefty:a1,rightshoulder:b5,righttrigger:b13,start:b15,x:b2,y:b3,platform:Linux," \
	dlv --listen=:2345 --headless=true --api-version=2 --accept-multiclient exec ~/projects/${PROGRAM_NAME}/bin/${PROGRAM_NAME}

mips:
	 CC='/opt/gcw0-toolchain/usr/bin/mipsel-gcw0-linux-uclibc-gcc' \
	 CGO_ENABLED=1 \
	 CGO_LDFLAGS='-lSDL2 -lpng16' \
	 GOARCH=mipsle \
	 GOMIPS=softfloat \
	 GOOS=linux \
	 PKG_CONFIG='/opt/gcw0-toolchain/usr/bin/pkg-config' \
	 go build -o bin/${PROGRAM_NAME}.gcw github.com/geniot/${PROGRAM_NAME}/src/cmd/${PROGRAM_NAME}

squash:
	mksquashfs bin/${PROGRAM_NAME}.gcw src/res/media/${PROGRAM_NAME}.png src/res/default.gcw0.desktop bin/${PROGRAM_NAME}.opk -all-root -no-xattrs -noappend -no-exports

opk: clean mips squash

#on PG2 use opkrun


