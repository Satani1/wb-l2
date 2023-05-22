package strings_sort

import (
	"fmt"
	"os"
	"strings"
	"unicode"
)

func DeleteDuplicate(data []string) []string {
	exist := make(map[string]struct{})
	result := make([]string, 0, len(data))

	for _, value := range data {
		if _, found := exist[value]; found {
			continue
		}

		result = append(result, value)
		exist[value] = struct{}{}
	}
	return result
}

func OutTerminal(data []string) {
	for _, value := range data {
		fmt.Fprintf(os.Stdout, "%v\n", value)
	}
}

func TrimNonNum(data string) string {
	return strings.TrimRightFunc(data, func(r rune) bool {
		return !unicode.IsNumber(r)
	})
}
