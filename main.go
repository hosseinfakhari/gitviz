package main

import (
	"flag"
)

func main() {
	var folder string
	var email string
	flag.StringVar(&folder, "add", "", "add new folder to scan for Git repository")
	flag.StringVar(&email, "email", "fakhari.hossein@gmail.com", "the email to scan")
	flag.Parse()

	if folder != "" {
		scan(folder)
		return
	}
	stats(email)
}
