package strings_sort

import "strconv"

type tableStrings struct {
	data    [][]string
	column  int
	numeric bool
}

func (ts tableStrings) Len() int {
	return len(ts.data)
}

func (ts tableStrings) Less(i, j int) bool {
	column := ts.column

	if column > len(ts.data[i])-1 || column > len(ts.data[j]) {
		column = 0
	}

	if ts.numeric {
		n1 := TrimNonNum(ts.data[i][column])
		n2 := TrimNonNum(ts.data[j][column])

		i1, err := strconv.Atoi(n1)
		if err != nil {
			return ts.data[i][column] < ts.data[j][column]
		}
		j1, err := strconv.Atoi(n2)
		if err != nil {
			return ts.data[i][column] < ts.data[j][column]
		}

		return i1 < j1
	}

	return ts.data[i][column] < ts.data[j][column]
}

func (ts tableStrings) Swap(i, j int) {
	ts.data[i], ts.data[j] = ts.data[j], ts.data[i]
}
