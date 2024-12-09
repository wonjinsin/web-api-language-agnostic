package model

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompany_SameID(t *testing.T) {
	tests := []struct {
		name     string
		company  Company
		inputID  uint64
		expected bool
	}{
		{
			name: "should return true when IDs match",
			company: Company{
				BaseModel: BaseModel{ID: 1},
				Name:      "Test Company",
			},
			inputID:  1,
			expected: true,
		},
		{
			name: "should return false when IDs don't match",
			company: Company{
				BaseModel: BaseModel{ID: 1},
				Name:      "Test Company",
			},
			inputID:  2,
			expected: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := tt.company.SameID(tt.inputID)
			assert.Equal(t, tt.expected, result)
		})
	}
}
