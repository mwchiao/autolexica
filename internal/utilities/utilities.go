package utilities

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

func SaveToFile(solution []string) error {

	fmt.Println("Saving to file \"solution.txt\"")
	f, err := os.Create("solution.txt")
	if err != nil {
		return err
	}

	defer f.Close()

	for _, line := range solution {
		_, err := f.WriteString(line + "\n")
		if err != nil {
			return err
		}
	}

	fmt.Println("Solution saved to file")
	return nil
}

func ReadGrid(path string) ([]rune, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	// Would using []byte instead save memory?
	var letters []rune
	for line := range strings.SplitSeq(string(data), "\n") {
		trimmed := strings.TrimSpace(line)
		if len(trimmed) == 0 {
			continue
		}

		for letter := range strings.SplitSeq(trimmed, " ") {
			letters = append(letters, []rune(letter)...)
		}
	}

	return letters, nil
}

func GetDictionary(path string) (map[string]bool, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var dict []string
	err = json.Unmarshal(data, &dict)
	if err != nil {
		return nil, err
	}

	dictionary := make(map[string]bool)
	for _, word := range dict {
		dictionary[word] = true
	}

	return dictionary, nil
}
