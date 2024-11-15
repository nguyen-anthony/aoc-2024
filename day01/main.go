package main

import (
    "log"
    "strconv"
    "github.com/nguyen-anthony/advent-of-code/utils"
    "unicode"
)

func FirstAndLastDigit(s string) (int, error) {
    var first, last rune
    foundFirst := false

    for _, char := range s {
        if unicode.IsDigit(s) {
            if !foundFirst {
                first = char
                foundFirst = true
            }
            last = char
        }
    }

    if !foundFirst {
        return 0, fmt.Errorf("no digits in the string")
    }

    firstInt, err := strconv.Atoi(string(first))
    if err != nil {
		return 0, fmt.Errorf("error converting first digit: %v", err)
	}

	lastInt, err := strconv.Atoi(string(last))
	if err != nil {
		return 0, fmt.Errorf("error converting last digit: %v", err)
	}

    result := firstInt*10 + lastInt
	return result, nil
}

func main() {
    // Use a function from the utils package
    lines, err := utils.ReadLines("input.txt")
    if err != nil {
        log.Fatalf("Failed to read file: %v", err)
    }

    fmt.Println("File lines:", lines)

    total := 0
    for _, line := range lines {
        number, _ := FirstAndLastDigit(line)
        if err != nil {
            fmt.Println(err)
        }
        total += number
    }

    fmt.Println("Total:", total)
}
