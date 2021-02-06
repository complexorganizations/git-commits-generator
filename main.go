package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

var (
	commitFileName = "Delete-This-File"
)

func main() {
	commandsRequirementsCheck()
	generateCommits()
}

// Check if GIT is installed in the system
func commandsRequirementsCheck() {
	if !commandExists("git") {
		log.Println("Error: Git was not discovered in the system.")
		os.Exit(0)
	}
}

// Generate a commit
func generateCommits() {
	for {
		ioutil.WriteFile(commitFileName, []byte(randomString(256)), 0644)
		cmd := exec.Command("git", "add", commitFileName)
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(10))
		cmd.Run()
	}
}

// Generate a string at random
func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

// Verify if there is a command
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
