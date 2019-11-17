package main

import (
	"io"
	"time"
	"fmt"
	"os"
)

// Sleeper interface is used to pause execution 
type Sleeper interface {
	Sleep()
}

// DefaultSleeper is the actual implementation of Sleeper interface
type DefaultSleeper struct {
	duration time.Duration
	sleep func(time.Duration)
}

// Sleep Pauses the execution time by 1 sec
func (d *DefaultSleeper) Sleep() {
	d.sleep(d.duration)
}

// CountPrinter prints countdown from 3 to 1 then Go
func CountPrinter(writer io.Writer, s Sleeper){
	countDownStart := 3
	countDownEnd := 0
	clincher := "Go!"
	for i:=countDownStart; i > countDownEnd; i-- {
		fmt.Fprintln(writer, i)
		s.Sleep()
	}
	fmt.Fprint(writer, clincher)
}

func main(){
	defaultSleeper := &DefaultSleeper{1 * time.Second, time.Sleep}
	CountPrinter(os.Stdout, defaultSleeper)
}