// Code generated with openapi-go DO NOT EDIT.
package v1

import (
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	oapi_rt "github.com/mworzala/openapi-go/pkg/oapi-rt"
	"github.com/go-chi/chi/v5"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

type PublicServer interface {
	GetTestPlainResp(ctx context.Context, testing int) error
	GetMapWorld(ctx context.Context, id string, abc *GetMapWorldAbc, accept string, req *MapManualTriggerWebhook) (*GetMapWorldResponse, *MapManualTriggerWebhook, error)
}

type PublicServerWrapper struct {
	log         *zap.SugaredLogger
	middlewares []oapi_rt.Middleware
	handler     PublicServer
}

type PublicServerWrapperParams struct {
	fx.In
	Log     *zap.SugaredLogger
	Handler PublicServer

	Middleware []oapi_rt.Middleware `group:"public_middleware"`
}

func NewPublicServerWrapper(p PublicServerWrapperParams) (*PublicServerWrapper, error) {
	sw := &PublicServerWrapper{
		log:         p.Log.With("handler", "public (wrapper)"),
		handler:     p.Handler,
		middlewares: p.Middleware,
	}

	return sw, nil
}

func (sw *PublicServerWrapper) Apply(r chi.Router) {
	r.Route("/v1/public", func(r chi.Router) {
		r.Get("/test/plain_resp", sw.GetTestPlainResp)
		r.Get("/maps/{id}/world", sw.GetMapWorld)
	})
}

func (sw *PublicServerWrapper) GetTestPlainResp(w http.ResponseWriter, r *http.Request) {
	var err error
	_ = err // Sometimes we don't use it but need that not to be an error

	// Read Parameters

	testingStr := r.URL.Query().Get("testing")
	var testing int
	if testingStr != "" {
		testing, err = strconv.Atoi(testingStr)
		if err != nil {
			oapi_rt.WriteGenericError(w, err)
			return
		}
	}

	var handler http.Handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := oapi_rt.NewContext(r.Context(), w, r)

		err := sw.handler.GetTestPlainResp(ctx, testing)
		if err != nil {
			oapi_rt.WriteGenericError(w, err)
			return
		}

		w.WriteHeader(302)
	})
	for _, middleware := range sw.middlewares {
		handler = middleware.Run(handler)
	}
	handler.ServeHTTP(w, r)
}

func (sw *PublicServerWrapper) GetMapWorld(w http.ResponseWriter, r *http.Request) {
	var err error
	_ = err // Sometimes we don't use it but need that not to be an error

	// Read Parameters

	var abc GetMapWorldAbc
	if err := oapi_rt.ReadExplodedQuery(r, &abc); err != nil {
		oapi_rt.WriteGenericError(w, err)
		return
	}

	accept := r.Header.Get("accept")

	id := chi.URLParam(r, "id")

	// Read Body
	var body MapManualTriggerWebhook

	if err = json.NewDecoder(r.Body).Decode(&body); err != nil {
		oapi_rt.WriteGenericError(w, err)
		return
	}

	var handler http.Handler
	handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := oapi_rt.NewContext(r.Context(), w, r)

		code200, code201, err := sw.handler.GetMapWorld(ctx, id, &abc, accept, &body)
		if err != nil {
			oapi_rt.WriteGenericError(w, err)
			return
		}

		if code200 != nil {
			switch {
			case code200.Polar != nil:
				w.Header().Set("content-type", "application/vnd.hollowcube.polar")
				w.WriteHeader(200)
				_, _ = w.Write(code200.Polar)
				return
			case code200.Anvil != nil:
				w.Header().Set("content-type", "application/vnd.hollowcube.anvil")
				w.WriteHeader(200)
				_, _ = w.Write(code200.Anvil)
				return
			}
		}
		if code201 != nil {
			w.Header().Set("content-type", "application/json")
			w.WriteHeader(201)
			if err = json.NewEncoder(w).Encode(code201); err != nil {
				sw.log.Errorw("failed to encode response", "err", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
			return
		}

		w.WriteHeader(204)
	})
	for _, middleware := range sw.middlewares {
		handler = middleware.Run(handler)
	}
	handler.ServeHTTP(w, r)
}
