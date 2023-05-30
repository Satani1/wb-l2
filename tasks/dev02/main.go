package main

import (
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"unicode"
)

type AppUnpacker struct {
	input string
}

func (a *AppUnpacker) Unpack(str string) (string, error) {
	var b strings.Builder
	runes := []rune(str)

	for i := 0; i < len(runes); i++ {
		if unicode.IsDigit(runes[i]) {
			if i == 0 {
				return "", nil
			}

			var num strings.Builder
			num.WriteRune(runes[i])
			letter := runes[i-1]

			for j := i + 1; j < len(runes)-1 && unicode.IsDigit(runes[j]); j++ {
				num.WriteRune(runes[j])
				i++
			}

			result, err := strconv.Atoi(num.String())
			if errors.Is(err, strconv.ErrRange) {
				// if number out of range, just print it
				b.WriteString(num.String())
				continue
			}
			if err != nil {
				return "", nil
			}

			for j := 0; j < result-1; j++ {
				b.WriteRune(letter)
			}

			continue

		}
		_, err := b.WriteRune(runes[i])
		if err != nil {
			return "", nil
		}
	}
	return b.String(), nil
}

func main() {
	appStr := AppUnpacker{input: "a4bc2d5e"}

	res, err := appStr.Unpack(appStr.input)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(res)
}
