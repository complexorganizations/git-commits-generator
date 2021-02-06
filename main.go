package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	commandsRequirementsCheck()
	generateCommits()
}

func commandsRequirementsCheck() {
	if !commandExists("git") {
		log.Fatal("Error: Git was not discovered in the system.")
	}
}

func generateCommits() {
	for {
		ioutil.WriteFile("Delete-This-File", []byte(randomString(128)), 0644)
		cmd := exec.Command("git", "add", "Delete-This-File")
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(10))
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
	_, err := exec.LookPath(cmd)
	return err == nil
}
