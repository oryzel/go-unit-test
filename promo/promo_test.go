package promo

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPromoCalculate(t *testing.T) {
	type testCase struct {
		name               string
		transactionAmount  int
		minimumAmount      int
		discountPercentage int
		expectedAmount     int
		wantErr            bool
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
		{
			name:               "It should return error",
			transactionAmount:  0,
			minimumAmount:      50000,
			discountPercentage: 10,
			expectedAmount:     0,
			wantErr:            true,
		},
		// {
		// 	name:               "It should return error wrong example",
		// 	transactionAmount:  10,
		// 	minimumAmount:      50000,
		// 	discountPercentage: 10,
		// 	expectedAmount:     0,
		// 	wantErr:            true,
		// },
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			promo := NewPromo(tc.minimumAmount, tc.discountPercentage)
			result, err := promo.Calculate(tc.transactionAmount)

			if tc.wantErr && err == nil {
				t.Fatalf("Minimum transaction amount shoul be greater than zero")
			} else {
				assert.Equal(t, tc.expectedAmount, result)
			}

		})

	}

}

func TestPromoCalculateShouldFailWithZeroAmount(t *testing.T) {
	promo := NewPromo(50000, 10)
	_, err := promo.Calculate(0)

	if err == nil {
		t.Fatalf("Minimum transaction amount shoul be greater than zero")
	}
}
