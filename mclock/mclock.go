package mclock

import (
	"github.com/aristanetworks/goarista/monotime"
	"time"
)

// AbsTime represents absolute monotonic time.
type AbsTime time.Duration

func Now() AbsTime {
	return AbsTime(monotime.Now())
}

func (t AbsTime) Add(d time.Duration) AbsTime {
	return t + AbsTime(d)
}

// Clock interface makes it possible to replace the monotonic system clock with a simulated clock.
type Clock interface {
	Now() AbsTime
	Sleep(time.Duration)
	After(time.Duration) <-chan time.Time
}

// System implements Clock using the system clock.
type System struct{}

// Now implements Clock.
func (System) Now() AbsTime {
	return AbsTime(monotime.Now())
}

// Sleep implements Clock.
func (System) Sleep(d time.Duration) {
	time.Sleep(d)
}

// After implements Clock.
func (System) After(d time.Duration) <-chan time.Time {
	return time.After(d)
}
