package repository

import "go.mongodb.org/mongo-driver/bson/primitive"

type userDocument struct {
	ID       primitive.ObjectID `bson:"_id"`
	Email    string             `bson:"email"`
	PassHash []byte             `bson:"pass_hash"`
	IsAdmin  bool               `bson:"is_admin,omitempty"`
}

type appDocument struct {
	ID     primitive.ObjectID `bson:"_id"`
	Name   string             `bson:"name"`
	Secret string             `bson:"secret"`
}
