package DesignPattern

import (
	"context"
	"os"
	"os/signal"
)

func Run(ctx context.Context) {
	ctx, cancel := context.WithCancel(ctx)
	quit := make(chan os.Signal, 1)
	signal.Notify(quit)
	defer func() {
		signal.Stop(quit)
		close(quit)
	}()
	go func() {
		<-quit
		cancel()
	}()

}
