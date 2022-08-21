package promo_mock

import (
	"errors"

	"github.com/oryzel/go-unit-test/database"
)

type Promo struct {
	minimumPurchase int
	discountRepo    database.Discount
}

func NewPromo(minumumPurchase int, discountRepo database.Discount) *Promo {
	return &Promo{
		minimumPurchase: minumumPurchase,
		discountRepo:    discountRepo,
	}
}

func (p *Promo) Calculate(transactionAmount int) (int, error) {

	if transactionAmount <= 0 {
		return 0, errors.New("Minimum transacation amount shoul be greater than zero")
	}

	if transactionAmount > p.minimumPurchase {
		discount := p.discountRepo.FindCurrentPromo()
		transactionAmount = transactionAmount - (transactionAmount * discount / 100)
	}

	return transactionAmount, nil
}
