package ansatonce

import graphs "game/Graphs"

func AnsAtOnceMain(question, answer string) int {
	userAns := graphs.InputString("Enter your answer: ")
	if userAns == answer{
		return 1
	}
	return 0
}
