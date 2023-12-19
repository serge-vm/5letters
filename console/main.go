package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/serge-vm/5letters/assets"
)

func main() {
	just_letters := []string{}
	postion_letters := map[int]string{}
	reader := bufio.NewReader(os.Stdin)
	var letter, position string
	var err error

	for {
		fmt.Print("В слове есть буква (Пусто - закончить): ")
		letter, err = reader.ReadString('\n')
		if err != nil {
			log.Println("Ошибка ввода: ", err)
		}
		if len(strings.TrimSpace(letter)) == 0 {
			break
		}

		fmt.Print("Её позиция (Пусто - неизвестно): ")
		position, err = reader.ReadString('\n')
		if err != nil {
			log.Println("Ошибка ввода")
		}
		strl := strings.Trim(letter, "\r\n")
		strp := strings.Trim(position, "\r\n")
		if len(strp) > 0 {
			intp, err := strconv.Atoi(strp)
			if err != nil {
				log.Println("Число не распознано: ", err)
			} else {
				postion_letters[intp] = strl
			}
		} else {
			just_letters = append(just_letters, strl)
		}
	}

	if len(just_letters) == 0 && len(postion_letters) == 0 {
		fmt.Println("Ничего не введено, выходим.")
		return
	}

	// fmt.Println(postion_letters, just_letters)
	file, err := assets.Assets.Open("russian5.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var goodWord bool

	for scanner.Scan() {
		goodWord = true
		word5 := scanner.Text()

		for _, l := range just_letters {
			if !strings.Contains(word5, l) {
				goodWord = false
				break
			}
		}

		runes5 := []rune(word5)
		for p, l := range postion_letters {
			rl := []rune(l)
			if runes5[p-1] != rl[0] {
				goodWord = false
				break
			}
		}
		if goodWord {
			fmt.Println(word5)
		}

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Print("Нажмите Enter Для завершения")
	reader.ReadBytes('\n')

}
