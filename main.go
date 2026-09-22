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


func main(){
	reader := bufio.NewReader(os.Stdin)
	fmt.Println(generateText(words))
	fmt.Print("> ")
	start := time.Now()
	text, _ := reader.ReadString('\n')
	elapsed := time.Since(start)
	text = strings.TrimSpace(text)
	wpm := float64(len(text) / 5) / float64(elapsed.Minutes())
	fmt.Println(wpm)
}

func generateText(words []string) string{	

	rand.Shuffle(len(words), func(i, j int) {
		words[i], words[j] = words[j], words[i]
	})

	generated_text := strings.Join(words, " ")

	return generated_text
}
