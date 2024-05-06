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
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "е"})
	result := checkUnorderedLetters("тесто", unorderedLetters)
	if result == false {
		t.Fatal("Letter е is present in word тесто and not first. Misdetection.")
	}
}

func TestUnorderedLetterExclude(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 2, L: "е"})
	result := checkUnorderedLetters("тесто", unorderedLetters)
	if result == true {
		t.Fatal("Letter е is present in word тесто and position is right. Misdetection.")
	}

}

func TestUnorderedLetterAbsent(t *testing.T) {
	var unorderedLetters []models.Unordered
	unorderedLetters = append(unorderedLetters, models.Unordered{P: 1, L: "а"})
	result := checkUnorderedLetters("тесто", unorderedLetters)
	if result == true {
		t.Fatal("Letter в is absent in word тесто. Misdetection.")
	}

}

func TestOrderedLetterFound(t *testing.T) {
	letters := make(map[int]string)
	letters[2] = "е"
	result := checkOrderedLetters("тесто", letters)
	if result == false {
		t.Fatal("Letter е not found in word тесто on position 2, but it should.")
	}
}

func TestOrderedLetterWrongPosition(t *testing.T) {
	letters := make(map[int]string)
	letters[1] = "е"
	result := checkOrderedLetters("тесто", letters)
	if result == true {
		t.Fatal("Letter е found in word тесто on position 1, but it shouldn't.")
	}
}

func TestOrderedLetterNoLetter(t *testing.T) {
	letters := make(map[int]string)
	letters[1] = "ы"
	result := checkOrderedLetters("тесто", letters)
	if result == true {
		t.Fatal("Letter ы found in word тесто on position 1, but it shouldn't.")
	}
}

func TestAbsentLettersAbsent(t *testing.T) {
	result := checkAbsentLetters("тесто", []string{"а"})
	if result == false {
		t.Fatal("Letter а found in word тесто, but it shouldn't.")
	}
}

func TestAbsentLettersPresent(t *testing.T) {
	result := checkAbsentLetters("тесто", []string{"е"})
	if result == true {
		t.Fatal("Letter е not found in word тесто, but it should.")
	}
}
