package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"unicode"
)

var ignored = map[rune]struct{}{
	'(': struct{}{},
	')': struct{}{},
}

type charFrequency struct {
	char  rune
	count int
}

type byFrequency []charFrequency

func (a byFrequency) Len() int           { return len(a) }
func (a byFrequency) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byFrequency) Less(i, j int) bool { return a[i].count > a[j].count }

func processTextFile(filePath string) (string, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var text strings.Builder
	for scanner.Scan() {
		text.WriteString(scanner.Text())
	}
	return text.String(), scanner.Err()
}

func processDirectory(directoryPath string) (string, error) {
	var allText strings.Builder
	err := filepath.Walk(directoryPath, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			fileText, err := processTextFile(path)
			if err != nil {
				return err
			}
			allText.WriteString(fileText)
		}
		return nil
	})
	return allText.String(), err
}

func printCharacterFrequencies(text string) {
	// Remove spaces
	cleanedText := strings.Map(func(r rune) rune {
		if unicode.IsSpace(r) {
			return -1
		}
		return r
	}, text)

	// Count character frequencies
	characterFrequencies := make(map[rune]int)
	for _, char := range cleanedText {
		characterFrequencies[char]++
	}

	// Sort characters by frequencies in descending order
	var frequencies []charFrequency
	for char, freq := range characterFrequencies {
		if _, ignore := ignored[char]; !ignore {
			frequencies = append(frequencies, charFrequency{char, freq})
		}
	}
	sort.Sort(byFrequency(frequencies))

	// Print sorted characters
	for _, cf := range frequencies {
		fmt.Printf("%c: %d\n", cf.char, cf.count)
	}
}

func main() {
	// Specify the directory containing text files
	directoryPath := "./txt"

	allText, err := processDirectory(directoryPath)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printCharacterFrequencies(allText)
}
