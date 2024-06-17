// Code generated by ogen, DO NOT EDIT.

package off

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

type GetSampleBadRequest Error

func (*GetSampleBadRequest) getSampleRes() {}

type GetSampleForbidden Error

func (*GetSampleForbidden) getSampleRes() {}

type GetSampleInternalServerError Error

func (*GetSampleInternalServerError) getSampleRes() {}

type GetSampleNotFound Error

func (*GetSampleNotFound) getSampleRes() {}

// GetSampleOK is response for GetSample operation.
type GetSampleOK struct{}

func (*GetSampleOK) getSampleRes() {}

type GetSampleUnauthorized Error

func (*GetSampleUnauthorized) getSampleRes() {}

type PostSamplBadRequest Error

func (*PostSamplBadRequest) postSamplRes() {}

type PostSamplForbidden Error

func (*PostSamplForbidden) postSamplRes() {}

type PostSamplInternalServerError Error

func (*PostSamplInternalServerError) postSamplRes() {}

type PostSamplNotFound Error

func (*PostSamplNotFound) postSamplRes() {}

// PostSamplOK is response for PostSampl operation.
type PostSamplOK struct{}

func (*PostSamplOK) postSamplRes() {}

type PostSamplReq struct {
	// Status.
	Status int `json:"status"`
}

// GetStatus returns the value of Status.
func (s *PostSamplReq) GetStatus() int {
	return s.Status
}

// SetStatus sets the value of Status.
func (s *PostSamplReq) SetStatus(val int) {
	s.Status = val
}

type PostSamplUnauthorized Error

func (*PostSamplUnauthorized) postSamplRes() {}
