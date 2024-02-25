package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// также пользуемся встроенной утилитой bash, как и в случае с grep
func main() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("$")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		cmd := exec.Command("bash", "-c", input)

		cmd.Stdin, cmd.Stdout, cmd.Stderr = os.Stdin, os.Stdout, os.Stderr
		cmd.Run()
	}
}
