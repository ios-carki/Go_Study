package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {
	// zsh 쉘 실행 명령어
	command := "zsh"

	// zsh 쉘에 실행할 명령어들 (예: ls, 파일 실행)
	commands := []string{
		"cd /Users/hessegg/Desktop/Temp",
		"open ./Jjang.PNG",
	}

	cmd := exec.Command(command, "-c", strings.Join(commands, " && "))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin

	err := cmd.Run()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}

}
