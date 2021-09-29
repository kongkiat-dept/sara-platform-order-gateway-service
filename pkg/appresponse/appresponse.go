package appresponse

import (
	"net/http"
)

var (
	Success             = Status{Code: http.StatusOK, Message: []string{"Success"}}
	BadRequest          = Status{Code: http.StatusBadRequest, Message: []string{"Sorry, Not responding because of incorrect syntax"}}
	Unauthorized        = Status{Code: http.StatusUnauthorized, Message: []string{"Sorry, We are not able to process your request. Please try again"}}
	Forbidden           = Status{Code: http.StatusForbidden, Message: []string{"Sorry, Permission denied"}}
	InternalServerError = Status{Code: http.StatusInternalServerError, Message: []string{"Internal Server Error"}}
)

// ResponseBody struct
type ResponseBody struct {
	Status Status      `json:"status,omitempty"`
	Data   interface{} `json:"data,omitempty"`
	Pagination
}

// Status struct
type Status struct {
	Code    int      `json:"code,omitempty"`
	Message []string `json:"message,omitempty"`
}

type Pagination struct {
	CurrentPage *int `json:"current_page,omitempty"`
	PerPage     *int `json:"per_page,omitempty"`
	TotalItem   *int `json:"total_item,omitempty"`
}
