package parse

import (
	"dyno/internal/result"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO Change the process of reading file according to SQS and S3 Buckets
const bugFile = 6

// Parses the fuzzing files from the bug_buckets folder and creates raw fuzzing results.
// File is filepath to the bug_buckets.txt file that stores all the bugs that has occured.
// Location is the current location of the file on the system.
func ParseRestlerFuzzResults(location string, file string) []result.DynoResult {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	dynoResults := []result.DynoResult{}
	// Creating raw results from fuzz
	for scanner.Scan() {
		line := scanner.Text()
		if line[:1] == "-" {
			scanner.Scan()
			line = scanner.Text()
			bugFileNames := strings.Fields(line)
			if len(bugFileNames) > 5 {
				bugFileName := bugFileNames[bugFile]
				fuzzError := strings.Split(bugFileName, "_")
				results := CreateResults(location, bugFileName, fuzzError)
				for i := 0; i < len(results); i++ {
					dynoResults = append(dynoResults, results[i])
				}
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return dynoResults
}

// Reads the a bug_bucket file that is specified by the category of the bug found by restler
// Creates the body of the issue in regards to the bug found by the fuzzer with details on the bug 
// and how to fix itInternalServerErrors creates a github Issue Request for 
// the categorized bug by restler providing a description on what the bug is and how to possibly fix the bug
// Inputs location is the location of the bug_buckets folder, 
// bugFileName is the name of the file that has the logs of the bug
// fuzzError is the categoried error created
// Returns a list of all the results from the fuzz error that has occured within a DynoResult struct
func CreateResults(location string, bugFileName string, fuzzError []string) ([]result.DynoResult) {
	dynoResults := []result.DynoResult{}
	dynoResult := &result.DynoResult{}
	
	f, err := os.Open(fmt.Sprintf(location+"%s", bugFileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()
	
	title := fmt.Sprintf("%s Invalid %s Response", fuzzError[0], fuzzError[1])
	dynoResult.Title = &title
	dynoResult.ErrorType = &fuzzError[0]

	scanner := bufio.NewScanner(f)

	// Creating DynoResult and adding the result to a list of all the results within the bug 
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[:1] == "-" {
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
func createMethod(requestSplit []string, dynoMethodInformation *result.DynoMethodInformation) (*result.DynoMethodInformation) {
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
func createResult(requestSplit []string, scanner *bufio.Scanner, dynoResult *result.DynoResult) (*result.DynoResult){
	method := strings.Trim(requestSplit[0], "\\r")
	dynoResult.Method = &method
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