package main

import "testing"

func TestAnagram(t *testing.T) {
	input := []string{"ПЯТКА", "пяТак", "тяпка", "листОк", "Слиток", "свисток", "столИК", "стол"}
	wantOutput := map[string][]string{
		"листок": {"слиток", "столик"},
		"пятка":  {"пятак", "тяпка"},
	}

	got := SearchAnagram(prepareWords(input))

	//compare got and want
	for k, v := range got {
		if len(v) != len(wantOutput[k]) {
			t.Errorf("Got [%s] want [%s]\n", v, wantOutput[k])
		}
		for i, x := range wantOutput[k] {
			if v[i] != x {
				t.Errorf("Got [%s] want [%s]\n", v[i], x)
			}
		}
	}

}
