package main

import (
	graphs "game/Graphs"
	ansatonce "game/ansAtOnce"
	ansbyletter "game/ansbyletter"
	"game/config"
)

func main() {
	var choice, res, doStop int
	for {
		graphs.MainMenu()
		choice = graphs.InputNum("> ")
		if choice == 1 {
			graphs.SystemClear()
			graphs.ShowOptions()
		} else if choice == 2 {
			graphs.SystemClear()
			question, answer := config.GetRandomQuestion()
			graphs.ShowQuestion(question)
			graphs.ShowOptions()
			choice = graphs.InputNum(">")
			graphs.SystemClear()
			switch choice {
			case 1:
				res = ansbyletter.AnswerByLetterMain(question, answer)
			case 2:
				res = ansatonce.AnsAtOnceMain(question, answer)
			default:
				continue
			}
			doStop = graphs.ShowRes(res)
			if doStop == 1 {
				continue
			} else {
				break
			}

		} else {
			break
		}

	}
}
