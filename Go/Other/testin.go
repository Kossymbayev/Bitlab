package main

import (
	"fmt"
	"time"
)

func typewriter(text string) {
	for _, char := range text {
		fmt.Printf("%c", char)
		time.Sleep(40 * time.Millisecond)
	}
	fmt.Println()
}

func pause(seconds time.Duration) { time.Sleep(seconds * time.Second) }

func clear() { fmt.Print("\033[H\033[2J") }

func main() {

	clear()

	var playerName string

	typewriter("Введи свое имя: ")
	fmt.Scan(&playerName)
	pause(1)

	typewriter(fmt.Sprintf("%v проиграл", playerName))
}
