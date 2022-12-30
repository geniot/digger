PROGRAM_NAME := digger

all: clean build

clean:
	rm bin/${PROGRAM_NAME}* -f

build:
	go build -gcflags="all=-N -l" -o bin/${PROGRAM_NAME} github.com/geniot/${PROGRAM_NAME}/cmd/${PROGRAM_NAME}

mips:
	CC='/opt/gcw0-toolchain/usr/bin/mipsel-gcw0-linux-uclibc-gcc' \
	 CGO_ENABLED=1 \
	 CGO_LDFLAGS='-lSDL2 -lpng16' \
	 GOARCH=mipsle \
	 GOMIPS=softfloat \
	 GOOS=linux \
	 PKG_CONFIG='/opt/gcw0-toolchain/usr/bin/pkg-config' \
	 go build -o bin/${PROGRAM_NAME}.gcw github.com/geniot/${PROGRAM_NAME}/cmd/${PROGRAM_NAME}

squash:
	mksquashfs bin/${PROGRAM_NAME}.gcw resources/media/${PROGRAM_NAME}.png resources/default.gcw0.desktop bin/${PROGRAM_NAME}.opk -all-root -no-xattrs -noappend -no-exports

opk: clean mips squash

#on PG2 use opkrun


