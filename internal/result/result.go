package result

type DynoResult struct {
	Title             *string                `json:"title,omitempty"`
	Endpoint          *string                `json:"endpoint,omitempty"`
	HTTPMethod        *string                `json:"httpMethod,omitempty"`
	Method            *string                `json:"method,omitempty"`
	MethodInformation *DynoMethodInformation `json:"methodInformation,omitempty"`
	TimeDelay         *string                `json:"timeDelay,omitempty"`
	AsyncTime         *string                `json:"asyncTime,omitempty"`
	PreviousResponse  *string                `json:"previousResponse,omitempty"`
	ErrorType         *string                `json:"errorType,omitempty"`
}

type DynoMethodInformation struct {
	AcceptedResponse *string `json:"acceptedResponse,omitempty"`
	Host             *string `json:"host,omitempty"`
	ContentType      *string `json:"contentType,omitempty"`
	Request          *string `json:"request,omitempty"`
}
