package repository_paymet

import (
	"time"

	"github.com/vladislavkovaliov/ledger/internal/domain/payment"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type paymentDocument struct {
	ID         bson.ObjectID `bson:"_id,omitempty"`
	Amount     int64         `bson:"amount"`
	CategoryID string        `bson:"category_id,omitempty"`
	CreateDate time.Time     `bson:"create_date"`
	UpdateDate time.Time     `bson:"update_date"`
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
		CreateDate: p.CreateDate,
		UpdateDate: p.UpdateDate,
	}
}

func toDomain(doc *paymentDocument) *payment.Payment {
	return &payment.Payment{
		ID:         doc.ID.Hex(),
		Amount:     doc.Amount,
		CategoryID: doc.CategoryID,
		CreateDate: doc.CreateDate,
		UpdateDate: doc.UpdateDate,
	}
}
