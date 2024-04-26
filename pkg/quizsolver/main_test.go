package quizsolver

import (
	"testing"

	"github.com/serge-vm/5letters/internal/models"
)

func TestSolverFindGrass(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "а"})
	letters := make(map[int]string)
	letters[4] = "в"
	letters[1] = "т"
	result := Solve(unorderedLetters, letters, []string{"ю"})
	for _, r := range result {
		if r == "трава" {
			return
		}
	}
	t.Fatal("Word трава not found, but it should be.")
}

func TestSolverFindNoGrass(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 3, L: "а"})
	letters := make(map[int]string)
	letters[4] = "в"
	letters[1] = "т"
	result := Solve(unorderedLetters, letters, []string{"ю"})
	for _, r := range result {
		if r == "трава" {
			t.Fatal("Word трава found, but it shouldn't be.")
		}
	}
}

func TestSolverFindMetro(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "о"})
	letters := make(map[int]string)
	letters[3] = "т"
	letters[1] = "м"
	result := Solve(unorderedLetters, letters, []string{"к"})
	for _, r := range result {
		if r == "метро" {
			return
		}
	}
	t.Fatal("Word метро not found, but it should be.")
}

func TestSolverFindNoMetro(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "о"})
	letters := make(map[int]string)
	letters[3] = "т"
	letters[1] = "м"
	result := Solve(unorderedLetters, letters, []string{"р"})
	for _, r := range result {
		if r == "метро" {
			t.Fatal("Word метро found, but it shouldn't be.")
		}
	}
}

func TestUnorderedLetterFound(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "e"})
	result := checkUnorderedLetters("tests", unorderedLetters)
	if result == false {
		t.Fatal("Letter e is present in word tests and not first. Misdetection.")
	}
}

func TestUnorderedLetterExclude(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 2, L: "e"})
	result := checkUnorderedLetters("tests", unorderedLetters)
	if result == true {
		t.Fatal("Letter e is present in word tests and position is right. Misdetection.")
	}

}

func TestUnorderedLetterAbsent(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "a"})
	result := checkUnorderedLetters("tests", unorderedLetters)
	if result == true {
		t.Fatal("Letter a is absent in word tests. Misdetection.")
	}

}

func TestOrderedLetterFound(t *testing.T) {
	letters := make(map[int]string)
	letters[2] = "e"
	result := checkOrderedLetters("tests", letters)
	if result == false {
		t.Fatal("Letter e not found in word tests on position 2, but it should.")
	}
}

func TestOrderedLetterWrongPosition(t *testing.T) {
	letters := make(map[int]string)
	letters[1] = "e"
	result := checkOrderedLetters("tests", letters)
	if result == true {
		t.Fatal("Letter e found in word tests on position 1, but it shouldn't.")
	}
}

func TestOrderedLetterNoLetter(t *testing.T) {
	letters := make(map[int]string)
	letters[1] = "a"
	result := checkOrderedLetters("tests", letters)
	if result == true {
		t.Fatal("Letter a found in word tests on position 1, but it shouldn't.")
	}
}

func TestAbsentLettersAbsent(t *testing.T) {
	result := checkAbsentLetters("tests", []string{"a"})
	if result == false {
		t.Fatal("Letter a found in word tests, but it shouldn't.")
	}
}

func TestAbsentLettersPresent(t *testing.T) {
	result := checkAbsentLetters("test", []string{"e"})
	if result == true {
		t.Fatal("Letter e not found in word tests, but it should.")
	}
}
