package main

import (
	"context"
	"errors"
	"net/http"

	"github.com/otakakot/sample-go-ogen-convenient-errors/pkg/on"
)

func main() {}

var _ on.Handler = (*Handler)(nil)

type Handler struct{}

// GetSample implements on.Handler.
func (h *Handler) GetSample(ctx context.Context, params on.GetSampleParams) (on.GetSampleRes, error) {
	if err := h.SwitchError(ctx, params.Status); err != nil {
		return nil, err
	}

	return &on.GetSampleOK{}, nil
}

// PostSample implements on.Handler.
func (h *Handler) PostSample(ctx context.Context, req *on.PostSampleReq) (on.PostSampleRes, error) {
	if err := h.SwitchError(ctx, req.Status); err != nil {
		return nil, err
	}

	return &on.PostSampleOK{}, nil
}

func (h *Handler) SwitchError(
	ctx context.Context,
	status int,
) error {
	switch status {
	case http.StatusOK:
		return nil
	case http.StatusBadRequest:
		return ErrBadRequest
	case http.StatusUnauthorized:
		return ErrUnauthorized
	case http.StatusForbidden:
		return ErrForbidden
	case http.StatusNotFound:
		return ErrNotFound
	default:
		return ErrInternal
	}
}

var (
	ErrBadRequest   = errors.New("bad request")
	ErrUnauthorized = errors.New("unauthorized")
	ErrForbidden    = errors.New("forbidden")
	ErrNotFound     = errors.New("not found")
	ErrInternal     = errors.New("internal server error")
)

// NewError implements on.Handler.
func (h *Handler) NewError(ctx context.Context, err error) *on.ErrorStatusCode {
	switch {
	case errors.Is(err, ErrBadRequest):
		return &on.ErrorStatusCode{
			StatusCode: http.StatusBadRequest,
			Response: on.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, ErrUnauthorized):
		return &on.ErrorStatusCode{
			StatusCode: http.StatusUnauthorized,
			Response: on.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, ErrForbidden):
		return &on.ErrorStatusCode{
			StatusCode: http.StatusForbidden,
			Response: on.Error{
				Message: err.Error(),
			},
		}
	case errors.Is(err, ErrNotFound):
		return &on.ErrorStatusCode{
			StatusCode: http.StatusNotFound,
			Response: on.Error{
				Message: err.Error(),
			},
		}
	}

	return &on.ErrorStatusCode{
		StatusCode: http.StatusInternalServerError,
		Response: on.Error{
			Message: err.Error(),
		},
	}
}
