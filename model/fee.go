package model

import "pikachu/util"

// Fee is model for fee
type Fee struct {
	BaseModel
	CountryCode util.CountryCode `json:"countryCode"`
	FeeRate     int64            `json:"feeRate"`
	TaxRate     int64            `json:"taxRate"`
	FeeScale    int              `json:"feeScale"`
	TaxScale    int              `json:"taxScale"`
}

// CalculateFeeWithTaxAmount is for calculate fee amount
func (f Fee) CalculateFeeWithTaxAmount(amount int64) int64 {
	return amount * f.FeeRate * f.TaxRate / util.Pow10(f.FeeScale) / util.Pow10(f.TaxScale)
}
