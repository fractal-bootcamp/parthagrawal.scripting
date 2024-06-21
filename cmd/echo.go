package cmd

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
)

func echo() {
	dance := exec.Command("echo", "hello")
	output, err := dance.Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Ran command!")
	fmt.Println(strings.TrimSpace(string(output)))
}
