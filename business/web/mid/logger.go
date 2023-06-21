package mid

import (
	"context"
	"github.com/mamalmaleki/go_sales_test/foundation/web"
	"go.uber.org/zap"
	"net/http"
	"time"
)

// Logger ..
func Logger(log *zap.SugaredLogger) web.Middleware {
	m := func(handler web.Handler) web.Handler {
		h := func(ctx context.Context, w http.ResponseWriter, r *http.Request) error {
			//traceID := "0000000000000000"
			//statusCode := http.StatusOK
			//now := time.Now()

			// If the context is missing this value, request the service
			// to be shutdown gracefully.
			v, err := web.GetValues(ctx)
			if err != nil {
				return err // web.NewShutdownError("web value missing from context)
			}

			log.Infow("request started",
				"trace_id", v.TraceID,
				"method", r.Method,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr)

			// Call the next handler.
			err = handler(ctx, w, r)

			log.Infow("request completed",
				"trace_id", v.TraceID,
				"method", r.Method,
				"path", r.URL.Path,
				"remote_addr", r.RemoteAddr,
				"status_code", v.StatusCode,
				"since", time.Since(v.Now),
			)

			// Return the error, so it can be handled further up the chain.
			return err
		}
		return h
	}
	return m
}
