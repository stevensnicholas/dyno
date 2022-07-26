package issue

import (
	"dyno/internal/result"
	"fmt"
)

type DynoIssue struct {
	Title      *string            `json:"title,omitempty"`
	Details    *string            `json:"details,omitempty"`
	Visualizer *string            `json:"visualizer,omitempty"`
	Body       *result.DynoResult `json:"body,omitempty"`
	Assignee   *string            `json:"assignee,omitempty"`
	Labels     *[]string          `json:"labels,omitempty"`
	State      *string            `json:"state,omitempty"`
	Milestone  *int               `json:"milestone,omitempty"`
}

// CreateIssues collects all the results from the raw fuzzing results and formats
// them using the DynoIssue struct to be presented as an issue on any communication platform.
// Inputted is a slice of dynoResults and returns all the issues as a []DynoIssue.
func CreateIssues(dynoResults []result.DynoResult) []DynoIssue {
	dynoIssues := []DynoIssue{}
	dynoIssue := &DynoIssue{}
	for i := range dynoResults {
		dynoIssue.Body = &dynoResults[i]
		dynoIssue = createIssue(*dynoResults[i].ErrorType, dynoIssue)
		if dynoIssue != nil {
			dynoIssues = append(dynoIssues, *dynoIssue)
		}
	}
	return dynoIssues
}

// createIssue sorts the bugs found by the fuzzer by there categories and creates a new DynoIssue
// Inputted is a fuzzError is the type of bug that has been found by the fuzzer and a
// dynoIssue is a struct of a issue. Returns *DynoIssue with all the relevant
// information regarding the certain bug
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
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func internalServerErrorsIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: InternalServerErrors at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: '500 Internal Server' Errors and any other 5xx errors are detected."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}

// resourceHierarchyCheckerIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func resourceHierarchyCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: ResourceHierarchyChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: Detects that a child resource can be accessed from a non-parent resource."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}

// nameSpaceRuleCheckerIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func nameSpaceRuleCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: NameSpaceRuleChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: Detects that an unauthorized user can access service resources."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}

// useAfterFreeCheckerIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func useAfterFreeCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: UseAfterFreeChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: Detects that a deleted resource can still being accessed after deletion."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}

// leakageRuleCheckerIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func leakageRuleCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: LeakageRuleChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: Detects that a failed resource creation leaks data in subsequent requests."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}

// invalidDynamicObjectCheckerIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func invalidDynamicObjectCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: InvalidDynamicObjectChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: Detects 500 errors or unexpected success status codes when invalid dynamic objects are sent in requests."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}

// payloadBodyCheckerIssue creates a github Issue Request for the categorized bug by restler
// providing a description on what the bug is and how to possibly fix the bug.
// Inputted is a *DynoIssue which is null and is a struct of a issue.
// Returns *DynoIssue with all the relevant information regarding the certain bug
func payloadBodyCheckerIssue(dynoIssue *DynoIssue) *DynoIssue {
	title := fmt.Sprintf("DYNO Fuzz: PayloadBodyChecker at Endpoint %s", *dynoIssue.Body.Endpoint)
	details := "Details: Detects 500 errors when fuzzing the JSON bodies of requests."
	visualizer := "Visualizer: [DYNO](the web url)"
	labels := []string{"bug"}
	dynoIssue.Title = &title
	dynoIssue.Labels = &labels
	dynoIssue.Details = &details
	dynoIssue.Visualizer = &visualizer
	return dynoIssue
}
