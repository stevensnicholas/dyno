package parse

import (
	"bufio"
	"dyno/internal/result"
	"fmt"
	"strings"
)

// Parses the fuzzing files from the bug_buckets folder and creates raw fuzzing results.
// fileContents is the contents of a file from the s3 bucket object
// fuzzError is the type of bug found
func ParseRestlerFuzzResults(fileContents string, fuzzError []string) []result.DynoResult {
	dynoResults := []result.DynoResult{}
	dynoResult := &result.DynoResult{}
	title := fmt.Sprintf("%s Invalid %s Response", fuzzError[0], fuzzError[1])
	dynoResult.Title = &title
	dynoResult.ErrorType = &fuzzError[0]
	file := strings.NewReader(fileContents)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 1 && line[0:2] == "->" {
			requestSplit := strings.Split(line, "\\n")
			dynoMethodInformation := &result.DynoMethodInformation{}
			dynoMethodInformation = createMethod(requestSplit, dynoMethodInformation)
			dynoResult.MethodInformation = dynoMethodInformation
			dynoResult = createResult(requestSplit, scanner, dynoResult)
			if dynoResult != nil {
				dynoResults = append(dynoResults, *dynoResult)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return dynoResults
}

// createMethod creates the method information within a DynoResult that is supplied with the
// accepted response, host, content and the request that was sent
// Inputed is the requestSplit string that contains information on the fuzz bug
// and dynoMethodInformation which holds all the information about the method of fuzzing.
// Returning DynoMethodInformation struct
func createMethod(requestSplit []string, dynoMethodInformation *result.DynoMethodInformation) *result.DynoMethodInformation {
	acceptedResponse := strings.Trim(requestSplit[1], "\\r")
	host := strings.Trim(requestSplit[2], "\\r")
	contentType := strings.Trim(requestSplit[3], "\\r")
	dynoMethodInformation.AcceptedResponse = &acceptedResponse
	dynoMethodInformation.Host = &host
	dynoMethodInformation.ContentType = &contentType
	request := ""
	for i := 4; i < len(requestSplit); i++ {
		request = request + requestSplit[i]
	}
	if request != "" {
		request = "Request: " + strings.Trim(request, "\\r")
	}
	dynoMethodInformation.Request = &request
	return dynoMethodInformation
}

// createResult creates the result structure of a DynoResult
// containing all the necessary information from the fuzzing
// Inputed is the requestSplit string that contains information on the fuzz bug,
// scanner of the bug file allowing to read the next line of text and the current dynoResult
// Returning a dynoResult struct
func createResult(requestSplit []string, scanner *bufio.Scanner, dynoResult *result.DynoResult) *result.DynoResult {
	method := strings.Trim(requestSplit[0], "\\r")
	httpMethod := strings.Split(method, " ")[1]
	dynoResult.Method = &method
	dynoResult.HTTPMethod = &httpMethod
	endpoint := strings.Split(requestSplit[0], " ")[2]
	dynoResult.Endpoint = &endpoint
	scanner.Scan()
	timeDelay := scanner.Text()
	dynoResult.TimeDelay = &timeDelay
	scanner.Scan()
	asyncTime := scanner.Text()
	dynoResult.AsyncTime = &asyncTime
	scanner.Scan()
	previousResponseText := scanner.Text()
	previousResponseSplit := strings.Split(previousResponseText, "\\n")
	prevresp := ""
	for i := 5; i < len(previousResponseSplit); i++ {
		prevresp = prevresp + previousResponseSplit[i]
	}
	if prevresp != "" {
		prevresp = " response:" + strings.Trim(prevresp, "\\r")
	}
	previousResponse := strings.Trim(previousResponseSplit[0], "\\r") + prevresp
	dynoResult.PreviousResponse = &previousResponse
	return dynoResult
}

// GetFuzzError is given the filename as a string and 
// returns the type of fuzzing error found by the fuzzer
func GetFuzzError(filename string) []string{
	fuzzError := strings.Split(filename, "_")
	if len(fuzzError) <= 1 {
		panic("FileName: not recognised as a RESTler File")
	}
	return strings.Split(filename, "_")
}