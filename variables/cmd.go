package variables

import (
	"fmt"
	"os/exec"
)

func BasicCMD(cmd *exec.Cmd) *exec.Cmd {
	cmd.Env = BasicEnv()
	return cmd
}

func FirstOrgCmd(cmd *exec.Cmd) *exec.Cmd {
	cmd.Env = FirstOrgEnv()
	return cmd
}

func SecondOrgCmd(cmd *exec.Cmd) *exec.Cmd {
	cmd.Env = SecondOrgEnv()
	return cmd
}

func handleCmd(b []byte, err error) {
	if err != nil {
		fmt.Println(err, string(b))
		return
	}
	fmt.Println(string(b))
}

func RunBasicCmd(cmd *exec.Cmd) {
	c := BasicCMD(cmd)
	fmt.Println(c.String())
	handleCmd(c.CombinedOutput())
}
func RunFirstOrgCmd(cmd *exec.Cmd) {
	c := FirstOrgCmd(cmd)
	fmt.Println(c.String())
	handleCmd(c.CombinedOutput())
}
func RunSecondOrgCmd(cmd *exec.Cmd) {
	c := SecondOrgCmd(cmd)
	fmt.Println(c.String())
	handleCmd(c.CombinedOutput())
}
