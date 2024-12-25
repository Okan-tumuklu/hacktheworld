package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strings"
	// "time"  // If you suspect a blocking factor on the target site, remove this line and line 57 from the comment section, and set the duration as desired in line 57.
)

func main() {
	
	
 fmt.Println("+-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+")
 fmt.Println("|d| |i| |r| |e| |c| |t| |o| |r| |y| |s| |c| |a| |n|")
 fmt.Println("İ am a jr ;) ")
 fmt.Println("+-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+ +-+")

	
	if len(os.Args) != 4 || os.Args[2] != "-w" {
		fmt.Println("Usage: go run directoryscan.go https://site/ -w /wordlist-path")
		return
	}

	siteURL := strings.TrimRight(os.Args[1], "/")
	wordlistPath := os.Args[3]

	file, err := os.Open(wordlistPath)
	if err != nil {
		fmt.Printf("Failed to open wordlist: %v\n", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		if word == "" {
			continue
		}
		fullURL := fmt.Sprintf("%s/%s", siteURL, word)
		resp, err := http.Get(fullURL)
		if err != nil {
			fmt.Printf("Request error: %v\n", err)
			continue
		}

		// Here, I took help from artificial intelligence to display the incoming responses properly.
		func() {
			defer resp.Body.Close()
			// Process only 200 OK responses
			if resp.StatusCode == http.StatusOK {
				fmt.Printf("[200] %s\n", fullURL)
			}
		}()
		// time.Sleep(3 * time.Millisecond)
	}

	if err := scanner.Err(); err != nil {
		fmt.Printf("File read error: %v\n", err)
	}
}
