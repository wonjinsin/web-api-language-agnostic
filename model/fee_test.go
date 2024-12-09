package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFee_CalculateFeeWithTaxAmount(t *testing.T) {
	tests := []struct {
		name     string
		fee      Fee
		amount   int64
		expected int64
	}{
		{
			name: "basic calculation",
			fee: Fee{
				FeeRate:  4,
				TaxRate:  110,
				FeeScale: 2,
				TaxScale: 2,
			},
			amount:   10000,
			expected: 440,
		},
		{
			name: "basic calculation2",
			fee: Fee{
				FeeRate:  200,
				TaxRate:  110,
				FeeScale: 4,
				TaxScale: 2,
			},
			amount:   1000000,
			expected: 22000,
		},
		{
			name: "zero amount",
			fee: Fee{
				FeeRate:  100,
				TaxRate:  110,
				FeeScale: 4,
				TaxScale: 2,
			},
			amount:   0,
			expected: 0,
		},
		{
			name: "different scales",
			fee: Fee{
				FeeRate:  150,
				TaxRate:  108,
				FeeScale: 4,
				TaxScale: 2,
			},
			amount:   2000000,
			expected: 32400,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.fee.CalculateFeeWithTaxAmount(tt.amount)
			assert.Equal(t, tt.expected, result)
		})
	}
}
