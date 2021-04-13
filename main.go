package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func init() {
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
}

func main() {
	generateCommits()
}

func generateCommits() {
	for {
		ioutil.WriteFile("Delete-This-File", []byte(randomString(128)), 0644)
		cmd := exec.Command("git", "add", "Delete-This-File")
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(10))
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

func commandExists(cmd string) bool {
	appName, err := exec.LookPath(cmd)
	if err != nil {
		return false
	}
	_ = appName // variable declared and not used
	return true
}
