package parse

import (
	"golambda/internal/result"
	"bufio"
	"fmt"
	"os"
	"strings"
)

// TODO Change the process of reading file according to SQS and S3 Buckets
const bugFile = 6

// Parses the fuzzing files from the bug_buckets folder and creates github issues
// Inputs:
//				file is filepath to the bug_buckets.txt file that stores all the bugs that has occured
func ParseRestlerFuzzResults(file string) []result.DynoResult {
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	location := "internal/tests/bug_buckets/"
	scanner := bufio.NewScanner(f)
	dynoResult := &result.DynoResult{}
	dynoResults := []result.DynoResult{}
	// Sending an issue for each error found through fuzz
	for scanner.Scan() {
		line := scanner.Text()
		if line[:1] == "-" {
			scanner.Scan()
			line = scanner.Text()
			bugFileNames := strings.Fields(line)
			if len(bugFileNames) > 5 {
				bugFileName := bugFileNames[bugFile]
				fuzzError := strings.Split(bugFileNames[bugFile], "_")
				dynoResult.ErrorType = &fuzzError[0]
				title := fmt.Sprintf("%s Invalid %s Response\n", fuzzError[0], fuzzError[1])
				details := AddDYNODetails(fuzzError[0])
				dynoResult.Title = &title
				dynoResult.Details = &details
				dynoResult = CreateResult(location, bugFileName, dynoResult)
				if dynoResult != nil {
					dynoResults = append(dynoResults, *dynoResult)
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
// Creates the body of the issue in regards to the bug found by the fuzzer with details on the bug and how to fix itInternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
//
// Inputs:
//				bugFileName is the name of the file that has the logs of the bug
//        body is the start of the body for the github issue
// Returns:
// 				body is the body of the issue
// 				endpoint is the endpoint that has the bug
func CreateResult(location string, bugFileName string, dynoResult *result.DynoResult) (*result.DynoResult) {
	f, err := os.Open(fmt.Sprintf(location+"%s", bugFileName))
	if err != nil {
		panic(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)

	// Creating body for IssueRequest in Github
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) > 0 && line[:1] == "-" {
			requestSplit := strings.Split(line, "\\n")
			dynoMethodInformation := &result.DynoMethodInformation{}
			dynoMethodInformation = CreateMethod(requestSplit, dynoMethodInformation)
			dynoResult.MethodInformation = dynoMethodInformation
			dynoResult = CreateBody(requestSplit, scanner, dynoResult)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
	return dynoResult
}

// AddDYNODetails adds the details and visualizer url to the body of the issue request that
// corresponds to the FuzzError that is received
// Inputs:
//				fuzzError is the type of bug that has been found by the fuzzer
// Returns:
// 				details a string with the specified details
// 				If not details are created leaves an empty string
func AddDYNODetails(fuzzError string) string {
	details := ""
	switch fuzzError {
	case "InternalServerErrors":
		details = "\nDetails: '500 Internal Server' Errors and any other 5xx errors are detected.\n\nVisualizer: [DYNO](the web url)\n"
	case "UseAfterFreeChecker":
		details = "\nDetails: Detects that a deleted resource can still being accessed after deletion.\n\nVisualizer: [DYNO](the web url)\n"
	case "NameSpaceRuleChecker":
		details = "\nDetails: Detects that an unauthorized user can access service resources.\n\nVisualizer: [DYNO](the web url)\n"
	case "ResourceHierarchyChecker":
		details = "\nDetails: Detects that a child resource can be accessed from a non-parent resource.\n\nVisualizer: [DYNO](the web url)\n"
	case "LeakageRuleChecker":
		details = "\nDetails: Detects that a failed resource creation leaks data in subsequent requests.\n\nVisualizer: [DYNO](the web url)\n"
	case "InvalidDynamicObjectChecker":
		details = "\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n\nVisualizer: [DYNO](the web url)\n"
	case "PayloadBodyChecker":
		details = "\nDetails: Detects 500 errors when fuzzing the JSON bodies of requests.\n\nVisualizer: [DYNO](the web url)\n"
	default:
		details = ""
	}
	return details
}

func CreateMethod(requestSplit []string, dynoMethodInformation *result.DynoMethodInformation) (*result.DynoMethodInformation) {
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

func CreateBody(requestSplit []string, scanner *bufio.Scanner, dynoResult *result.DynoResult) (*result.DynoResult){
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
	prevrequest := ""
	for i := 5; i < len(previousResponseSplit); i++ {
		prevrequest = prevrequest + previousResponseSplit[i]
	}
	if prevrequest != "" {
		prevrequest = " request:" + strings.Trim(prevrequest, "\\r")
	}
	previousResponse := strings.Trim(previousResponseSplit[0], "\\r") + prevrequest
	dynoResult.PreviousResponse = &previousResponse
	return dynoResult
}