PROGRAM_NAME := digger

all: clean build

clean:
	rm bin/${PROGRAM_NAME}* -f

build:
	go build -o bin/${PROGRAM_NAME} geniot.com/geniot/${PROGRAM_NAME}/cmd/${PROGRAM_NAME}

mips:
	CC='/opt/gcw0-toolchain/usr/bin/mipsel-gcw0-linux-uclibc-gcc' \
	CGO_CFLAGS='-I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include -I/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/include/libpng16 -D_REENTRANT' \
	 CGO_ENABLED=1 \
	 CGO_LDFLAGS='-L/opt/gcw0-toolchain/usr/mipsel-gcw0-linux-uclibc/sysroot/usr/lib -lSDL2 -lpng16' \
	 GOARCH=mipsle \
	 GOMIPS=softfloat \
	 GOOS=linux \
	 PKG_CONFIG='/opt/gcw0-toolchain/usr/bin/pkg-config' \
	 go build -o bin/${PROGRAM_NAME}.gcw geniot.com/geniot/${PROGRAM_NAME}/cmd/${PROGRAM_NAME}

squash:
	mksquashfs bin/${PROGRAM_NAME}.gcw resources/media/${PROGRAM_NAME}.png resources/default.gcw0.desktop bin/${PROGRAM_NAME}.opk -all-root -no-xattrs -noappend -no-exports

opk: clean mips squash

#on PG2 use opkrun


