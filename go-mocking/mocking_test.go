package main

import (
	"testing"
	"time"
	"bytes"
	"github.com/aszenz/basic-golang/go-test-helpers"
)

// ---------- //

// SpyTime 
type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration 
}

func TestDefaultSleeper(t *testing.T){
	sleepTime := 5 * time.Second
	spyTime := &SpyTime{}
	sleeper := DefaultSleeper{sleepTime, spyTime.Sleep} 
	sleeper.Sleep()
	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v but slept for %v", sleepTime, spyTime.durationSlept)
	}
}

// ---------- //

type SpySleeper struct {
	Calls []string
}

const sleep = "sleep"
const write = "write"

// Sleep for SpySleeper is the mocked sleep
func (s *SpySleeper) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

func (s *SpySleeper) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return 
}

func TestCountPrinter(t *testing.T){
	t.Run("prints 3 to Go!", func(t *testing.T){
		buffer := &bytes.Buffer{}
		CountPrinter(buffer, &SpySleeper{})
		got := buffer.String()
		want := "3\n2\n1\nGo!"
		testhelpers.AssertStringsEqual(t, got, want)
	})
	t.Run("sleep before every print", func(t *testing.T){
		spySleeper := &SpySleeper{}
		CountPrinter(spySleeper, spySleeper)
		got := spySleeper.Calls
		want := []string{"write", "sleep", "write", "sleep", "write", "sleep", "write"} 
		testhelpers.AssertStringArraysEqual(t, got, want)
	})
}

// ---------- //
