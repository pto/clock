// clock counts down to or up from a target time.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	target := time.Date(2016, 7, 10, 0, 0, 0, 0, time.UTC)
	motto := "Just Go"
	printTargetTime(target, motto)
	exitOnEnterKey()

	var previous time.Time
	for {
		now := time.Now()
		now = now.Add(time.Duration(-now.Nanosecond())) // truncate to second
		if now != previous {
			previous = now
			remaining := target.Sub(now)
			printTimeRemaining(now, remaining)
		}
		time.Sleep(50 * time.Millisecond)
	}
}

func exitOnEnterKey() {
	go func() {
		buf := make([]byte, 1)
		_, _ = os.Stdin.Read(buf)
		os.Exit(0)
	}()
}

const (
	indent         = "\t"
	highlightStart = "\x1b[1;35m"
	highlightEnd   = "\x1b[0m"
)

func printTargetTime(target time.Time, motto string) {
	fmt.Print(indent, highlightStart, motto, highlightEnd, "\n")
	fmt.Print(indent, target.Format(time.UnixDate), "\n")
}

func printTimeRemaining(now time.Time, remaining time.Duration) {
	var sign string
	if remaining > 0 {
		sign = "-" // countdown is "T minus..."
	} else {
		sign = "+" // count up is "T plus..."
		remaining = -remaining
	}

	var days int
	if remaining >= 24*time.Hour {
		days = int(remaining / (24 * time.Hour))
		remaining = remaining % (24 * time.Hour)
	}

	fmt.Print(indent, now.Format(time.UnixDate), "  ", sign)
	if days > 0 {
		fmt.Print(days, "d")
	}
	fmt.Print(remaining, "          \r")
}
