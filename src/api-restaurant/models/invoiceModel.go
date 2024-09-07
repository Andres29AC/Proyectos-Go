package models

import(
	"time"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Invoice struct{
	ID                primitive.ObjectID       `bson:"_id"`
	Invoice_id        string                   `json:"invoice_id"`
	Order_id		  string                   `json:"order_id"`
	Pay_method       *string                   `json:"pay_method" validate:"eq=CARD|eq=CASH|eq="`
	Pay_status       *string                   `json:"pay_status" validate:"required,eq=PAID|eq=PENDING|eq=FAILED|eq="`
	Pay_due_date     time.Time                 `json:"pay_due_date"`
	Created_at       time.Time                 `json:"created_at"`
	Updated_at       time.Time                 `json:"updated_at"`
}

//NOTE:
// *string: el asterisco indica que el campo puede ser nulo
// eq=CARD|eq=CASH|eq=": solo puede ser CARD o CASH
