package main

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

var updating = false

func getVersion() (string, error) {
	output, err := exec.Command("git", "rev-parse", "HEAD").Output()
	return string(output), err
}

func rebuildElichika() (string, error) {
	output, err := exec.Command("git", "submodule", "deinit", "-f", "assets").Output()
	if err != nil {
		return "", errors.New(fmt.Sprint("error running git submodule deinit.\noutput: ", string(output), "\nerror: ", err))
	}
	output, err = exec.Command("git", "submodule", "update", "--init", "--recursive", "--checkout", "assets").Output()
	if err != nil {
		return "", errors.New(fmt.Sprint("error running git submodule update.\noutput: ", string(output), "\nerror: ", err))
	}
	output, err = exec.Command("go", "build").Output()
	if err != nil {
		cmd := exec.Command("go", "build")
		cmd.Env = append(os.Environ(), "CGO_ENABLED=0")
		output, err = cmd.Output()
	}
	if err != nil {
		return "", errors.New(fmt.Sprint("error building elichika.\noutput: ", string(output), "\nerror: ", err))
	}

	cmd := exec.Command("./elichika", "rebuild_assets")
	output, err = cmd.Output()
	if err != nil {
		return "", errors.New(fmt.Sprint("error building elichika assets.\noutput: ", string(output), "\nerror: ", err))
	}

	version, _ := getVersion()
	return "Updated elichika to version: " + version, nil
}

func updateElichika() (string, error) {
	if updating {
		return "Elichika is already being updated", nil
	}
	updating = true
	output, err := exec.Command("git", "pull").Output()
	if err != nil {
		return "", errors.New(fmt.Sprint("error running git pull.\noutput: ", string(output), "\nerror: ", err))
	}
	if strings.Contains(string(output), "Already up to date.") {
		return "", errors.New(fmt.Sprint("elichika already up to date.\ngit pull output: ", string(output)))
	}
	outputStr, err := rebuildElichika()
	return outputStr, err
}
