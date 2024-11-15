package utils

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func ReadLines(filename string) ([]string, error) {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return strings.Split(strings.TrimSpace(string(content)), "\n"), nil
}

func ReadInts(filename string) ([]int, error) {
	lines, err := ReadLines(filename)
	if err != nil {
		return nil, err
	}
	ints := make([]int, len(lines))
	for i, line := range lines {
		var val int
		_, err := fmt.Sscanf(line, "%d", &val)
		if err != nil {
			return nil, err
		}
		ints[i] = val
	}
	return ints, nil
}

