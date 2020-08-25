package model

import (
	"time"
)

// Identity main entity in the system
type Identity struct {
	ID                string              `json:"id" db:"id"`
	RecoveryAddresses []RecoveryAddresses `json:"recovery_addresses"`
	SchemaID          string              `json:"schema_id" db:"schema_id"`
	SchemaURL         string              `json:"schema_url" db:"schema_url"`
	VerifiableAddresses []VerifiableAddresses `json:"verifiable_addresses"`
}

// Address ...
type Address struct {
	ID    string `json:"id" db:"id"`
	Value string `json:"value" db:"value"`
	Via   string `json:"via" db:"via"`
}


// RecoveryAddresses ...
type RecoveryAddresses struct {
	Address
	Identity string `db:"identity_id"`
}

// VerifiableAddresses ...
type VerifiableAddresses struct {
	Address
	ExpiresAt  time.Time `json:"expires_at" db:"expires_at"`
	Verified   bool      `json:"verified" db:"verified"`
	VerifiedAt time.Time `json:"verified_at" db:"verified_at"`
	Identity  string    `db:"identity_id"`
}

// // VerifiableAddressesWrap response wrap
// type VerifiableAddressesWrap struct {
// 	VerifiableAddresses VerifiableAddresses `json:"verifiable_addresses"`
// }

// // RecoveryAddressesWrap response wrap
// type RecoveryAddressesWrap struct {
// 	RecoveryAddresses RecoveryAddresses `json:"verifiable_addresses"`
// }
