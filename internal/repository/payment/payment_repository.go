package repository_paymet

import (
	"context"
	"time"

	"github.com/vladislavkovaliov/ledger/internal/domain/payment"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
)

type PaymentRepository struct {
	collection *mongo.Collection
}

func NewPaymentRepository(col *mongo.Collection) *PaymentRepository {
	return &PaymentRepository{collection: col}
}

func (r *PaymentRepository) Save(ctx context.Context, p *payment.Payment) error {
	if p.CreateDate.IsZero() {
		p.CreateDate = time.Now()
	}

	p.UpdateDate = time.Now()

	doc := toDocument(p)

	res, err := r.collection.InsertOne(ctx, doc)

	p.ID = res.InsertedID.(bson.ObjectID).Hex()

	return err
}

func (r *PaymentRepository) FindByID(ctx context.Context, id string) (*payment.Payment, error) {
	var doc paymentDocument

	err := r.collection.FindOne(ctx, bson.M{"_id": id}).Decode(&doc)

	if err != nil {
		return nil, err
	}

	return toDomain(&doc), nil
}

func (r *PaymentRepository) List(ctx context.Context) ([]*payment.Payment, error) {
	cursor, err := r.collection.Find(ctx, bson.M{})

	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)

	var payments []*payment.Payment

	for cursor.Next(ctx) {
		var doc paymentDocument

		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}

		payments = append(payments, toDomain(&doc))
	}

	return payments, nil
}
