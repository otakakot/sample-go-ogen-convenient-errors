package main

import (
	"context"
	"net/http"

	"github.com/otakakot/sample-go-ogen-convenient-errors/pkg/off"
)

func main() {}

var _ off.Handler = (*Handler)(nil)

type Handler struct{}

// GetSample implements off.Handler.
func (h *Handler) GetSample(ctx context.Context, params off.GetSampleParams) (off.GetSampleRes, error) {
	switch params.Status {
	case http.StatusOK:
		return &off.GetSampleOK{}, nil
	case http.StatusBadRequest:
		return &off.GetSampleBadRequest{
			Message: "",
		}, nil
	case http.StatusUnauthorized:
		return &off.GetSampleUnauthorized{
			Message: "",
		}, nil
	case http.StatusForbidden:
		return &off.GetSampleForbidden{
			Message: "",
		}, nil
	case http.StatusNotFound:
		return &off.GetSampleNotFound{
			Message: "",
		}, nil
	default:
		return &off.GetSampleInternalServerError{
			Message: "",
		}, nil
	}
}

// PostSampl implements off.Handler.
func (h *Handler) PostSampl(ctx context.Context, req *off.PostSamplReq) (off.PostSamplRes, error) {
	switch req.Status {
	case http.StatusOK:
		return &off.PostSamplOK{}, nil
	case http.StatusBadRequest:
		return &off.PostSamplBadRequest{
			Message: "",
		}, nil
	case http.StatusUnauthorized:
		return &off.PostSamplUnauthorized{
			Message: "",
		}, nil
	case http.StatusForbidden:
		return &off.PostSamplForbidden{
			Message: "",
		}, nil
	case http.StatusNotFound:
		return &off.PostSamplNotFound{
			Message: "",
		}, nil
	default:
		return &off.PostSamplInternalServerError{
			Message: "",
		}, nil
	}
}
