package cmd

import (
	"bufio"
	"fmt"
	"io"
	"os/exec"
	"strings"
)

func ExecMvn(goals []string, pomfile string, settingsFile string) {
	args := make([]string, 0)
	for _, goal := range goals {
		args = append(args, goal)
	}

	if len(strings.Trim(pomfile, "")) > 0 {
		args = append(args, "-f")
		args = append(args, pomfile)
	}
	if len(strings.Trim(settingsFile, "")) > 0 {
		args = append(args, "-s")
		args = append(args, settingsFile)
	}
	fmt.Printf("Start executing: %s\n\n", getWholeCmd(args))
	printRealCmdOutput(exec.Command("mvn", args...))
}

func getWholeCmd(args []string) string {
	result := "mvn"
	for _, arg := range args {
		result = result + " " + arg
	}
	return result
}

func printRealCmdOutput(cmd *exec.Cmd) {
	stdout, _ := cmd.StdoutPipe()
	cmd.Start()
	for {
		r := bufio.NewReader(stdout)
		line, _, err := r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Printf(err.Error())
			break
		}
		fmt.Println(string(line))
	}
}
