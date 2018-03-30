package main

import (
	"os"
	"fmt"
	"strings"
	//"os/exec"
	"os/exec"
)

func main() {
	path := os.Args[1]
	parts := strings.Split(path, "/")
	host := parts[0]
	projectPath := strings.Join(parts[1:], "/")

	cloneUrl := fmt.Sprintf("git@%v:%v", host, projectPath)

	basePath, _ := os.LookupEnv("GOPATH")

	cloneDir := fmt.Sprintf("%v/src/%v", basePath, path)

	output, err := exec.Command("git", "clone", cloneUrl, cloneDir).CombinedOutput()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}

	fmt.Println(string(output))

}
