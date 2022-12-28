package mclock

import (
	"sync"
	"time"
)

type Simulated struct {
	now       AbsTime
	scheduled []event
	mu        sync.RWMutex
	cond      *sync.Cond
}

type event struct {
	do func()
	at AbsTime
}

func (s *Simulated) init() {
	if s.cond == nil {
		s.cond = sync.NewCond(&s.mu)
	}
}

func (s *Simulated) insert(d time.Duration, do func()) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.init()

	at := s.now + AbsTime(d)
	l, h := 0, len(s.scheduled)
	ll := h
	for l != h {
		m := (l + h) / 2
		if at < s.scheduled[m].at {
			h = m
		} else {
			l = m + 1
		}
	}
	s.scheduled = append(s.scheduled, event{})
	copy(s.scheduled[l+1:], s.scheduled[l:ll])
	s.scheduled[l] = event{do: do, at: at}
	s.cond.Broadcast()
}

func (s *Simulated) Run(d time.Duration) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.init()

	end := s.now + AbsTime(d)
	for len(s.scheduled) > 0 {
		ev := s.scheduled[0]
		if ev.at > end {
			break
		}
		s.now = ev.at
		ev.do()
		s.scheduled = s.scheduled[1:]
	}
	s.now = end
}

// Now implements Clock.
func (s *Simulated) Now() AbsTime {
	s.mu.RLock()
	defer s.mu.RUnlock()

	return s.now
}

// Sleep implements Clock.
func (s *Simulated) Sleep(d time.Duration) {
	<-s.After(d)
}

// After implements Clock.
func (s *Simulated) After(d time.Duration) <-chan time.Time {
	after := make(chan time.Time, 1)
	s.insert(d, func() {
		after <- (time.Time{}).Add(time.Duration(s.now))
	})
	return after
}
