package main

import (
	"fmt"
	"github.com/common-nighthawk/go-figure"
	"os"
	"os/signal"
	"time"
)

const timeFormat = "03:04:05PM"

func main() {
	quit := startDrawing()

	interruptChan := make(chan os.Signal, 1)
	signal.Notify(interruptChan, os.Interrupt)
	for range interruptChan {
		quit <- struct{}{}
		break
	}
}

func startDrawing() chan interface{} {
	ticker := time.NewTicker(time.Second)
	quit := make(chan interface{})
	go drawLoop(ticker, quit)
	return quit
}

func drawLoop(ticker *time.Ticker, quit chan interface{}) {
	for {
		select {
		case currentTime := <-ticker.C:
			draw(currentTime)
		case <-quit:
			ticker.Stop()
			return
		}
	}
}

func draw(currentTime time.Time) {
	timeString := currentTime.Format(timeFormat)
	myFigure := figure.NewFigure(timeString, "", true)
	lines := myFigure.Slicify()

	// TODO: This is bad.  It will go back to the top of the visible terminal.
	// This should clear out an empty section below where you run the program
	// and place the clock there, refreshing only in that area.
	fmt.Printf("\033[0;0H")
	for _, line := range lines {
		fmt.Println(line)
	}
}
