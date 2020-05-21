package middleware

import (
	"context"

	"github.com/sirupsen/logrus"
	"github.com/spiffe/spire/pkg/server/api/rpccontext"
)

const (
	// serverAPIPrefix is the prefix on the SPIRE server APIs full method names.
	// We strip it out to keep things concise.
	serverAPIPrefix = "/spire.api.server."
)

// WithLogger returns logging middleware that provides a per-rpc logger with
// some initial fields set. If unset, it also provides name metadata on
// to the handler context.
func WithLogger(log logrus.FieldLogger) Middleware {
	return Preprocess(func(ctx context.Context, fullMethod string) (context.Context, error) {
		ctx, names := withNames(ctx, fullMethod)
		log := log.WithFields(logrus.Fields{
			"service": names.Service,
			"method":  names.Method,
		})
		return rpccontext.WithLogger(ctx, log), nil
	})
}
