package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/Coderovshik/auth/internal/app"
	"github.com/Coderovshik/auth/internal/config"
)

func main() {
	// godotenv.Load("config/auth.env")

	cfg := config.Get()

	ctx := context.Background()
	application := app.New(ctx, cfg)

	go application.Run()

	waitForSignal()
	application.Stop(ctx)
}

func waitForSignal() {
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT, syscall.SIGQUIT)

	sig := <-stop
	log.Printf("stopping app signal=%s", sig.String())
}
