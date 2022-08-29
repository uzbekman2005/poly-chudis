package main

import (
	"fmt"
	ansatonce "game/ansAtOnce"
	ansbyletter "game/ansbyletter"
)

func main() {
	fmt.Println("Hello world")
	answer := "tashkent"
	question := "What is capital of Uzbekistan?"
	fmt.Println(ansatonce.AnsAtOnceMain(question, answer))
	fmt.Println(ansbyletter.AnswerByLetterMain(question, answer))
}
