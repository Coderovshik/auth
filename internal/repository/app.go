package repository

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/Coderovshik/auth/internal/domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ domain.AppRepository = (*App)(nil)

type App struct {
	coll *mongo.Collection
}

func NewApp(appColl *mongo.Collection) *App {
	return &App{
		coll: appColl,
	}
}

func (a *App) Get(ctx context.Context, id string) (domain.App, error) {
	const op = "repository.App.Get"

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("ERROR: invalid id: %s", err.Error())

		return domain.App{}, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.D{
		{"_id", objectID},
	}

	var result appDocument
	err = a.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("ERROR: app not found: %s\n", err.Error())

			return domain.App{}, fmt.Errorf("%s: %w", op, ErrAppNotFound)
		}

		log.Printf("ERROR: failed to find document: %s\n", err.Error())

		return domain.App{}, fmt.Errorf("%s: %w", op, err)
	}

	return toDomainApp(result), nil
}
