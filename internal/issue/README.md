# Issues Library 
The Issues Library converts DynoResults into issues that can then be utilized by collaborative software API's. 

Issues are decoupled and expandable with a structure that is versatile to be used by any collaborative coding software. 

- Title contains the title of the issue created in the format `DYNO Fuzz: BUG at Endpoint ENDPOINT using METHOD Method`. 
- Details contains the specific details for the bug found by the dynamic analysis tool. 
- Visualizer is a link to the frontend page to be used by managers to have an overview of all the issues currently suffered within the API. 
- Body contains a DynoResult which is explained in [Results](../result/README.md). 
- Assignee is a user specified by the client for the given issue. 
- Labels specifies the bug and severity level of the bug being either high, medium or low. 
- State specifies if the current issue is closed or open. 
- Milestone is specified by the user and indicates the milestone the issue should be resolved in. 

## Issue Structure 

```go
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
```

## Severity Levels 
- High: Given to issues where there is potential for accessing potential secrets, data leaks, unauthorized user access
- Medium: Given to issues with unexpected errors or success codes and access to deleted resources
- Low: Given to issues where there is a change to the json bodies of a request
