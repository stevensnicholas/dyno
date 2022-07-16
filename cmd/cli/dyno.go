package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"io/ioutil"
)

// Create a new type for a list of Strings
type stringList []string

// Implement the flag.Value interface
func (s *stringList) String() string {
	return fmt.Sprintf("%v", *s)
}

func (s *stringList) Set(value string) error {
	*s = strings.Split(value, ",")
	return nil
}

func main() {
	// Subcommands
	sendCommand := flag.NewFlagSet("send", flag.ExitOnError)

	// Count subcommand flag pointers
	// Adding a new choice for --metric of 'substring' and a new --substring flag

	filePath := sendCommand.String("dest", "", "file to send. (Required)")

	// Verify that a subcommand has been provided
	// os.Arg[0] is the main command
	// os.Arg[1] will be the subcommand
	if len(os.Args) < 2 {
		fmt.Println("send subcommand is required")
		os.Exit(1)
	}

	// Switch on the subcommand
	// Parse the flags for appropriate FlagSet
	// FlagSet.Parse() requires a set of arguments to parse as input
	// os.Args[2:] will be all arguments starting after the subcommand at os.Args[1]
	switch os.Args[1] {
	case "send":
		sendCommand.Parse(os.Args[2:])
	default:
		flag.PrintDefaults()
		os.Exit(1)
	}

	// Check which subcommand was Parsed using the FlagSet.Parsed() function. Handle each case accordingly.
	// FlagSet.Parse() will evaluate to false if no flags were parsed (i.e. the user did not provide any flags)

	if sendCommand.Parsed() {
		if *filePath == "" {
			sendCommand.PrintDefaults()
			os.Exit(1)
		}

		data, err := os.Open(*filePath)
    if err != nil {
        log.Fatal(err)
    }

		// will be replaced with actual api-endpoint, 
		url := "https://o8cnchwjji.execute-api.ap-southeast-2.amazonaws.com/v1/post_json"
    fmt.Println("URL:>", url)

    req, err := http.NewRequest("POST", url, data)
    req.Header.Set("X-Custom-Header", "myvalue")
    req.Header.Set("Content-Type", "application/json")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }
    defer resp.Body.Close()

    fmt.Println("response Status:", resp.Status)
    fmt.Println("response Headers:", resp.Header)
    body, _ := ioutil.ReadAll(resp.Body)
    fmt.Println("response Body:", string(body))

	}

}