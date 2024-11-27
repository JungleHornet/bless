package bless

import (
	"errors"
	"github.com/mattn/go-tty"
	"log"
	"strings"
)

func (b *Blessing) constructOptions(optionsArg []string, index int) (string, error) {
	options := make([]string, len(optionsArg))
	copy(options, optionsArg)
	if index > len(options) {
		return "", errors.New("error: index out of bounds")
	} else if index < 0 {
		return "", errors.New("error: index out of bounds")
	}

	for thisIndex, option := range options {
		option = strings.TrimSpace(option)

		if thisIndex == index {
			option = b.settings.selectorOpen + strings.ToUpper(option) + b.settings.selectorClose
		} else {
			option = " " + option + " "
		}
		options[thisIndex] = option
	}

	res := ""
	for _, this := range options {
		res += this
	}

	strings.TrimPrefix(res, " ")

	return res, nil
}

func (b *Blessing) HorizontalOptions(prompt string, options ...string) int {
	b.Println(prompt)
	tty, err := tty.Open()
	if err != nil {
		log.Fatal(err)
	}
	defer tty.Close()

	option := 0

	optionStr, _ := b.constructOptions(options, option)
	line := b.Print(optionStr)
	for {
		var r rune
		r, err = tty.ReadRune()
		if err != nil {
			log.Fatal(err)
		}

		{
			if r == 27 {
				var r2 rune
				r2, err = tty.ReadRune()
				if err != nil {
					log.Fatal(err)
				}
				if r2 == 91 {
					var r3 rune
					r3, err = tty.ReadRune()
					if err != nil {
						log.Fatal(err)
					}

					if r3 == 67 {
						if option != len(options)-1 {
							option++
						} else {
							option = 0
						}
					} else if r3 == 68 {
						if option != 0 {
							option--
						} else {
							option = len(options) - 1
						}
					} else {
						r = r3
						continue
					}
				} else {
					r = r2
					continue
				}
			} else if r == 13 {
				b.Println()
				return option
			}
		}

		optionStr, _ = b.constructOptions(options, option)
		b.Overwrite(line, optionStr)
	}
}

func (b *Blessing) VerticalOptions(prompt string, options ...string) int {
	// TODO make vertical options function that uses up and down arrow keys
	return 0
}

func (b *Blessing) MatrixOptions(prompt string, options ...string) int {
	// TODO make matrix options function that uses up, down, left, and right arrow keys
	return 0
}
