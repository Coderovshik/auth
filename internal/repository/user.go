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

var _ domain.UserRepository = (*User)(nil)

type User struct {
	coll *mongo.Collection
}

func NewUser(userColl *mongo.Collection) *User {
	return &User{
		coll: userColl,
	}
}

func (u *User) Save(ctx context.Context, email string, passHash []byte) (string, error) {
	const op = "repository.User.Save"

	res, err := u.coll.InsertOne(
		ctx,
		bson.D{
			{"email", email},
			{"pass_hash", passHash},
		},
	)
	if err != nil {
		log.Printf("ERROR: failed to insert document: %s\n", err.Error())

		return "", fmt.Errorf("%s: %w", op, err)
	}

	hex := res.InsertedID.(primitive.ObjectID).Hex()
	return hex, nil
}

func (u *User) Get(ctx context.Context, email string) (domain.User, error) {
	const op = "repository.User.Get"

	filter := bson.D{
		{"email", email},
	}

	var result userDocument
	err := u.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("ERROR: user not found: %s\n", err.Error())

			return domain.User{}, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		log.Printf("ERROR: failed to find document: %s\n", err.Error())

		return domain.User{}, fmt.Errorf("%s: %w", op, err)
	}

	return toDomainUser(result), nil
}

func (u *User) IsAdmin(ctx context.Context, id string) (bool, error) {
	const op = "repository.User.IsAdmin"

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Printf("ERROR: invalid id: %s", err.Error())

		return false, fmt.Errorf("%s: %w", op, err)
	}

	filter := bson.D{
		{"_id", objectID},
	}

	var result userDocument
	err = u.coll.FindOne(ctx, filter).Decode(&result)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			log.Printf("ERROR: user not found: %s\n", err.Error())

			return false, fmt.Errorf("%s: %w", op, ErrUserNotFound)
		}

		log.Printf("ERROR: failed to find document: %s\n", err.Error())

		return false, fmt.Errorf("%s: %w", op, err)
	}

	return result.IsAdmin, nil
}
