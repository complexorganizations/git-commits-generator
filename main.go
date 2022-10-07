package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

var (
	removeThisFile = ".delete"
	commitCount    int
	forever        bool
	err            error
	cmd            *exec.Cmd
)

func init() {
	// Make sure all the required applications are installed in the system.
	commandExists("git")
	// If the user passed on how many commits to generate use that, otherwise use the default value.
	if len(os.Args) > 1 {
		tempCommitCount := flag.Int("commit", 0, "The total number of git commits that should be generated.")
		tempForever := flag.Bool("forever", false, "Run the generator forever.")
		flag.Parse()
		commitCount = *tempCommitCount
		forever = *tempForever
	} else {
		log.Fatal("Error: No flags provided. Please use -help for more information.")
	}
	// Make sure the user provded how many commits to generate.
	if commitCount == 0 {
		log.Fatal("Error: Please increase the default value of -commit from 0 to a positive integer.")
	}
}

func main() {
	// Start generating all the commits
	generateCommits()
}

func generateCommits() {
	for loop := 0; loop <= commitCount; loop++ {
		// Generate a random string, and write it to a file.
		err = os.WriteFile(removeThisFile, []byte(randomString(256)), 0644)
		// log all the errors to the console.
		if err != nil {
			log.Println(err)
		}
		// Run the git command to commit the file.
		cmd = exec.Command("git", "add", removeThisFile)
		err = cmd.Run()
		if err != nil {
			log.Println(err)
		}
		// Add a random commit message
		cmd = exec.Command("git", "commit", "-m", randomString(25))
		err = cmd.Run()
		if err != nil {
			log.Println(err)
		}
		// Show the number of commits generated
		log.Println("Commit:", loop, "/", commitCount)
	}
	// Remove the file to cleanup.
	err = os.Remove(removeThisFile)
	if err != nil {
		log.Println(err)
	}
	// Run the git command to commit the file.
	cmd = exec.Command("git", "add", removeThisFile)
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
	// Add a random commit message
	cmd = exec.Command("git", "commit", "-m", randomString(25))
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
	// Once everything is done, push the repo.
	cmd = exec.Command("git", "push")
	err = cmd.Run()
	if err != nil {
		log.Println(err)
	}
	// If the user wants to run the generator forever, run it again.
	if forever {
		generateCommits()
	}
}

// Generate a random string
func randomString(bytesSize int) string {
	rand.Seed(time.Now().UTC().UnixNano())
	randomBytes := make([]byte, bytesSize/2)
	rand.Read(randomBytes)
	randomString := fmt.Sprintf("%X", randomBytes)
	return randomString
}

// To see if a command is available
func commandExists(appname string) {
	_, err := exec.LookPath(appname)
	if err != nil {
		log.Fatalln(err)
	}
}
