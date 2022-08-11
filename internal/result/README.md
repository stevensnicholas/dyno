# Result Library 

The results library contains various extensible structures that are used by the [parse library](../parse) to format raw dynamic analysis results. 

### DynoResult Structure 
- Title: A customized title that is currently in the format `VULNERABILITY Invalid HTTP Response` where it displays the vulnerability found and the http response that was executed
- Endpoint: The endpoint that was fuzzed
- HTTPMethod: The http method that was executed
- Method: The HTTPMethod and the endpoint
- MethodInformation a DynoMethodInformation structure
- TimeDelay: The time between sending the request and retrieving a response 
- AsyncTime: The time between sending the request and retrieving a response 
- PreviousResponse: The previous response that was executed on the endpoint useful for data leaks
- ErrorType: The vulnerability that was found 
```go
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
```

### DynoMethodInformation Structure 
- Accepted Response: The type of data that was accepted 
- Host: Where the fuzzing was hosted
- ContentType: The type of content that was accepted 
- Request: The request that was sent 
```go
type DynoMethodInformation struct {
	AcceptedResponse *string `json:"acceptedResponse,omitempty"`
	Host             *string `json:"host,omitempty"`
	ContentType      *string `json:"contentType,omitempty"`
	Request          *string `json:"request,omitempty"`
}
```