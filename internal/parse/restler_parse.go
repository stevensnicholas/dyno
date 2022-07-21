package parse

import (
	"bufio"
	"context"
	"fmt"
	"github.com/google/go-github/v45/github"
	"golambda/internal/platform"
	"os"
	"strings"
	"golambda/internal/issue"
)

// TODO Change the process of reading file according to SQS and S3 Buckets
const bugFile = 6

// Parses the fuzzing files from the bug_buckets folder and creates github issues
// Inputs:
//				token is the user token
//				repoName is the user's repo
//				owner is the owner of the repo
//				file is filepath to the bug_buckets.txt file that stores all the bugs that has occured
func ParseFuzz(token string, repoName string, owner string, file string) {
	ctx := context.Background()
	client := platform.CreateClient(ctx, &token)
	f, err := os.Open(file)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	location := "cmd/internal/tests/bug_buckets/"
	scanner := bufio.NewScanner(f)
	dynoIssue := &issue.DynoIssue{}
	dynoIssueBody := &issue.DynoIssueBody{}
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
				title := fmt.Sprintf("# %s Invalid %s Response\n", fuzzError[0], fuzzError[1])
				details := AddDYNODetails(fuzzError[0])
				dynoIssueBody.Title = &title
				dynoIssueBody.Details = &details
				dynoIssueBody = ReadBugFile(location, bugFileName, dynoIssueBody)
				dynoIssue.Body = dynoIssueBody
				client.Issues.Create(ctx, owner, repoName, &github.IssueRequest{
					Title:     dynoIssue.Title,
					Body:      FormatBody(dynoIssueBody),
					Labels:    dynoIssue.Labels,
					Assignee:  dynoIssue.Assignee,
					State:     dynoIssue.State,
					Milestone: dynoIssue.Milestone,
				})
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

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
func ReadBugFile(location string, bugFileName string, dynoIssueBody *issue.DynoIssueBody) (*issue.DynoIssueBody) {
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
			dynoMethodInformation := &issue.DynoMethodInformation{}
			dynoMethodInformation = CreateMethod(requestSplit, dynoMethodInformation)
			dynoIssueBody.MethodInformation = dynoMethodInformation
			CreateBody(requestSplit, scanner, dynoIssueBody)
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	return dynoIssueBody
}

// CreateIssue sorts the bugs found by the fuzzer by there categories and creates a new github issueRequest
// Inputs:
//				fuzzError is the type of bug that has been found by the fuzzer
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*github.IssueRequest with all the relevant information regarding the certain bug
func CreateIssue(fuzzError string, dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	switch fuzzError {
	case "InternalServerErrors":
		dynoIssue = InternalServerErrors(dynoIssue)
	case "UseAfterFreeChecker":
		dynoIssue = UseAfterFreeChecker(dynoIssue)
	case "NameSpaceRuleChecker":
		dynoIssue = NameSpaceRuleChecker(dynoIssue)
	case "ResourceHierarchyChecker":
		dynoIssue = ResourceHierarchyChecker(dynoIssue)
	case "LeakageRuleChecker":
		dynoIssue = LeakageRuleChecker(dynoIssue)
	case "InvalidDynamicObjectChecker":
		dynoIssue = InvalidDynamicObjectChecker(dynoIssue)
	default:
		dynoIssue = PayloadBodyChecker(dynoIssue)
	}
	return dynoIssue
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

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*github.IssueRequest with all the relevant information regarding the certain bug
func InternalServerErrors(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: InternalServerErrors at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*issue.DynoIssue with all the relevant information regarding the certain bug
func ResourceHierarchyChecker(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: ResourceHierarchyChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*issue.DynoIssue with all the relevant information regarding the certain bug
func NameSpaceRuleChecker(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: NameSpaceRuleChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*issue.DynoIssue with all the relevant information regarding the certain bug
func UseAfterFreeChecker(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: UseAfterFreeChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*issue.DynoIssue with all the relevant information regarding the certain bug
func LeakageRuleChecker(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: LeakageRuleChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*issue.DynoIssue with all the relevant information regarding the certain bug
func InvalidDynamicObjectChecker(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

// InternalServerErrors creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*issue.DynoIssue with all the relevant information regarding the certain bug
func PayloadBodyChecker(dynoIssue *issue.DynoIssue) *issue.DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: PayloadBodyChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	return dynoIssue
}

func CreateMethod(requestSplit []string, dynoMethodInformation *issue.DynoMethodInformation) (*issue.DynoMethodInformation) {
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

func CreateBody(requestSplit []string, scanner *bufio.Scanner, dynoIssueBody *issue.DynoIssueBody) (*issue.DynoIssueBody){
	method := strings.Trim(requestSplit[0], "\\r")
	dynoIssueBody.Method = &method
	endpoint := strings.Split(requestSplit[0], " ")[2]
	dynoIssueBody.Endpoint = &endpoint
	scanner.Scan()
	timeDelay := scanner.Text()
	dynoIssueBody.TimeDelay = &timeDelay
	scanner.Scan()
	asyncTime := scanner.Text()
	dynoIssueBody.AsyncTime = &asyncTime
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
	dynoIssueBody.PreviousResponse = &previousResponse
	return dynoIssueBody
}

func FormatBody(dynoIssueBody *issue.DynoIssueBody) *string {
	body := *dynoIssueBody.Title + *dynoIssueBody.Details 
	if *dynoIssueBody.MethodInformation.ContentType != ""{
		body = body + "\n" + *dynoIssueBody.Method + "\n" + "\n" + "- " + *dynoIssueBody.MethodInformation.AcceptedResponse + "\n" + "- " + *dynoIssueBody.MethodInformation.Host + "\n" + "- " + *dynoIssueBody.MethodInformation.ContentType
	} else {
		body = body + "\n" + *dynoIssueBody.Method + "\n" + "\n" + "- " + *dynoIssueBody.MethodInformation.AcceptedResponse + "\n" + "- " + *dynoIssueBody.MethodInformation.Host
	}
	if *dynoIssueBody.MethodInformation.Request != "" {
		body = body + "\n" + "- " + *dynoIssueBody.MethodInformation.Request 
	}
	body = body + "\n" + "\n" + *dynoIssueBody.TimeDelay + "\n" + *dynoIssueBody.AsyncTime + "\n" + "\n" + *dynoIssueBody.PreviousResponse
	body = body + "\n"
	return &body
}