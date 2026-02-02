package repository_paymet

import (
	"time"

	"github.com/vladislavkovaliov/ledger/internal/domain/payment"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type paymentDocument struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Amount     int64         `json:"Amount" bson:"amount"`
	CategoryID string        `json:"CategoryID" bson:"category_id,omitempty"`
	CreatedAt  time.Time     `json:"CreatedAt" bson:"created_at"`
	UpdatedAt  time.Time     `json:"UpdatedAt" bson:"updated_at"`
}

func toDocument(p *payment.Payment) *paymentDocument {
	var id bson.ObjectID
	var err error

	if p.ID != "" {
		id, err = bson.ObjectIDFromHex(p.ID)

		if err != nil {
			return nil
		}
	}

	return &paymentDocument{
		ID:         id,
		Amount:     p.Amount,
		CategoryID: p.CategoryID,
		CreatedAt:  p.CreatedAt,
		UpdatedAt:  p.UpdatedAt,
	}
}

func toDomain(doc *paymentDocument) *payment.Payment {
	return &payment.Payment{
		ID:         doc.ID.Hex(),
		Amount:     doc.Amount,
		CategoryID: doc.CategoryID,
		CreatedAt:  doc.CreatedAt,
		UpdatedAt:  doc.UpdatedAt,
	}
}
