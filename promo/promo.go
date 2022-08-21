package promo

import "errors"

type Promo struct {
	minimumPurchase    int
	discountPercentage int
}

func NewPromo(minimumPurchase int, discouontPercentage int) *Promo {
	return &Promo{
		minimumPurchase:    minimumPurchase,
		discountPercentage: discouontPercentage,
	}
}

func (p *Promo) Calculate(transactionAmount int) (int, error) {

	if transactionAmount <= 0 {
		return 0, errors.New("Minimum transaction amount shoul be greater than zero")
	}

	if transactionAmount > p.minimumPurchase {
		transactionAmount = transactionAmount - (transactionAmount * p.discountPercentage / 100)
	}

	return transactionAmount, nil

}
