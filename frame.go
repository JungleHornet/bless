package bless

import (
	"fmt"
	"strings"
	"unicode/utf8"

	"github.com/lunixbochs/vtclean"
)

func (b *Blessing) Print(output ...any) int {
	var outputStr string
	for _, each := range output {
		outputStr += fmt.Sprint(each) + " "
	}
	outputStr = strings.TrimSuffix(outputStr, " ")
	line := b.write(outputStr, -1)
	b.runFrame()
	return line
}

func (b *Blessing) write(output string, startLine int) int {
	appending := false
	if startLine == -1 {
		startLine = b.getLine()
		appending = true
	}
	outputSlice := strings.Split(output, "\n")

	for startLine >= len(b.terminal.innerFrame) {
		b.terminal.innerFrame = append(b.terminal.innerFrame, "")
	}

	if appending {
		b.terminal.innerFrame[startLine] += outputSlice[0]
	} else {
		b.terminal.innerFrame[startLine] = outputSlice[0]
	}
	startLine += 1

	if len(outputSlice) > 1 {
		outputSlice = outputSlice[1:]
	} else {
		b.cleanTerminal()
		return startLine - 1
	}

	if startLine > b.getLine() || startLine == b.getLine() {
		if !appending && startLine == b.getLine() {
			for i, each := range outputSlice {
				b.terminal.innerFrame[startLine+i] = each

			}
		}
		b.terminal.innerFrame = append(b.terminal.innerFrame, outputSlice...)
	} else {
		for i := startLine; i < startLine+len(outputSlice); i++ {
			b.terminal.innerFrame[i] = outputSlice[0]
			outputSlice = outputSlice[1:]
		}
	}
	b.cleanTerminal()
	return startLine - 1
}

func printableLen(str string) int {
	return utf8.RuneCount([]byte(vtclean.Clean(str, false)))
}

func (b *Blessing) cleanTerminal() {
	for i, each := range b.terminal.innerFrame {
		if printableLen(each) > b.terminal.width-3 {
			this := each[:b.terminal.width-3]
			next := " " + each[b.terminal.width-3:]
			b.terminal.innerFrame = append(append(b.terminal.innerFrame[:i], this, next), b.terminal.innerFrame[i+1:]...)
		}
	}

	innerHeight := b.terminal.height - 2
	for len(b.terminal.innerFrame) > innerHeight {
		b.terminal.pastFrame = append(b.terminal.pastFrame, b.terminal.innerFrame[0])
		b.terminal.innerFrame = b.terminal.innerFrame[1:]
	}
}

func (b *Blessing) Println(output ...any) int {
	output = append(output, "\n")
	return b.Print(output...)
}

func (b *Blessing) Overwrite(startLine int, output ...any) {
	var outputStr string
	for _, each := range output {
		outputStr += fmt.Sprint(each) + " "
	}
	outputStr = strings.TrimSuffix(outputStr, " ")
	b.write(outputStr, startLine)
	b.runFrame()
}

func (b *Blessing) RmLines(lines ...int) {
	for i, x := range lines {
		x = x - i
		b.terminal.innerFrame = append(b.terminal.innerFrame[:x], b.terminal.innerFrame[x+1:]...)
	}
	b.runFrame()
}

func (b *Blessing) getLine() int {
	return len(b.terminal.innerFrame) - 1
}

func (b *Blessing) constructFrame() {
	f := string(b.settings.frame)
	frame := strings.Repeat(f, b.terminal.width-1) + "\n"
	for i := 0; i < (b.terminal.height - 1); i++ {
		var line string
		if i == b.terminal.height-2 {
			line = strings.Repeat(f, b.terminal.width-2)
		} else if len(b.terminal.innerFrame) > i {
			line = " " + b.terminal.innerFrame[i]
			if printableLen(line) < b.terminal.width-2 {
				line += strings.Repeat(" ", b.terminal.width-2-printableLen(line))
			}
		} else {
			line = strings.Repeat(" ", b.terminal.width-2)
		}
		line = f + line + f
		frame += line + "\n"
	}
	b.terminal.frame = strings.TrimSuffix(frame, "\n")
}

func (b *Blessing) printFrame() {
	fmt.Print(strings.Repeat("\b", b.terminal.height*b.terminal.width))
	fmt.Print(b.terminal.frame)
	fmt.Print(strings.Repeat("\b", (b.terminal.width*2)-1))
}

func (b *Blessing) runFrame() {
	b.updateBlessingSize()
	b.constructFrame()
	b.printFrame()
}
