package customCache

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Message struct {
	Sequence primitive.ObjectID `bson:"_id,omitempty"`
	Message  string             `bson:"message"`
	Received bool               `bson:"received"`
}
