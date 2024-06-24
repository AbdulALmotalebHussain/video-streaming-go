package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Video struct {
    ID       primitive.ObjectID `bson:"_id,omitempty"`
    Filename string             `bson:"filename"`
    Path     string             `bson:"path"`
}

