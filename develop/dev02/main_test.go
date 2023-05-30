package main

import (
	"testing"
)

func TestUnpack(t *testing.T) {
	input := []string{"a4bc2d5e", "abcd", "", "qwe5"}
	wantOutput := []string{"aaaabccddddde", "abcd", "", "qweeeee"}

	for i, str := range input {
		app := AppUnpacker{input: str}
		got, err := app.Unpack(str)
		if err != nil {
			t.Error("Uncorrect data", err)
		}
		want := wantOutput[i]
		if got != want {
			t.Errorf("Got [%s] want [%s]\n", got, want)
		}
	}
}
