package process

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
)

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

func destroyRepo(repoLocation string) {
	os.RemoveAll(repoLocation)
	fmt.Printf("Repo removed: %s\n", repoLocation)
}

func getWholeCmd(cmd string, args []string) string {
	result := cmd
	for _, arg := range args {
		result = result + " " + arg
	}
	return result
}
