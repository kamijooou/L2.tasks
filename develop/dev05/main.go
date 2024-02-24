package main

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

// пользуемся утилитой самого терминала
func main() {
	input, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	input = strings.TrimSpace(input)
	args := strings.Split(input, " ")

	cmd := exec.Command("grep", args...)
	cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr

	cmd.Run()
}
