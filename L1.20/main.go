package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"unicode"
)

func FirstMethod() {
	scanner := bufio.NewScanner(os.Stdin)
	var line []rune
	for scanner.Scan() {
		line = []rune(" " + scanner.Text())
		break
	}

	length := len(line)
	lastSpacePos := length
	for i := length - 1; i >= 0; i-- {
		if unicode.IsSpace(line[i]) || i == (0) {
			for j := i + 1; j < lastSpacePos; j++ {
				line = append(line, line[j])
			}
			line = append(line, rune(' '))
			line = slices.Delete(line, i, lastSpacePos)
			lastSpacePos = i
		}
	}

	fmt.Println(string(line))
} //Время: O(2n) ~> O(n) | Память... - страдает))

func SecondMethod() {
	revers := func(line []rune, strt, end int) {
		for i, j := strt, end; i < j; i, j = i+1, j-1 {
			line[i], line[j] = line[j], line[i]
		}
	}

	scanner := bufio.NewScanner(os.Stdin)
	var line []rune
	for scanner.Scan() {
		line = []rune(scanner.Text())
		break
	}

	length := len(line)
	revers(line, 0, length-1)

	position := 0
	for i := 0; i <= length; i++ {
		if i == length || unicode.IsSpace(line[i]) {
			revers(line, position, i-1)
			position = i + 1
		}
	}
	fmt.Println(string(line))

} //Время: O(n) | Память: O(1)

func main() {
	FirstMethod()
	SecondMethod()
}
