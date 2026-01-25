package repository_user

import (
	"time"

	"github.com/vladislavkovaliov/ledger/internal/domain/user"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type userDocument struct {
	ID           bson.ObjectID `bson:"_id,omitempty"`
	Email        string        `json:"email" bson:"email"`
	PasswordHash string        `json:"-" bson:"password_hash"`
	CreatedAt    time.Time     `json:"created_at" bson:"created_at"`
}

func toDocument(p *user.User) *userDocument {
	var id bson.ObjectID
	var err error

	if p.ID != "" {
		id, err = bson.ObjectIDFromHex(p.ID)

		if err != nil {
			return nil
		}
	}

	return &userDocument{
		ID:           id,
		Email:        p.Email,
		PasswordHash: p.PasswordHash,
		CreatedAt:    p.CreatedAt,
	}
}

func toDomain(doc *userDocument) *user.User {
	return &user.User{
		ID:           doc.ID.Hex(),
		Email:        doc.Email,
		PasswordHash: doc.PasswordHash,
		CreatedAt:    doc.CreatedAt,
	}
}

func toUserResponseDomain(doc *userDocument) *user.UserResponse {
	return &user.UserResponse{
		ID:        doc.ID.Hex(),
		Email:     doc.Email,
		CreatedAt: doc.CreatedAt,
	}
}
