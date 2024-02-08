package domain

import "context"

type App struct {
	Info   AppInfo
	Name   string
	Secret string
}

type AppInfo struct {
	ID string
}

type AppRepository interface {
	Get(ctx context.Context, id string) (App, error)
}
