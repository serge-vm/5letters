package quizsolver

import (
	"bufio"
	"log"
	"strings"

	"github.com/serge-vm/5letters/assets"
)

func Solve(unordered []string, ordered map[int]string, absent []string) []string {
	suggestions := []string{}

	file, err := assets.Assets.Open("russian5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		word5 := scanner.Text()

		if !checkUnorderedLetters(word5, unordered) {
			continue
		}

		if !checkOrderedLetters(word5, ordered) {
			continue
		}

		if !checkAbsentLetters(word5, absent) {
			continue
		}

		suggestions = append(suggestions, word5)

	}

	return suggestions
}

func checkUnorderedLetters(word string, unorderedLetters []string) bool {
	for _, l := range unorderedLetters {
		if !strings.Contains(word, l) {
			return false
		}
	}
	return true
}

func checkOrderedLetters(word string, orderedLettes map[int]string) bool {
	runes5 := []rune(word)
	for p, l := range orderedLettes {
		rl := []rune(l)
		if runes5[p-1] != rl[0] {
			return false
		}
	}
	return true
}

func checkAbsentLetters(word string, absentLetters []string) bool {
	for _, l := range absentLetters {
		if strings.Contains(word, l) {
			return false
		}
	}
	return true
}
