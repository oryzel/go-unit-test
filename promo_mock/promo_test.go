package promo_mock

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type PromoRepositoryMock struct {
	mock.Mock
}

func (prm PromoRepositoryMock) FindCurrentPromo() int {
	args := prm.Called()

	return args.Int(0)
}

func TestPromoCalculate(t *testing.T) {

	type testCase struct {
		name               string
		transactionAmount  int
		minimumAmount      int
		discountPercentage int
		expectedAmount     int
	}

	testCases := []testCase{
		{
			name:               "It should apply 63000",
			transactionAmount:  70000,
			minimumAmount:      50000,
			discountPercentage: 10,
			expectedAmount:     63000,
		},
		{
			name:               "It should apply 40000",
			transactionAmount:  40000,
			minimumAmount:      50000,
			discountPercentage: 10,
			expectedAmount:     40000,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			promoRepositoryMock := PromoRepositoryMock{}
			promoRepositoryMock.On("FindCurrentPromo").Return(tc.discountPercentage)

			promo := NewPromo(tc.minimumAmount, promoRepositoryMock)
			result, _ := promo.Calculate(tc.transactionAmount)

			assert.Equal(t, tc.expectedAmount, result)
		})

	}
}

func TestPromoCalculateShouldFailWithZeroAmount(t *testing.T) {

	promoRepositoryMock := PromoRepositoryMock{}
	promo := NewPromo(50000, promoRepositoryMock)
	_, err := promo.Calculate(0)

	if err == nil {
		t.Fatalf("Minimum transaction amount shoul be greater than zero")
	}
}
