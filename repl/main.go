package repl

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"time"

	"github.com/knetic/govaluate"
)

const NAME string = "cas"
const EXIT string = "q"

func printRepl() {
	fmt.Printf("%s> ", NAME)
}

func recoverExp(text string) {
	if r := recover(); r != nil {
		printRepl()
		fmt.Println("unknown command ", text)
	}
}

func printInvalidCmd(text string) {
	// We might have a panic here we so need DEFER + RECOVER
	defer recoverExp(text)
	// \n Will be ignored
	t := strings.TrimSuffix(text, "\n")
	if t != "" {
		expression, errExp := govaluate.NewEvaluableExpression(text)
		result, errEval := expression.Evaluate(nil)
		// Before we need to know if is not a Math expr
		printRepl()
		if errExp == nil && errEval == nil {
			fmt.Println(result)
		} else {
			fmt.Println("unknow command " + t)
		}
	}
}

func getLine(r *bufio.Reader) string {
	t, _ := r.ReadString('\n')
	return strings.TrimSpace(t)
}

func shouldContinue(text string) bool {
	if strings.EqualFold(EXIT, text) {
		return false
	}
	return true
}

func help() {
	fmt.Println("Welcome to CAS Repl! ")
	fmt.Println("(C) Matt Hall 2020")
	fmt.Println("Under the MIT License")
	fmt.Println("Avaliable commands: ")
	fmt.Println("	help - display help")
	fmt.Println("	cls - clear terminal screen")
	fmt.Println("	q - exit the repl")
	fmt.Println("	express - enter an expression")
}

func cls() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func now() {
	fmt.Println("go-repl> ", time.Now().Format(time.RFC850))
}

func processLine(line string) error {
	fmt.Printf("You said: '%s'\n", line)
	if 0 == 9 {
		return errors.New("f u")
	}
	return nil
}

func main() {
	// commands := map[string]interface{}{
	// 	"help": help,
	// 	"cls":  cls,
	// 	"time": now,
	// }
	reader := bufio.NewReader(os.Stdin)
	help()
	printRepl()
	text := getLine(reader)
	for ; shouldContinue(text); text = getLine(reader) {
		err := processLine(text)
		if err != nil {
			fmt.Println("FUCK")
			printInvalidCmd(text)
		} else {

		}
		// if value, exists := commands[text]; exists {
		// 	value.(func())()
		// } else {
		// 	printInvalidCmd(text)
		// }
		printRepl()
	}
	fmt.Println("-- session ended --")

}
