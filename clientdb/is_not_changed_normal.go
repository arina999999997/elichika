//go:build !embedded

package clientdb

import (
	"elichika/config"
	"elichika/log"

	"os/exec"
)

func isNotChanged(file string) bool {
	cmd := exec.Command("git", "diff", "--exit-code", "--quiet", file)
	cmd.Dir = config.AssetPath
	err := cmd.Run()
	if err == nil {
		return true // exit code is 0
	}
	exitError, ok := err.(*exec.ExitError)
	if !ok {
		log.Panic(err)
	}
	if exitError.ExitCode() != 1 {
		log.Panic(err)
	}
	return false
}
