package main

var (
	// ErrMessageCouldNotParseRequest is thrown when an invalid request object is provided
	ErrMessageCouldNotParseRequest = "request unable to be parsed"
)

// ExampleRequest is an example of a request object that might be passed in as part of an API request
type ExampleRequest struct {
	LoanID     string  `json:"loanID"`                      // sensitivity flag omitted on purpose (sensitive:"false" is default)
	LoanAmount float64 `json:"loanAmount" sensitive:"true"` // sensitivity flag declared on purpose to omit this entry
}
