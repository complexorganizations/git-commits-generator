package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"os/exec"
)

func main() {
	for {
		ioutil.WriteFile("FileName", []byte(randomString(1024)), 0644)
		cmd := exec.Command("git", "add", "FileName")
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		cmd.Run()
		cmd = exec.Command("git", "push")
		cmd.Run()
	}
}

func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}
