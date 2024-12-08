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

// CalculateFeeAmount is for calculate fee amount
func (f Fee) CalculateFeeAmount(amount int64) int64 {
	return (amount * f.FeeRate) / util.Pow10(f.FeeScale)
}

// CalculateTaxAmount is for calculate tax amount
func (f Fee) CalculateTaxAmount(amount int64) int64 {
	return (amount * f.TaxRate) / util.Pow10(f.TaxScale)
}
