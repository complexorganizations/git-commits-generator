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
	err            error
)

func init() {
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
	tempCommitCount := flag.Int("commit", 10000, "The ammount of commits")
	tempRemoveThisFile := flag.String("file", "remove-this-file", "The name of the file to remove")
	flag.Parse()
	commitCount = *tempCommitCount
	removeThisFile = *tempRemoveThisFile
}

func main() {
	generateCommits()
}

func generateCommits() {
	for loop := 0; loop <= commitCount; loop++ {
		if fileExists(removeThisFile) {
			err = os.Remove(removeThisFile)
			handleErrors(err)
		} else if !fileExists(removeThisFile) {
			err = os.WriteFile(removeThisFile, []byte(randomString(256)), 0644)
			handleErrors(err)
		}
		cmd := exec.Command("git", "add", removeThisFile)
		err = cmd.Run()
		handleErrors(err)
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		err = cmd.Run()
		handleErrors(err)
		log.Println("Commit:", loop, "/", commitCount)
	}
	cmd := exec.Command("git", "pull")
	err = cmd.Run()
	handleErrors(err)
	cmd = exec.Command("git", "push")
	err = cmd.Run()
	handleErrors(err)
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

func handleErrors(err error) {
	if err != nil {
		log.Println(err)
	}
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
