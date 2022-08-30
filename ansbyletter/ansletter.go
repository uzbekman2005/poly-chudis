package ansbyletter

import (
	"fmt"
	"strings"
)

func AnswerByLetterMain(question, answer string) int {
	showAnswer([]string{"a", "-", "b", "-", "-", "t"})
	return 1
}

func showAnswer(list []string) {
	var temp []string
	for i := 0; i < len(list); i++{
		temp = append(temp, "---")
	}
	fmt.Print("+", strings.Join(temp,"+"), "+\n")
	fmt.Println("|", strings.Join(list, " | "), "|")
	fmt.Print("+", strings.Join(temp,"+"), "+\n")


}
