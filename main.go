package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {

	args := os.Args[1:]

	fmt.Println(args)

	file := fmt.Sprintf("./day%s/part%s.go", args[0], args[1])
	arg := fmt.Sprintf("./day%s/%s.txt", args[0], args[2])

	out, _ := exec.Command("go", "run", file, arg).Output()
	// out, _ := exec.Command("echo", "test").Output()
	fmt.Println(string(out))

}
