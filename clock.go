// clock counts down to or up from a target time.
package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	const (
		indent          = "\t"
		highlight_start = "\x1b[1;36m"
		highlight_end   = "\x1b[0m"
	)

	fmt.Print(indent, highlight_start, "Just Go", highlight_end, "\n")
	target := time.Date(2016, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Print(indent, target.Format(time.UnixDate), "\n")

	var (
		previous time.Time
		days     int
		sign     string
	)

	// Exit on Enter key
	go func() {
		buf := make([]byte, 1)
		_, _ = os.Stdin.Read(buf)
		os.Exit(0)
	}()

	for {
		now := time.Now()
		now = now.Add(time.Duration(-now.Nanosecond())) // truncate to second
		if now != previous {
			previous = now
			remaining := target.Sub(now)
			if remaining > 0 {
				sign = "-" // countdown is "T minus..."
			} else {
				sign = "+" // count up is "T plus..."
				remaining = -remaining
			}
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
		time.Sleep(50 * time.Millisecond)
	}
}
