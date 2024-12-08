package model

// BankAccount ...
type BankAccount struct {
	BaseModel
	CompanyID   uint64 `json:"companyID" gorm:"index"`
	BankName    string `json:"bankName"`
	BranchName  string `json:"branchName"`
	AccountNo   string `json:"accountNo"`
	AccountName string `json:"accountName"`
}
