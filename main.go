package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readFile(fileName string) string {
	content, err := os.ReadFile(fileName)
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return string(content)
}

func validateCommand(command string, arg ...string) bool {
	if command == "man" {
		if len(arg) != 1 {
			fmt.Println("Invalid Command")
			return false
		}
	} else if command == "wc" {
		if len(arg) != 2 {
			fmt.Println("Invalid Command")
			return false
		}
	}
	return true
}

func commandHelper(command string) {
	switch command {
	case "wc":
		fmt.Println(readFile("man_wc.txt"))
	default:
		fmt.Println("Command not supported")

	}
}

func executeCommand(cmd string, arg ...string) {
	if validateCommand(cmd, arg...) {
		switch cmd {
		case "wc":
			content := readFile(strings.Trim(arg[1], "\n"))
			if content != "" {
				switch arg[0] {
				case "-c":
					count := len(content)
					fmt.Println(count, " ", arg[1])
				case "-w":
					count := len(strings.Fields(content))
					fmt.Println(count, " ", arg[1])
				case "-l":
					count := len(strings.Split(content, "\n"))
					fmt.Println(count, " ", arg[1])
				default:
					fmt.Println("Invalid subcommand")
				}
			}
		case "man":
			commandHelper(arg[0])
		default:
			fmt.Println("Command not supported")
		}
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Welcome to Unix terminal")

	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")

	for text != "quit" {
		inputArray := strings.Split(text, " ")
		executeCommand(inputArray[0], inputArray[1:]...)
		text, _ = reader.ReadString('\n')
		text = strings.Trim(text, "\n")
	}
}
