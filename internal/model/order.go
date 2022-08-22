package model

import (
	"fmt"
	"time"
)

type Order struct {
	ID          string    `json:"id"`
	BankID      string    `json:"bank_id" validate:"required"`
	Amount      int64     `json:"amount" validate:"required"`
	Information string    `json:"information"`
	CreatedOn   time.Time `json:"-"`
}

func (order Order) String() string {
	return fmt.Sprintf(`Order ID: %s. Bank ID: %s. Amount: "%d Euros". Information: "%s".`, order.ID, order.BankID, order.Amount, order.Information)
}
