package main

import (
	"bytes"
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

var messageTemplate = `
Hi all,
We have many canvassing sessions scheduled this week. Please join us to help Win ‘24! 🌹

%s`

var eventTemplate = `🗓️ Date/Time: %s
📍 Meeting Point: %s, %s
🔗 Link: %s

`

func main() {
	fmt.Println("Formatting Labour Events!")

	argsWithoutProg := os.Args[1:]

	if len(argsWithoutProg) < 1 {
		os.Exit(-1)
	}

	csvFilePath := argsWithoutProg[0]
	fmt.Println("Opening...", csvFilePath)

	in, err := os.Open(csvFilePath)
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}

	defer func(in *os.File) {
		err := in.Close()
		if err != nil {
			os.Exit(-2)
		}
	}(in)

	var events []*Event
	var eventStringBuffer bytes.Buffer

	if err := gocsv.UnmarshalFile(in, &events); err != nil {
		fmt.Println("Error parsing csv:", err)
		os.Exit(-2)
	}

	for _, event := range events {
		fmt.Println("Adding", event.EventTitle, event.StartTime)
		eventStringBuffer.WriteString(fmt.Sprintf(eventTemplate, event.StartTime, event.Location, event.Postcode, event.URL))
	}

	fmt.Println(fmt.Sprintf(messageTemplate, eventStringBuffer.String()))
}
