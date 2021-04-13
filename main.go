package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

var (
	removeThisFile = "removeThisFile"
	commitCount    = 10000
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
	for loop := 0; loop <= commitCount; loop++ {
		ioutil.WriteFile(removeThisFile, []byte(randomString(256)), 0644)
		cmd := exec.Command("git", "add", removeThisFile)
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		cmd.Run()
		log.Println("Commit:", loop, "/", commitCount)
	}
	cmd := exec.Command("git", "pull")
	cmd.Run()
	cmd = exec.Command("git", "push")
	cmd.Run()
	main()
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
	_ = appName
	return true
}
