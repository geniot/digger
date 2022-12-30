package loop

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/geniot/digger/internal/ctx"
	"github.com/tevino/abool/v2"
	"os"
	"time"
)

type LoopImpl struct {
	isRunning *abool.AtomicBool
}

func NewLoop() *LoopImpl {
	return &LoopImpl{abool.New()}
}

func (loop LoopImpl) Start() {
	go loop.StartInput()
	loop.isRunning.Set()
	for loop.isRunning.IsSet() {
		ctx.EventLoopIns.Run()
		ctx.PhysicsLoopIns.Run()
		ctx.RenderLoopIns.Run()
	}
}

func (loop LoopImpl) StartInput() {
	f, err := os.Open("/dev/input/event2")
	if err != nil {
		println(err.Error())
		return
	}
	defer f.Close()
	b := make([]byte, 24)
	for {
		f.Read(b)
		sec := binary.LittleEndian.Uint64(b[0:8])
		usec := binary.LittleEndian.Uint64(b[8:16])
		t := time.Unix(int64(sec), int64(usec))
		fmt.Println(t)
		var value int32
		typ := binary.LittleEndian.Uint16(b[16:18])
		code := binary.LittleEndian.Uint16(b[18:20])
		binary.Read(bytes.NewReader(b[20:]), binary.LittleEndian, &value)
		fmt.Printf("type: %x\ncode: %d\nvalue: %d\n", typ, code, value)
	}
}

func (loop LoopImpl) Stop() {
	loop.isRunning.UnSet()
}
