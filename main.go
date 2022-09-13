package main

import (
	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	app := IrisRouter().InitRouter()
	srv := NewHTTPServer(app)
	app.Logger().Info("Server started on http://localhost:8080")
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			app.Logger().Info("Server closed")
			app.Logger().Fatal(err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	app.Logger().Info("Server shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		app.Logger().Fatal(err)
	}

	select {
	case <-ctx.Done():
		app.Logger().Info("Server shutdown timeout")
	default:
		app.Logger().Info("Server shutdown")
	}
}
