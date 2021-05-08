package main

import (
	"crypto/rand"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
)

var (
	removeThisFile string
	commitCount    int
)

func init() {
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
	tempCommitCount := flag.Int("commit", 10000, "The ammount of commits")
	tempRemoveThisFile := flag.String("file", "removeThisFile", "The ammount of commits")
	flag.Parse()
	commitCount = *tempCommitCount
	removeThisFile = *tempRemoveThisFile
}

func main() {
	generateCommits()
}

func generateCommits() {
	for loop := 0; loop <= commitCount; loop++ {
		os.WriteFile(removeThisFile, []byte(randomString(256)), 0644)
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
	generateCommits()
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
