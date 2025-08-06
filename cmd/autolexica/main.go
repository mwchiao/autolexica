package main

import (
	"flag"
	"fmt"
	"log"
	"slices"

	"github.com/mwchiao/autolexica/internal/board"
	"github.com/mwchiao/autolexica/internal/utilities"
)

const minWordLength = 3

var gridPath string
var dictPath string
var saveFile bool

func init() {
	flag.StringVar(&gridPath, "grid", "./grid.txt", "Specify the path to your desired grid.txt file")
	flag.StringVar(&dictPath, "dictionary", "./dictionary.json", "Specify the path to your desired dictionary.json file")
	flag.BoolVar(&saveFile, "save", false, "Whether or not to save solution as a file to disk")
}

func main() {
	flag.Parse()

	letters, err := utilities.ReadGrid(gridPath)
	if err != nil {
		log.Fatal(err)
	}

	grid, err := board.PopulateBoard(letters)
	if err != nil {
		log.Fatal(err)
	}

	dictionary, err := utilities.GetDictionary(dictPath)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Finding solutions ")
	words := findWords(&grid)

	var validWords []string
	for _, word := range words {
		if _, ok := dictionary[word]; ok {
			if !slices.Contains(validWords, word) {
				validWords = append(validWords, word)
			}
		}
	}
	slices.Sort(validWords)

	if saveFile {
		err = utilities.SaveToFile(validWords)
		if err != nil {
			log.Print(err)
		}
	} else {
		for _, word := range validWords {
			fmt.Println(word)
		}
	}

	fmt.Println("Done")
}

func findWords(grid *[]board.Tile) []string {
	gridSize := len(*grid)
	visited := make([]bool, gridSize)

	var foundWords []string
	for i := range gridSize {
		// Find all words with each letter of the grid as the first letter
		words := searchGrid(grid, i, &visited)
		foundWords = append(foundWords, words...)
		// reset visited after iterating over each tile
		visited = make([]bool, gridSize)
	}

	return foundWords
}

func searchGrid(grid *[]board.Tile, index int, visited *[]bool) []string {
	var words []string
	words = append(words, dfs(grid, index, visited, []rune{})...)

	return words
}

func dfs(grid *[]board.Tile, index int, visited *[]bool, runes []rune) []string {
	var words []string

	(*visited)[index] = true
	runes = append(runes, (*grid)[index].Letter)
	if len(runes) >= minWordLength {
		words = append(words, string(runes))
	}

	for _, n := range (*grid)[index].Neighbors {
		if (*visited)[n] == false {
			words = append(words, dfs(grid, n, visited, runes)...)
		}
	}

	(*visited)[index] = false
	runes = runes[:(len(runes) - 1)]
	return words
}
