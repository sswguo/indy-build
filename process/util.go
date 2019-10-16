package process

import (
	"fmt"
	"os"
	"os/exec"
)

func printRealCmdOutput(cmd *exec.Cmd) {
	// stdout, _ := cmd.StdoutPipe()
	fmt.Printf("Running cmd: %v in %v\n", cmd, cmd.Dir)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	cmd.Run()
	// for {
	// 	r := bufio.NewReader(stdout)
	// 	line, _, err := r.ReadLine()
	// 	if err == io.EOF {
	// 		break
	// 	}
	// 	if err != nil {
	// 		fmt.Printf(err.Error())
	// 		break
	// 	}
	// 	fmt.Println(string(line))
	// }
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
