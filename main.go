package main

import (
    "os"
	"os/exec"
    "fmt"
    "bufio"
    "strings"
    "regexp"
)

func record() {
    runos("defaults", []string{"write", "com.spotify.client", "NSTraceEvents", "YES"})

    output := runos("/Applications/Spotify.app/Contents/MacOS/Spotify", []string{})
    scanner := bufio.NewScanner(strings.NewReader(output))
    for scanner.Scan() {
        r, _ := regexp.Compile(`\d{4}-\d{2}-\d{2} (\d{2}:\d{2}:\d{2}.\d{3}) (.+)\[\d+:\d+\] Received event: (.+) at: (.+)`)
        submatch := r.FindStringSubmatch(scanner.Text())
        if len(submatch) == 5 {
            //time := submatch[1]
            //client_name := submatch[2]
            action := submatch[3]
            //details := submatch[4]
            if action == "Gesture" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "MouseMoved" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "MouseEntered" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "SysDefined" {
                scanner.Scan()
            } else if action == "MouseExited" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "LMouseDown" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "LMouseUp" {
                scanner.Scan()
                fmt.Println(scanner.Text())
                scanner.Scan()
                fmt.Println(scanner.Text())
            } else if action == "LMouseDragged" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "RMouseDragged" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "RMouseDown" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "RMouseUp" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "ScrollWheel" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "Kitdefined" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "KeyUp" {
                scanner.Scan()
            } else if action == "KeyDown" {
                scanner.Scan()
                fmt.Println(scanner.Text())
            } else if action == "Resume" {
                scanner.Scan()
                fmt.Println(scanner.Text())
            } else if action == "FlagsChanged" {
                scanner.Scan()
                scanner.Scan()
            } else if action == "AppDefined" {
                scanner.Scan()
            } else {
                fmt.Println("Unknown action: " + action)
            }
        }
    }

    runos("defaults", []string{"write", "com.spotify.client", "NSTraceEvents", "NO"})
}

func runos(cmd string, args []string) string {
    var (
		output []byte
		err    error
    )
    
	if output, err = exec.Command(cmd, args...).CombinedOutput(); err != nil {
		fmt.Fprintln(os.Stderr, "There was an error running command: ", err)
		os.Exit(1)
	}
    
	return string(output)
}

func main() {
	args := os.Args
	
	fmt.Println("Starting...\n")

    if len(args) < 2 {
        fmt.Println("Need 2 args")
        return
    }
    if args[1] == "-r" || args[1] == "--record" {
        record()
        return
    }
}