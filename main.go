package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strings"
	"time"
)

var words = []string{
	"the",
	"quick",
	"brown",
	"the",
	"hello",
	"world",
	"computer",
	"house",
	"small",
	"large",
	"program",
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	target := generateText(words)

	fmt.Println(target)
	fmt.Print("> ")

	start := time.Now()
	text, _ := reader.ReadString('\n')

	elapsed := time.Since(start)
	text = strings.TrimSpace(text)

	correct := checkText(target, text)

	wpm := (float64(correct) / 5.0) / elapsed.Minutes()

	accuracy := 0.0
	if len(text) > 0 {
		accuracy = float64(correct) / float64(len(text)) * 100
	}

	fmt.Printf("WPM: %.2f\n", wpm)
	fmt.Printf("Accuracy: %.2f%%\n", accuracy)
}

func generateText(words []string) string {
	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	return strings.Join(words, " ")
}

func checkText(target string, typed string) int {
	targetWords := strings.Fields(target)
	typedWords := strings.Fields(typed)

	correct := 0
	limit := min(len(targetWords), len(typedWords))

	for i := range limit {
		correct += checkWord(targetWords[i], typedWords[i])

		if i < limit - 1{
			correct++
		}
	}

	return correct
}

func checkWord(target string, typed string) int {
	correct := 0
	limit := min(len(target), len(typed))

	for i := range limit {
		if target[i] == typed[i] {
			correct++
		}
	}

	return correct
}
