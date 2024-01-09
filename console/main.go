package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/serge-vm/5letters/internal/quizsolver"
)

func consoleReadLetter(reader bufio.Reader) (string, bool) {
	letter, err := reader.ReadString('\n')
	if err != nil {
		log.Println("Ошибка ввода: ", err)
		return "", false
	}
	clean := strings.TrimSpace(letter)
	if len(clean) == 0 {
		return "", false
	}
	return string([]rune(clean)[0]), true
}

func consoleReadPosition(reader bufio.Reader) (int, bool) {
	strp, ok := consoleReadLetter(reader)
	if !ok {
		return 0, false
	}

	intp, err := strconv.Atoi(strp)
	if err != nil {
		log.Println("Число не распознано: ", err)
		return 0, false
	}

	return intp, true
}

func main() {
	unorderedLetters := []string{}
	orderedLetters := map[int]string{}
	absentLetters := []string{}
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("В слове есть буква (Пусто - следующий шаг): ")
		letter, ok := consoleReadLetter(*reader)
		if !ok {
			break
		}

		fmt.Print("Её позиция (Пусто - неизвестно): ")
		position, ok := consoleReadPosition(*reader)
		if ok {
			orderedLetters[position] = letter
		} else {
			unorderedLetters = append(unorderedLetters, letter)
		}
	}

	for {
		fmt.Print("В слове нет буквы (Пусто - закончить): ")
		letter, ok := consoleReadLetter(*reader)
		if !ok {
			break
		}

		absentLetters = append(absentLetters, letter)

	}

	if len(unorderedLetters) == 0 && len(orderedLetters) == 0 && len(absentLetters) == 0 {
		fmt.Println("Ничего не введено, выходим.")
		return
	}

	goodWords := quizsolver.Solve(unorderedLetters, orderedLetters, absentLetters)

	if len(goodWords) >= 1 {
		for _, gword := range goodWords {
			fmt.Println(gword)
		}
	} else {
		fmt.Println("Подходящих слов не нашлось.")
	}

	fmt.Print("Нажмите Enter Для завершения")
	reader.ReadBytes('\n')

}
