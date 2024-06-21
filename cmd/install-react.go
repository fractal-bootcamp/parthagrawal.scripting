package cmd

import (
	"fmt"
	"log"
	"os/exec"
)

func install_react() {

	install := exec.Command("bun", "create", "vite")
	output, err := install.Output()
	if err != nil {
		log.Println(err)
	}
	fmt.Println("Ran install!")
	fmt.Println(string(output))

}
