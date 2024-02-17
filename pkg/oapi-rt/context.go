package oapi_rt

import (
	"context"
	"net/http"
)

// I am not in love with this because it introduces a dependency on the generator in the generated code
// (for fetching the accept header), but I don't have a better idea.

type contextKey string

const (
	writerKey  contextKey = "writer"
	requestKey contextKey = "request"
	acceptKey  contextKey = "accept"
)

func NewContext(ctx context.Context, w http.ResponseWriter, r *http.Request) context.Context {
	ctx = context.WithValue(ctx, writerKey, w)
	ctx = context.WithValue(ctx, requestKey, r)
	return context.WithValue(ctx, acceptKey, r.Header.Get("Accept"))
}

// RequestFromContext returns the request from the context.
// Bad bad very bad do not use unless necessary
func RequestFromContext(ctx context.Context) *http.Request {
	return ctx.Value(requestKey).(*http.Request)
}

// WriterFromContext returns the writer from the context.
// Bad bad very bad by zak
func WriterFromContext(ctx context.Context) http.ResponseWriter {
	return ctx.Value(writerKey).(http.ResponseWriter)
}
