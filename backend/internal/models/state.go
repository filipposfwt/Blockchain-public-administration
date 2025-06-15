package models

// BaseCertificate represents common fields for all certificates
type BaseCertificate struct {
	FullName           string `json:"fullName"`
	RegistrationNumber string `json:"registrationNumber"`
	Municipality       string `json:"municipality"`
	IssuedDate         string `json:"issuedDate"`
	ExpiryDate         string `json:"expiryDate"`
}

// BirthCertificate represents a birth certificate
type BirthCertificate struct {
	BaseCertificate
	FatherName string `json:"fatherName"`
	MotherName string `json:"motherName"`
	BirthDate  string `json:"birthDate"`
	BirthPlace string `json:"birthPlace"`
}

// DeathCertificate represents a death certificate
type DeathCertificate struct {
	BaseCertificate
	DeathDate  string `json:"deathDate"`
	DeathPlace string `json:"deathPlace"`
}

// MarriageCertificate represents a marriage certificate
type MarriageCertificate struct {
	BaseCertificate
	SpouseName    string `json:"spouseName"`
	MarriageDate  string `json:"marriageDate"`
	MarriagePlace string `json:"marriagePlace"`
}

// StateInfo represents the state information structure
type StateInfo struct {
	ID   string `json:"id"`
	Type string `json:"type"`
	// Certificate data based on type
	BirthCertificate    *BirthCertificate    `json:"birthCertificate,omitempty"`
	DeathCertificate    *DeathCertificate    `json:"deathCertificate,omitempty"`
	MarriageCertificate *MarriageCertificate `json:"marriageCertificate,omitempty"`
}

// State represents a state verification record
type State struct {
	StateHash string    `json:"stateHash"`
	StateInfo StateInfo `json:"stateInfo"`
}

// StateVerificationRequest represents the request payload for state verification
type StateVerificationRequest struct {
	StateInfo  StateInfo `json:"stateInfo" binding:"required"`
	PrivateKey string    `json:"privateKey" binding:"required"`
	RecordID   string    `json:"recordId,omitempty"` // Optional field for recordId
}
