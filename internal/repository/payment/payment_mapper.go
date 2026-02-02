package repository_paymet

import (
	"time"

	"github.com/vladislavkovaliov/ledger/internal/domain/payment"
	"go.mongodb.org/mongo-driver/v2/bson"
)

type paymentDocument struct {
	ID          bson.ObjectID `bson:"_ID,omitempty"`
	Amount      int64         `bson:"Amount"`
	CategoryID  string        `bson:"categoryId,omitempty"`
	CreatedDate time.Time     `bson:"createdDate"`
	UpdatedDate time.Time     `bson:"updatedDate"`
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
		ID:          id,
		Amount:      p.Amount,
		CategoryID:  p.CategoryID,
		CreatedDate: p.CreateDate,
		UpdatedDate: p.UpdateDate,
	}
}

func toDomain(doc *paymentDocument) *payment.Payment {
	return &payment.Payment{
		ID:         doc.ID.Hex(),
		Amount:     doc.Amount,
		CategoryID: doc.CategoryID,
		CreateDate: doc.CreatedDate,
		UpdateDate: doc.UpdatedDate,
	}
}
