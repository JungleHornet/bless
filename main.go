package bless

import (
	"fmt"
	"golang.org/x/term"
	"os"
	"strings"
)

type settings = struct {
	selectorOpen  string
	selectorClose string
	frame         rune
}

type terminal = struct {
	width, height int
	frame         string
	pastFrame     []string
	innerFrame    []string
	futureFrame   []string
}

type Blessing struct {
	settings settings
	terminal terminal
}

func (b *Blessing) updateBlessingSize() {
	w, h, err := term.GetSize(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	if b.terminal.width != w {
		b.terminal.width = w
	}
	if b.terminal.height != h {
		b.terminal.height = h
	}
}

func New(selectorOpen string, selectorClose string, frame rune) Blessing {
	newBlessing := Blessing{settings: settings{selectorOpen: selectorOpen, selectorClose: selectorClose, frame: frame}, terminal: terminal{}}

	newBlessing.terminal.innerFrame = append(newBlessing.terminal.innerFrame, "")
	newBlessing.updateBlessingSize()
	newBlessing.runFrame()

	return newBlessing
}

func (b *Blessing) Close() {
	fmt.Println(strings.Repeat(string(8), len(b.terminal.frame)-(b.terminal.width*2)+1))
}
