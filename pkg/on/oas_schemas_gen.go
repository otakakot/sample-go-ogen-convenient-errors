// Code generated by ogen, DO NOT EDIT.

package on

import (
	"fmt"
)

func (s *ErrorStatusCode) Error() string {
	return fmt.Sprintf("code %d: %+v", s.StatusCode, s.Response)
}

// Ref: #/components/schemas/Error
type Error struct {
	Message string `json:"message"`
}

// GetMessage returns the value of Message.
func (s *Error) GetMessage() string {
	return s.Message
}

// SetMessage sets the value of Message.
func (s *Error) SetMessage(val string) {
	s.Message = val
}

// ErrorStatusCode wraps Error with StatusCode.
type ErrorStatusCode struct {
	StatusCode int
	Response   Error
}

// GetStatusCode returns the value of StatusCode.
func (s *ErrorStatusCode) GetStatusCode() int {
	return s.StatusCode
}

// GetResponse returns the value of Response.
func (s *ErrorStatusCode) GetResponse() Error {
	return s.Response
}

// SetStatusCode sets the value of StatusCode.
func (s *ErrorStatusCode) SetStatusCode(val int) {
	s.StatusCode = val
}

// SetResponse sets the value of Response.
func (s *ErrorStatusCode) SetResponse(val Error) {
	s.Response = val
}

type GetSampleBadRequest Error

func (*GetSampleBadRequest) getSampleRes() {}

type GetSampleForbidden Error

func (*GetSampleForbidden) getSampleRes() {}

type GetSampleNotFound Error

func (*GetSampleNotFound) getSampleRes() {}

// GetSampleOK is response for GetSample operation.
type GetSampleOK struct{}

func (*GetSampleOK) getSampleRes() {}

type GetSampleUnauthorized Error

func (*GetSampleUnauthorized) getSampleRes() {}

type PostSampleBadRequest Error

func (*PostSampleBadRequest) postSampleRes() {}

type PostSampleForbidden Error

func (*PostSampleForbidden) postSampleRes() {}

type PostSampleNotFound Error

func (*PostSampleNotFound) postSampleRes() {}

// PostSampleOK is response for PostSample operation.
type PostSampleOK struct{}

func (*PostSampleOK) postSampleRes() {}

type PostSampleReq struct {
	// Status.
	Status int `json:"status"`
}

// GetStatus returns the value of Status.
func (s *PostSampleReq) GetStatus() int {
	return s.Status
}

// SetStatus sets the value of Status.
func (s *PostSampleReq) SetStatus(val int) {
	s.Status = val
}

type PostSampleUnauthorized Error

func (*PostSampleUnauthorized) postSampleRes() {}
