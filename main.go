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
	removeThisFile = ".delete"
	commitCount    int
	err            error
	cmd            *exec.Cmd
)

func init() {
	// Make sure it's all installed on your computer
	commandExists("git")
	// Take user flags
	tempCommitCount := flag.Int("commit", 1, "The ammount of commits")
	flag.Parse()
	commitCount = *tempCommitCount
}

func main() {
	generateCommits()
}

func generateCommits() {
	// Begin a cycle
	for loop := 0; loop <= commitCount; loop++ {
		// Create the document.
		err = os.WriteFile(removeThisFile, []byte(randomString(256)), 0644)
		handleErrors(err)
		// Make the required updates to the git repository
		cmd = exec.Command("git", "add", removeThisFile)
		err = cmd.Run()
		handleErrors(err)
		// Add a git commit message to the mix
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		err = cmd.Run()
		handleErrors(err)
		log.Println("Commit:", loop, "/", commitCount)
	}
}

// Generate a random string
func randomString(bytesSize int) string {
	randomBytes := make([]byte, bytesSize)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

// To see if a command is available
func commandExists(appname string) {
	appExists, err := exec.LookPath(appname)
	if err != nil {
		handleErrors(err)
	}
	_ = appExists
}

// Handle all the errors
func handleErrors(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
