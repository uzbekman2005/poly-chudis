package main

import (
	"fmt"
	graphs "game/Graphs"
	// ansatonce "game/ansAtOnce"
	ansbyletter "game/ansbyletter"
	"game/config"
)

func main() {
	fmt.Println("Hello world")
	question, answer := config.GetRandomQuestion()
	// fmt.Println(ansatonce.AnsAtOnceMain(question, answer))
	fmt.Println(ansbyletter.AnswerByLetterMain(question, answer))
	graphs.ShowQuestion(question)
}
