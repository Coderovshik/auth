package app

import (
	"context"
	"log"
	"net"

	"github.com/Coderovshik/auth/internal/api"
	"github.com/Coderovshik/auth/internal/api/controller"
	"github.com/Coderovshik/auth/internal/config"
	"github.com/Coderovshik/auth/internal/repository"
	"github.com/Coderovshik/auth/internal/usecase"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	desc "github.com/Coderovshik/auth/pkg/auth"
)

type App struct {
	cfg           *config.Config
	grpcServer    *grpc.Server
	serviceServer desc.AuthServer
	mongoClient   *mongo.Client
}

func New(ctx context.Context, cfg *config.Config) *App {
	// TODO: configurate storage

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(cfg.MongoURI()),
	)
	if err != nil {
		log.Fatalf("FATAL: failed to create database connection: %s", err.Error())
	}

	// TODO: configurate serviceServer

	userRepo := repository.NewUser(
		client.Database(cfg.MongoDB.Database).Collection("user"),
	)
	appRepo := repository.NewApp(
		client.Database(cfg.MongoDB.Database).Collection("app"),
	)

	loginUsecase := usecase.NewLogin(userRepo, appRepo, cfg.App.TokenTTL)
	registerUsecase := usecase.NewRegister(userRepo)
	isAdminUsecase := usecase.NewIsAdmin(userRepo)

	loginController := controller.NewLogin(loginUsecase)
	registerController := controller.NewRegister(registerUsecase)
	isAdminController := controller.NewIsAdmin(isAdminUsecase)

	serviceServer := api.NewServer(
		loginController,
		isAdminController,
		registerController,
	)

	return &App{
		cfg:           cfg,
		mongoClient:   client,
		serviceServer: serviceServer,
	}
}

func (a *App) startGRPC() error {
	addr := a.cfg.AddressGRPC()
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("FATAL: failed to initialize listener address=%s", addr)
	}

	var serverOptions []grpc.ServerOption

	a.grpcServer = grpc.NewServer(serverOptions...)
	desc.RegisterAuthServer(a.grpcServer, a.serviceServer)
	reflection.Register(a.grpcServer)

	log.Printf("grpc server is running address=%s\n", lis.Addr().String())
	return a.grpcServer.Serve(lis)
}

func (a *App) Run() {
	const op = "app.Run"

	if err := a.startGRPC(); err != nil {
		log.Fatalf("FATAL: failed to start gRPC server: %s", err.Error())
	}
}

func (a *App) Stop(ctx context.Context) {
	const op = "app.Stop"

	log.Printf("stopping gRPC server addr=%s\n", a.cfg.AddressGRPC())

	a.grpcServer.GracefulStop()
	if err := a.mongoClient.Disconnect(ctx); err != nil {
		log.Fatalf("FATAL: failed to close db connection: %s", err.Error())
	}

	log.Println("server stopped")
}
