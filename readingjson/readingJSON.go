package main

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"
)

func main() {
	const jsonStream = `
	[
		{"Name": "Ed", "Text": "Knock knock.", "Sl": [1, 2, 3], "Submessage": {"M":"somesub1"}},
		{"Name": "Sam", "Text": "Who's there?", "Sl": [1, 2, 3, 4], "Submessage": {"M":"somesub2"}}
	]
`

	type Sub struct {
		M string
	}

	type Message struct {
		Name, Text string
		Sl         []int
		Submessage Sub
	}

	dec := json.NewDecoder(strings.NewReader(jsonStream))

	// read open bracket
	t, err := dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

	// while the array contains values
	for dec.More() {
		var m Message
		// decode an array value (Message)
		err := dec.Decode(&m)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%v: %v    %v    %v \n", m.Name, m.Text, m.Sl, m.Submessage)
	}

	// read closing bracket
	t, err = dec.Token()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%T: %v\n", t, t)

}
