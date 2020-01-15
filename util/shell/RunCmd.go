package shell

import (
	"os/exec"
)

func RunCmd(cmdStr string) string {
	cmd := exec.Command("sh", "-c", cmdStr)
	bytes, err := cmd.Output()

	if err != nil {
		panic(err)
	}

	return string(bytes)
}
