package main

import (
	"strconv"
	"strings"
)

var openbracket rune
var closebracket rune

func init() {
	openbracket = '('
	closebracket = ')'
}

func unParseStringWhichHaveMultiplierToTemplateStrings(input string) string {
	start := make([]int, 0)

	multiplier := 1
	multiplierString := false

	val, err := strconv.Atoi(string(input[0]))
	if err != nil {
		//fmt.Println("Error converting to Int")
	} else {
		multiplier = val
		multiplierString = true
	}

	i := 0

	for i<len(input) {
		if rune(input[i]) == openbracket {
			start = append(start, i-1)
		} else if rune(input[i]) == closebracket {
			input = input[:start[len(start)-1]] + unParseStringWhichHaveMultiplierToTemplateStrings(input[start[len(start)-1]:i]) + input[i+1:]
			start = start[:len(start) - 1]
		}
		i++
	}
	if !multiplierString {
		return input
	}
	return strings.Repeat(input[2:len(input)], multiplier)
}
