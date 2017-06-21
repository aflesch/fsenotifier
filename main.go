// +build darwin

// fsenotifier - application that wrap fsevents package, gets as argument the watch point path
//               and outout to stdout a json events
// Version 0.1

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"bitbucket.org/minutelab/mlab/sync/mnotify/fsenotifier/fsedata"
	"github.com/fsnotify/fsevents"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("%s", answerError(fsedata.ErrorNoWatchPoint))
		os.Exit(1)
	}

	watchpath := os.Args[1]
	info, err := os.Stat(watchpath)
	if err != nil {
		fmt.Printf("%s", answerError(fsedata.ErrorWrongWatchPoint))
		os.Exit(1)
	}
	// Must be a directory not a file
	if !info.IsDir() {
		fmt.Printf("%s", answerError(fsedata.ErrorNoDirWatchPoint))
		os.Exit(1)
	}

	if err := start(watchpath); err != nil {
		fmt.Printf("%s", answerError(*err))
		os.Exit(1)
	}
}

func start(watchPointPath string) *fsedata.Errors {

	dev, err := fsevents.DeviceForPath(watchPointPath)
	if err != nil {
		log.Fatalf("Failed to retrieve device for path: %v", err)
	}

	es := &fsevents.EventStream{
		Paths:   []string{watchPointPath},
		Latency: 100 * time.Millisecond,
		Device:  dev,
		Flags:   fsevents.FileEvents | fsevents.WatchRoot}

	es.Start()
	readLoop(es)
	return nil
}

func answerError(err fsedata.Errors) string {
	answer := fsedata.FSEvents{
		Fserror: err,
	}

	jsonAnswer, _ := json.Marshal(answer)

	return string(jsonAnswer[:])
}

func readLoop(es *fsevents.EventStream) {
	ec := es.Events
	for msg := range ec {
		var events fsedata.FSEvents
		for _, event := range msg {
			newEvent := fsedata.FSEvent{Path: event.Path, Flags: fsedata.EventFlags(event.Flags)}
			events.Events = append(events.Events, newEvent)
		}
		anserEvent(events)
	}
}

func anserEvent(events fsedata.FSEvents) {
	jsonAnswer, _ := json.Marshal(events)
	fmt.Printf("%s", jsonAnswer)
}
