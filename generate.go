package main

import (
    "fmt"
    "bufio"
    "os"
    "time"
    "strings"
)

func check(e error) {
    if e != nil {
        panic(e)
    }
}

func main() {
    // get title
    reader := bufio.NewReader(os.Stdin)
    fmt.Print("Enter title: ")
    title, _ := reader.ReadString('\n')
    title = strings.Replace(title, "\n", "", -1)
    nonSpaceTitle := strings.Replace(title, " ", "-", -1)

    // generate time string
    now := time.Now()
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
