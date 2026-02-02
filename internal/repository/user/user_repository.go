package repository_user

import (
	"context"
	"time"

	"github.com/vladislavkovaliov/ledger/internal/domain/user"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type UserRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(col *mongo.Collection) *UserRepository {
	return &UserRepository{collection: col}
}

func (r *UserRepository) Save(ctx context.Context, u *user.User) error {
	if u.CreatedAt.IsZero() {
		u.CreatedAt = time.Now()
	}

	doc := toDocument(u)

	res, err := r.collection.InsertOne(ctx, doc)

	u.ID = res.InsertedID.(bson.ObjectID).Hex()

	return err
}

func (r *UserRepository) FindByEmail(ctx context.Context, email string) (*user.User, error) {
	var doc userDocument

	err := r.collection.FindOne(ctx, bson.M{"email": email}).Decode(&doc)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, nil
		}

		return nil, err
	}

	return toDomain(&doc), nil
}

func (r *UserRepository) FindByID(ctx context.Context, id string) (*user.User, error) {
	var doc userDocument

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)

	if err != nil {
		return nil, err
	}

	return toDomain(&doc), nil
}

func (r *UserRepository) List(ctx context.Context) ([]*user.UserResponse, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var users []*user.UserResponse

	for cursor.Next(ctx) {
		var doc userDocument

		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		users = append(users, toUserResponseDomain(&doc))
	}

	return users, nil
}
