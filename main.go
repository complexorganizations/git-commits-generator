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
)

// Check for GIT
func commandsRequirementsCheck() {
	if !commandExists("git") {
		log.Println("Error: Git was not discovered in the system.")
		os.Exit(0)
	}
}

// Run the app
func main() {
	commandsRequirementsCheck()
	for loop := 1; loop <= numberOfCommits; loop++ {
		ioutil.WriteFile("FileName", []byte(randomString(1024)), 0644)
		cmd := exec.Command("git", "add", "FileName")
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		cmd.Run()
	}
	cmd := exec.Command("git", "push")
	cmd.Run()
}

// Generate Random String
func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

// Check if a command exists
func commandExists(cmd string) bool {
	_, err := exec.LookPath(cmd)
	return err == nil
}
