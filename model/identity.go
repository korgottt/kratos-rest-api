package model

import (
	"time"
)

// Identity1 main entity in the system
type Identity1 struct {
	ID                string `json:"id"`
	RecoveryAddresses []struct {
		ID    string `json:"id"`
		Value string `json:"value"`
		Via   string `json:"via"`
	} `json:"recovery_addresses"`
	SchemaID  string `json:"schema_id"`
	SchemaURL string `json:"schema_url"`
	Traits    struct {
	} `json:"traits"`
	VerifiableAddresses []struct {
		ExpiresAt  time.Time `json:"expires_at"`
		ID         string    `json:"id"`
		Value      string    `json:"value"`
		Verified   bool      `json:"verified"`
		VerifiedAt time.Time `json:"verified_at"`
		Via        string    `json:"via"`
	} `json:"verifiable_addresses"`
}



//-----------------------------------------------------------

// Identity2 main entity in the system
type Identity2 struct {
	ID        string                `json:"id"`
	RecAddr   []RecoveryAddresses   `json:"recovery_addresses"`
	SchemaID  string                `json:"schema_id"`
	SchemaURL string                `json:"schema_url"`
	Traits    struct{}              `json:"traits"`
	VerifAddr []VerifiableAddresses `json:"verifiable_addresses"`
}

// RecoveryAddresses ...
type RecoveryAddresses struct {
	ID    string `json:"id"`
	Value string `json:"value"`
	Via   string `json:"via"`
}

// VerifiableAddresses ...
type VerifiableAddresses struct {
	ExpiresAt  time.Time `json:"expires_at"`
	ID         string    `json:"id"`
	Value      string    `json:"value"`
	Verified   bool      `json:"verified"`
	VerifiedAt time.Time `json:"verified_at"`
	Via        string    `json:"via"`
}

// VerifiableAddressesWrap response wrap
type VerifiableAddressesWrap struct {
	VerifiableAddresses VerifiableAddresses `json:"verifiable_addresses"`
}

// RecoveryAddressesWrap response wrap
type RecoveryAddressesWrap struct {
	RecoveryAddresses RecoveryAddresses `json:"verifiable_addresses"`
}
