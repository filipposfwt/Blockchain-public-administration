package models

// Credential represents a verifiable credential
type Credential struct {
	ID         string  `json:"id"`
	Type       string  `json:"type"`
	Issuer     string  `json:"issuer"`
	Holder     string  `json:"holder"`
	IssuedDate string  `json:"issuedDate"`
	ExpiryDate string  `json:"expiryDate"`
	DegreeName string  `json:"degreeName"`
	University string  `json:"university"`
	GradePoint float64 `json:"gradePoint"`
}
