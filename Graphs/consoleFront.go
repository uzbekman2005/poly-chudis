package graphs

import (
	"fmt"
	"strconv"
	"strings"
)

func MainMenu() {
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("+-              Assalomu alaykum!                 -+")
	fmt.Println("+-      Polychudes o'yiniga xo'sh kelibsiz!       -+")
	fmt.Println("+--------------------------------------------------+")
	fmt.Print("\n")
	fmt.Println(" ---> 1 - Bosing > Shartlar bilan tanishish uchun<  ")
	fmt.Println(" ---> 2 - Bosing 	   > O`yini boshlash uchun<      ")
	// 1 Start
	// 2 End
}

func ShowOptions() {
	fmt.Println("+--------------------------------------------------+")
	fmt.Println("+-      O`yin 2 xil yo`l bilan o`ylaniladi        -+")
	fmt.Println("+-                                                -+")
	fmt.Println("+-    -> 1. So`zlarni harflar orqali topasiz <-   -+")
	fmt.Println("+- -> 2. So`zlarni bitta so`z kiritib topasiz <-  -+")
	fmt.Println("+--------------------------------------------------+")
}

func ShowRes(res int) {
	// if 1 win else if 0 lose
	if res == 1 {
		fmt.Println("+--------------------------------------------------+")
		fmt.Println("+-				TABRIKLAYMIZ!!!!!!				   -+")
		fmt.Println("+-      		  SIZ YUTDINGIZ                    -+")
		fmt.Println("+--------------------------------------------------+")
		fmt.Print("\n")
		fmt.Println(" ---> 1. Bosing >> O`yin qaytadan o`ynash uchun     ")
		fmt.Println(" ---> 2. Bosing >> O`yindan chiqish uchun           ")
	} else {

		fmt.Println("****************************************************")
		fmt.Println("#              AFSUSKI SIZ TOPA OLMADINGIZ         #")
		fmt.Println("****************************************************")
		fmt.Print("\n")
		fmt.Println(" ---> 1. Bosing >> O`yin qaytadan o`ynash uchun     ")
		fmt.Println(" ---> 2. Bosing >> O`yindan chiqish uchun           ")
	}
}

func InputNum(hint string) int {
	var temp string

	for {
		fmt.Print(hint, ": ")
		fmt.Scan(&temp)
		res, err := strconv.Atoi(temp)
		if err == nil {
			return res
		}
		fmt.Println("Faqat butun son kriting!")
	}

}

func ShowQuestion(question string) {
	listQuestion := strings.Split(question, " ")
	// total width is 52 bytes
	curLength := 0
	l := len(listQuestion)
	var temp []string
	fmt.Println("+--------------------------------------------------+")
	for i := 0; i < l; i++ {
		if curLength+len(listQuestion[i])+len(temp) > 45 {
			fmt.Printf("+- %-46s -+\n", strings.Join(temp, " "))
			temp = []string{listQuestion[i]}
			curLength = len(listQuestion[i])
		} else {
			curLength += len(listQuestion[i])
			temp = append(temp, listQuestion[i])
		}
	}
	if len(temp) > 0 {
		fmt.Printf("+- %-46s -+\n", strings.Join(temp, " "))
	}
	fmt.Println("+--------------------------------------------------+")
}

func InputString(hint string) string {
	fmt.Print(hint + ": ")
	var temp string
	fmt.Scan(&temp)
	return temp
}
