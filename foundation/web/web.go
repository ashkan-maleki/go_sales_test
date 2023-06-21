package web

import (
	"context"
	"github.com/dimfeld/httptreemux/v5"
	"github.com/google/uuid"
	"net/http"
	"os"
	"syscall"
	"time"
)

// Handler is a type that handles a http request within out own little mux
// mini framework.
type Handler func(ctx context.Context, w http.ResponseWriter, r *http.Request) error

// App is the entrypoint into out application and what configures our context
// object for each of out http handlers. Feel free to add any configuration
// data/logic on this App struct.
type App struct {
	*httptreemux.ContextMux
	shutdown chan os.Signal
	mw       []Middleware
}

// NewApp creates an App value that handles a set of routes for the application.
func NewApp(shutdown chan os.Signal, mw ...Middleware) *App {
	return &App{
		ContextMux: httptreemux.NewContextMux(),
		shutdown:   shutdown,
		mw:         mw,
	}
}

// SignalShutdown is used to gracefully shut down the app when an integrity
// issue is identified.
func (a *App) SignalShutdown() {
	a.shutdown <- syscall.SIGTERM
}

// Handle sets a handler function for a given HTTP method and path pair
// to the application server mux.
func (a *App) Handle(method, group, path string, handler Handler, mw ...Middleware) {
	// First wrap handler specific middleware around this handler.
	handler = wrapMiddleware(mw, handler)

	// Add the application's general middleware to the handler chain.
	handler = wrapMiddleware(a.mw, handler)

	// The function to execute for each request.
	h := func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		// Set the context with the required values to process the request.
		v := Values{
			//TraceID: span.SpanContext().TraceID().String(),
			TraceID: uuid.New().String(),
			Now:     time.Now(),
		}
		ctx = context.WithValue(ctx, key, &v)

		// Call the wrapped handler function.
		if err := handler(ctx, w, r); err != nil {
			a.SignalShutdown()
			return
		}

		// PRE CODE PROCESSING
		// Call the wrapped handler function.
		if err := handler(r.Context(), w, r); err != nil {
			// Error handling
			return
		}
	}
	finalPath := path
	if group != "" {
		finalPath = "/" + group + path
	}
	a.ContextMux.Handle(method, finalPath, h)
}
