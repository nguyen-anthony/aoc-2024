package utils

import (
	"fmt"
	"strings"
)

func SplitToInts(s, sep string) ([]int, error) {
	parts := strings.Split(s, sep)
	ints := make([]int, len(parts))
	for i, part := range parts {
		var val int
		_, err := fmt.Sscanf(part, "%d", &val)
		if err != nil {
			return nil, err
		}
		ints[i] = val
	}
	return ints, nil
}

