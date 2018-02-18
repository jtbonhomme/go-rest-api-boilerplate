package handler

// A NotFoundError is a swagger response to represent error when resource is not found
//
// swagger:response notFoundError
type notFoundError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// A GenericError is a swagger response to represent any kind of error
//
// swagger:response genericError
type genericError struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}

// An OkError is a swagger response to represent the standard response when no problem occurred
//
// swagger:response okResponse
type okResponse struct {
	// in: body
	Body struct {
		Code    int32  `json:"code"`
		Message string `json:"message"`
	} `json:"body"`
}
