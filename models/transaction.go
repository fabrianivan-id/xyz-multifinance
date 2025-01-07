package models

type Transaction struct {
	ContractNumber string  `json:"contract_number"`
	NIK            string  `json:"nik"`
	OTR            float64 `json:"otr"`
	AdminFee       float64 `json:"admin_fee"`
	Installment    float64 `json:"installment"`
	Interest       float64 `json:"interest"`
	AssetName      string  `json:"asset_name"`
}