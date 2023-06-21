package testgrp

import (
	"context"
	"github.com/mamalmaleki/go_sales_test/foundation/web"
	"go.uber.org/zap"
	"log"
	"net/http"
)

// Handlers manages the set of check endpoints.
type Handlers struct {
	Log *zap.SugaredLogger
	//DB    *sqlx.DB
}

func (h Handlers) Test(ctx context.Context, w http.ResponseWriter,
	r *http.Request) error {
	status := struct {
		Status string
	}{
		Status: "OK",
	}

	v, err := web.GetValues(ctx)
	if err != nil {
		return err // web.NewShutdownError("web value missing from context)
	}
	log.Println("*************")
	log.Println(v.TraceID)

	return web.Respond(ctx, w, status, http.StatusOK)
}
