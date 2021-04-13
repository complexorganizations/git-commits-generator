package main

import (
	"crypto/rand"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"strconv"
)

var userInput string
var commitCount int
var removeThisFile string

func init() {
	if !commandExists("git") {
		log.Fatal("Error: The application git was not found in the system.")
	}
	commitCount = 10000
	if len(os.Args) > 1 {
		userInput = os.Args[1]
		commitCount, _ := strconv.ParseInt(userInput, 0, 0)
		_ = commitCount
	}
}

func main() {
	generateCommits()
}

func generateCommits() {
	for loop := 0; loop <= commitCount; loop++ {
		removeThisFile = "removeThisFile"
		ioutil.WriteFile(removeThisFile, []byte(randomString(256)), 0644)
		cmd := exec.Command("git", "add", removeThisFile)
		cmd.Run()
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		cmd.Run()
		log.Println("Commit:", loop, "/", commitCount)
	}
	cmd := exec.Command("git", "push")
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
	_ = appName // variable declared and not used
	return true
}
