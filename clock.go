// Clock counts down to or up from a target time.
package main

import (
	"fmt"
	"os"
	"time"
)

// Update target and motto as desired.
var (
	target = time.Date(2020, 5, 1, 0, 0, 0, 0, time.Local)
	motto  = "Simply Go"
)

func main() {
	printTargetTime(target, motto)
	exitOnEnterKey()

	var previous time.Time
	for {
		now := time.Now().Truncate(time.Second)
		if now != previous {
			previous = now
			countdown := now.Sub(target) // Negative times are before the target
			printCountdown(now.In(target.Location()), countdown)
		}
		time.Sleep(100 * time.Millisecond)
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
	highlightStart = "\x1b[1;35m"
	highlightEnd   = "\x1b[0m"
	indent         = "\t"
)

func printTargetTime(target time.Time, motto string) {
	fmt.Print(indent, highlightStart, motto, highlightEnd, "\n")
	fmt.Print(indent, target.Format(time.UnixDate), "\n")
}

func printCountdown(now time.Time, countdown time.Duration) {
	var sign string
	if countdown >= 0 {
		sign = "+"
	} else {
		sign = "-"
		countdown = -countdown
	}

	days := int(countdown / (24 * time.Hour))
	countdown = countdown % (24 * time.Hour)

	fmt.Print(indent, now.Format(time.UnixDate), "  ", sign)
	if days > 0 {
		fmt.Print(days, "d")
	}
	fmt.Print(countdown, "          \r")
	os.Stdout.Sync()
}
