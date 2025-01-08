package models

import "time"

type Transaction struct {
	ContractNumber string    `json:"contract_number"`
	NIK            string    `json:"nik"`
	OTR            float64   `json:"otr"`
	AdminFee       float64   `json:"admin_fee"`
	Tenor          int       `json:"tenor"`
	Installment    float64   `json:"installment"`
	Interest       float64   `json:"interest"`
	AssetName      string    `json:"asset_name"`
	Status         string    `json:"status"`        // New field for credit status
	CreditReason   string    `json:"credit_reason"` // New field for reason if credit is stuck/not paid
	Date           time.Time `json:"date"`          // New field for transaction date and time
	DueDate        time.Time `json:"due_date"`      // New field for due date of every tenor
}
