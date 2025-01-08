package models

type Consumer struct {
	NIK          string          `json:"nik"`
	FullName     string          `json:"full_name"`
	LegalName    string          `json:"legal_name"`
	PlaceOfBirth string          `json:"place_of_birth"`
	DateOfBirth  string          `json:"date_of_birth"`
	Salary       float64         `json:"salary"`
	Limit        map[int]float64 `json:"limit"` // tenor: amount
	CreditScore  int             `json:"credit_score"`
	Status       string          `json:"status"`
}
