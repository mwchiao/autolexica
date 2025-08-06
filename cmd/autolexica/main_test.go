package main

import (
	"slices"
	"testing"

	"github.com/mwchiao/autolexica/internal/board"
)

func TestMain(t *testing.T) {
	// Test that generated wordlist matches expected word list
	expected := []string{
		"abc",
		"aei",
		"dei",
		"feg",
		"ceg",
		"ghi",
		"ghif",
		"efhi",
		"abced",
	}
	slices.Sort(expected)

	dictionary := map[string]bool{
		"abc":   true,
		"aei":   true,
		"dei":   true,
		"feg":   true,
		"ceg":   true,
		"ghi":   true,
		"abcd":  true,
		"ghif":  true,
		"efhi":  true,
		"abcde": true,
		"abced": true,
		"defgh": true,
		"zoo":   true,
		"buzz":  true,
		"foo":   true,
		"bar":   true,
	}

	boardLetters := []string{
		"a", "b", "c",
		"d", "e", "f",
		"g", "h", "i",
	}

	letters := stringsToRunes(boardLetters)

	grid, _ := board.PopulateBoard(letters)
	words := findWords(&grid)
	var got []string
	for _, word := range words {
		if _, ok := dictionary[word]; ok {
			if !slices.Contains(got, word) {
				got = append(got, word)
			}
		}
	}
	slices.Sort(got)

	if !slices.Equal(got, expected) {
		t.Log("Generated solution does not match expected solution")
		t.Errorf("\tExpected: %v\n\tGot: %v", expected, got)
	}
}

func stringsToRunes(strs []string) []rune {
	runes := make([]rune, 0, len(strs))
	for _, s := range strs {
		runes = append(runes, []rune(s)[0])
	}

	return runes
}
