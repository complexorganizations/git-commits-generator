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
	numberOfCommits = 1000000
	commitFileName  = "file"
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
	for loop := 1; loop <= numberOfCommits; loop++ {
		ioutil.WriteFile(commitFileName, []byte(randomString(512)), 0644)
		cmd := exec.Command("git", "add", commitFileName)
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		cmd.Run()
	}
	cmd := exec.Command("git", "push")
	cmd.Run()
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
