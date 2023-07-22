package responser

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorBuilder struct {
	statusCode int
	message    string
	details    []errorDetail
}

type errorDetail struct {
	Field       string `json:"field,omitempty"`
	Description string `json:"description,omitempty"`
}

func NewErrorBuilder() *ErrorBuilder {
	return &ErrorBuilder{
		message: "An error occurred",
		details: []errorDetail{},
	}
}

func (e *ErrorBuilder) SetMessage(message string) *ErrorBuilder {
	e.message = message
	return e
}

func (e *ErrorBuilder) SetMessagef(format string, vals ...interface{}) *ErrorBuilder {
	e.message = fmt.Sprintf(format, vals...)
	return e
}
func (e *ErrorBuilder) SetStatusCode(s int) *ErrorBuilder {
	e.statusCode = s
	return e
}
func (e *ErrorBuilder) SetDetail(field, description string) *ErrorBuilder {
	e.details = append(e.details, errorDetail{
		Field:       field,
		Description: description,
	})
	return e
}

// ... Add other methods for adding error details ...

func (e *ErrorBuilder) Build() *ErrorResponse {
	return &ErrorResponse{
		StatusCode: e.statusCode,
		Message:    e.message,
		Details:    e.details,
	}
}

type ErrorResponse struct {
	StatusCode int           `json:"status"`
	Message    string        `json:"message"`
	Details    []errorDetail `json:"details,omitempty"`
}

func (r *ErrorResponse) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *ErrorResponse) Respond(ctx echo.Context) error {
	if r.StatusCode == 0 {
		r.StatusCode = http.StatusInternalServerError
	}
	return ctx.JSON(r.StatusCode, r)
}

func (e *ErrorBuilder) Respond(ctx echo.Context) error {
	errorResponse := e.Build()
	return errorResponse.Respond(ctx)
}

func (r *ErrorResponse) WriteTo(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.StatusCode)
	json.NewEncoder(w).Encode(r)
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewResponse(status int, message string, data interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Data:    data,
	}
}
func NewDefaultResponse(data interface{}) *Response {
	return &Response{
		Status:  200,
		Message: "success",
		Data:    data,
	}
}

func (r *Response) ToJSON() ([]byte, error) {
	return json.Marshal(r)
}
