package ansbyletter

import (
	"fmt"
	graphs "game/Graphs"
	"strings"
)

func AnswerByLetterMain(question, answer string) int {
	outputAnswer := generateOutPutAnswer(len(answer))
	rightAnswer := strings.Split(answer, "")
	numberOfRounds := int(float64(len(rightAnswer)) * 1.5)
	graphs.ShowQuestion(question)
	for numberOfRounds > 0 {
		showAnswer(outputAnswer)
		ans := graphs.InputString("Enter a letter")
		ans = strings.ToLower(ans)
		isRightAnswer(ans, rightAnswer, outputAnswer)
		graphs.SystemClear() //console is cleared here
		graphs.ShowQuestion(question)
		if strings.Join(outputAnswer, "") == answer {
			return 1
		}
		numberOfRounds--
	}
	showAnswer(rightAnswer)
	return 0
}

func isRightAnswer(letter string, rightAnswer []string, outPutAnswer []string) {
	for i := 0; i < len(rightAnswer); i++ {
		if string(rightAnswer[i]) == letter && string(outPutAnswer[i]) == "-" {
			outPutAnswer[i] = rightAnswer[i]
		}
	}
}

func generateOutPutAnswer(l int) []string {
	var res []string
	for i := 0; i < l; i++ {
		res = append(res, "-")
	}
	return res
}

func showAnswer(list []string) {
	var temp []string
	for i := 0; i < len(list); i++ {
		temp = append(temp, "---")
	}
	fmt.Print("+", strings.Join(temp, "+"), "+\n")
	fmt.Println("|", strings.Join(list, " | "), "|")
	fmt.Print("+", strings.Join(temp, "+"), "+\n")

}
