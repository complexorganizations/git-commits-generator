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
	fileName      string
	commitMessege string
	byteSize      string
)

func init() {
	fileName = randomString(10)
	commitMessege = randomString(10)
	byteSize = randomString(128)
}

func main() {
	commandsRequirementsCheck()
	generateCommits()
}

func commandsRequirementsCheck() {
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
}

func generateCommits() {
	for {
		ioutil.WriteFile(fileName, []byte(byteSize), 0644)
		cmd := exec.Command("git", "add", fileName)
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", commitMessege)
		cmd.Run()
		os.Remove(fileName)
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
