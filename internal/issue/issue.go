package issue

type DynoIssue struct {
	Title *string `json:"title,omitempty"`
	Body *DynoIssueBody `json:"body,omitempty"`
	Assignee *string `json:"assignee,omitempty"`
	Labels *[]string `json:"labels,omitempty"`
}

type DynoIssueBody struct {
	Title *string `json:"title,omitempty"`
	Details *string  `json:"details,omitempty"`
	Visualizer *string  `json:"visualizer,omitempty"`
	Endpoint *string `json:"endpoint,omitempty"`
	Method *string  `json:"method,omitempty"`
	MethodInformation *DynoMethodInformation `json:"methodInformation,omitempty"`
	TimeDelay *string  `json:"timeDelay,omitempty"`
	AsyncTime *string `json:"asyncTime,omitempty"`
	PreviousResponse *string  `json:"previousResponse,omitempty"`
}

type DynoMethodInformation struct {
	AcceptedResponse *string `json:"acceptedResponse,omitempty"`
	Host *string `json:"host,omitempty"`
	ContentType *string `json:"contentType,omitempty"`
	Request *string `json:"request,omitempty"`
}