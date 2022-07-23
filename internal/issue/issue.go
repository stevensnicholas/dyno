package issue

import (
	"fmt"
	"golambda/internal/result"
)

type DynoIssue struct {
	Title *string `json:"title,omitempty"`
	Details *string  `json:"details,omitempty"`
	Visualizer *string  `json:"visualizer,omitempty"`
	Body *result.DynoResult `json:"body,omitempty"`
	Assignee *string `json:"assignee,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
	State *string `json:"state,omitempty"`
	Milestone *int `json:"milestone,omitempty"`
}

func CreateIssues(dynoResults []result.DynoResult) []DynoIssue{
	dynoIssues := []DynoIssue{}
	dynoIssue := &DynoIssue{}
	for _, dynoResult := range dynoResults {
		dynoIssue.Body = &dynoResult
		dynoIssue = createIssue(*dynoResult.ErrorType, dynoIssue)
		if dynoIssue != nil {
			dynoIssues = append(dynoIssues, *dynoIssue)
		}
	}
	return dynoIssues 
}

// createIssue sorts the bugs found by the fuzzer by there categories and creates a new github issueRequest
// Inputs:
//				fuzzError is the type of bug that has been found by the fuzzer
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*github.IssueRequest with all the relevant information regarding the certain bug
func createIssue(fuzzError string, dynoIssue *DynoIssue) *DynoIssue {
	switch fuzzError {
	case "InternalServerErrors":
		dynoIssue = internalServerErrorsIssue(dynoIssue)
	case "UseAfterFreeChecker":
		dynoIssue = useAfterFreeCheckerIssue(dynoIssue)
	case "NameSpaceRuleChecker":
		dynoIssue = nameSpaceRuleCheckerIssue(dynoIssue)
	case "ResourceHierarchyChecker":
		dynoIssue = resourceHierarchyCheckerIssue(dynoIssue)
	case "LeakageRuleChecker":
		dynoIssue = leakageRuleCheckerIssue(dynoIssue)
	case "InvalidDynamicObjectChecker":
		dynoIssue = invalidDynamicObjectCheckerIssue(dynoIssue)
	default:
		dynoIssue = payloadBodyCheckerIssue(dynoIssue)
	}
	return dynoIssue
}

// internalServerErrorsIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug
// Inputs:
//        body is the body of the github issue
// 				endpoint is the endpoint that has the bug
// 				assignee is if there is a specified github user that should be assigned for checking this certain type of bug
// 				state is the current state of the issue
// 				milestone specifies if the issue should be linked to a certain milestone on the users repo
// Returns:
// 				*github.IssueRequest with all the relevant information regarding the certain bug
func internalServerErrorsIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: InternalServerErrors at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: '500 Internal Server' Errors and any other 5xx errors are detected.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
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
// 				*DynoIssue with all the relevant information regarding the certain bug
func resourceHierarchyCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: ResourceHierarchyChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: Detects that a child resource can be accessed from a non-parent resource.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details 
	dynoIssue.Visualizer = &visualizer
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
// 				*DynoIssue with all the relevant information regarding the certain bug
func nameSpaceRuleCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: NameSpaceRuleChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: Detects that an unauthorized user can access service resources.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details 
	dynoIssue.Visualizer = &visualizer
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
// 				*DynoIssue with all the relevant information regarding the certain bug
func useAfterFreeCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: UseAfterFreeChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: Detects that a deleted resource can still being accessed after deletion.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details 
	dynoIssue.Visualizer = &visualizer
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
// 				*DynoIssue with all the relevant information regarding the certain bug
func leakageRuleCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: LeakageRuleChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: Detects that a failed resource creation leaks data in subsequent requests.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details 
	dynoIssue.Visualizer = &visualizer
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
// 				*DynoIssue with all the relevant information regarding the certain bug
func invalidDynamicObjectCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details 
	dynoIssue.Visualizer = &visualizer
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
// 				*DynoIssue with all the relevant information regarding the certain bug
func payloadBodyCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: PayloadBodyChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "\nDetails: Detects 500 errors when fuzzing the JSON bodies of requests.\n"
	visualizer := "\nVisualizer: [DYNO](the web url)\n"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels 
	dynoIssue.Details = &details 
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}