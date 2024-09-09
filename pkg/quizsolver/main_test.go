package quizsolver

import (
	"testing"

	"github.com/serge-vm/5letters/internal/models"
	"github.com/stretchr/testify/assert"
)

func TestSolver(t *testing.T) {
	testCases := []struct {
		name       string
		unordered  []models.Unordered
		ordered    map[int]string
		absent     []string
		targetWord string
		found      bool
	}{
		{
			name: "Трава present",
			unordered: []models.Unordered{
				{P: 1, L: "а"},
			},
			ordered:    map[int]string{1: "т", 4: "в"},
			absent:     []string{"ю"},
			targetWord: "трава",
			found:      true,
		},
		{
			name: "Трава absent",
			unordered: []models.Unordered{
				{P: 3, L: "а"},
			},
			ordered:    map[int]string{1: "т", 4: "в"},
			absent:     []string{"ю"},
			targetWord: "трава",
			found:      false,
		},
		{
			name: "Метро present",
			unordered: []models.Unordered{
				{P: 1, L: "о"},
			},
			ordered:    map[int]string{1: "м", 3: "т"},
			absent:     []string{"к"},
			targetWord: "метро",
			found:      true,
		},
		{
			name: "Метро absent",
			unordered: []models.Unordered{
				{P: 1, L: "о"},
			},
			ordered:    map[int]string{1: "м", 3: "т"},
			absent:     []string{"р"},
			targetWord: "метро",
			found:      false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			var found bool
			result := Solve(tt.unordered, tt.ordered, tt.absent)
			for _, r := range result {
				if r == tt.targetWord {
					found = true
				}
			}
			assert.Equal(t, tt.found, found)
		})
	}
}

func TestUnordered(t *testing.T) {
	testCases := []struct {
		name      string
		unordered []models.Unordered
		checkWord string
		found     bool
	}{
		{
			name: "е not 1 in тесто",
			unordered: []models.Unordered{
				{P: 1, L: "е"},
			},
			checkWord: "тесто",
			found:     true,
		},
		{
			name: "е 2 in тесто",
			unordered: []models.Unordered{
				{P: 2, L: "е"},
			},
			checkWord: "тесто",
			found:     false,
		},
		{
			name: "a 1 in тесто",
			unordered: []models.Unordered{
				{P: 1, L: "а"},
			},
			checkWord: "тесто",
			found:     false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.found, checkUnorderedLetters(tt.checkWord, tt.unordered))
		})
	}
}

func TestOrdered(t *testing.T) {
	testCases := []struct {
		name      string
		ordered   map[int]string
		checkWord string
		found     bool
	}{
		{
			name:      "е 2 in тесто",
			ordered:   map[int]string{2: "е"},
			checkWord: "тесто",
			found:     true,
		},
		{
			name:      "е 1 in тесто",
			ordered:   map[int]string{1: "е"},
			checkWord: "тесто",
			found:     false,
		},
		{
			name:      "ы 1 in тесто",
			ordered:   map[int]string{1: "ы"},
			checkWord: "тесто",
			found:     false,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.found, checkOrderedLetters(tt.checkWord, tt.ordered))
		})
	}
}

func TestAbsent(t *testing.T) {
	testCases := []struct {
		name      string
		absent    []string
		checkWord string
		found     bool
	}{
		{
			name:      "а not in тесто",
			absent:    []string{"a"},
			checkWord: "тесто",
			found:     false,
		},
		{
			name:      "е in тесто",
			absent:    []string{"е"},
			checkWord: "тесто",
			found:     true,
		},
	}

	for _, tt := range testCases {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, !tt.found, checkAbsentLetters(tt.checkWord, tt.absent))
		})
	}
}
