package model

// Company ...
type Company struct {
	BaseModel
	Name       string `json:"name"`
	OwnerName  string `json:"ownerName"`
	Number     string `json:"number"`
	PostNumber string `json:"postNumber"`
	Address    string `json:"address"`
}