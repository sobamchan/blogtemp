package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
    "flag"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
    forYesterdayPtr := flag.Bool("yesterday", false, "Is it for yesterday?")
    flag.Parse()

	// get title
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter title: ")
	title, _ := reader.ReadString('\n')
	title = strings.Replace(title, "\n", "", -1)
	nonSpaceTitle := strings.Replace(title, " ", "-", -1)

    now := time.Now()
	// generate time string
    if *forYesterdayPtr {
	    now = now.AddDate(0, 0, -1)
    } else {
	    now = time.Now()
    }

	nowString := fmt.Sprintf("%d-%02d-%02d", now.Year(), now.Month(), now.Day())

	// generate filename
	filename := fmt.Sprintf("%s-%s.md", nowString, nonSpaceTitle)

	// generate file header
	header := fmt.Sprintf(`---
title: %s
updated: %d-%02d-%02d %02d:%02d
---`, title, now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute())

	// output to file
	f, err := os.Create(filename)
	check(err)

	defer f.Close()

	f.Sync()

	f.WriteString(header)

}
