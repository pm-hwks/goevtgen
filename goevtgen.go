package main

import (
	"fmt"
	"github.com/urfave/cli"
	"golang.org/x/sys/windows/svc/eventlog"
	"log"
	"os"
	"strconv"
	"time"
)

// Write to windows event log directly from go
func writeEvtLog(eps int, duration int, eventSource string, eventID int, eventType string, message string) {

	const name = "samplelog"
	const supports = eventlog.Error | eventlog.Warning | eventlog.Info

	// Register Event
	err := eventlog.InstallAsEventCreate(name, supports)
	if err != nil {
		fmt.Printf("Install failed: %s", err)
	}
	defer func() {
		err = eventlog.Remove(name)
		if err != nil {
			fmt.Printf("Remove failed: %s", err)
		}
	}()

	//Open event log for writing
	l, err := eventlog.Open(name)
	if err != nil {
		fmt.Printf("Open failed: %s", err)
	}
	defer l.Close()

	// Set up ticker that goes off every second
	ticker := time.NewTicker(1000 * time.Millisecond)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				starttime := time.Now()

				fmt.Println("Tick at", t)
				for i := 0; i < eps; i++ {
					// Write to event log
					l.Info(1, "info from go2")
				}

				elapsed := time.Since(starttime)
				fmt.Println("Sent " + strconv.Itoa(eps) + " messages to evt log in " + elapsed.String())
			}
		}
	}()

	time.Sleep(time.Duration(duration) * 1000 * time.Millisecond)
	ticker.Stop()
	//done <- true
	fmt.Println("Ticker stopped")

}

func main() {

	var eps int
	var duration int
	var eventSource string
	var eventID int
	var eventType string
	var message string

	app := cli.NewApp()
	app.Name = "goevtgen"
	app.Usage = "Generate windows eventlogs"
	app.Version = "0.1"
	app.Compiled = time.Now()

	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "eps",
			Value:       100,
			Usage:       "no of events per second to be generated (tested with 20K eps)",
			Destination: &eps,
		},
		cli.IntFlag{
			Name:        "duration",
			Value:       3,
			Usage:       "No of seconds to run the generation program",
			Destination: &duration,
		},
		cli.StringFlag{
			Name:        "EventSource",
			Usage:       "Event source name",
			Value:       "samplesource",
			Destination: &eventSource,
		},
		cli.IntFlag{
			Name:        "EventID",
			Usage:       "Event ID",
			Value:       123,
			Destination: &eventID,
		},
		cli.StringFlag{
			Name:        "EventType",
			Usage:       "Event Type (Info, Warning, Error)",
			Value:       "Info",
			Destination: &eventType,
		},
		cli.StringFlag{
			Name:        "Message",
			Usage:       "Event payload - message",
			Value:       "This is a sample log",
			Destination: &message,
		},
	}

	app.Action = func(c *cli.Context) error {

		fmt.Println("Initiating goevtgen with the following parameters !")

		fmt.Println("EPS : " + strconv.Itoa(eps))
		fmt.Println("Duration : " + strconv.Itoa(duration))
		fmt.Println("Event Source : " + eventSource)
		fmt.Println("Event ID : " + strconv.Itoa(eventID))
		fmt.Println("Event Type : " + eventType)
		fmt.Println("Message : " + message)
		fmt.Println("\n")

		// Write windows event log
		writeEvtLog(eps, duration, eventSource, eventID, eventType, message)

		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
