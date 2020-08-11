package datastore

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/korgottt/go-real-world-api/model"
	"github.com/korgottt/go-real-world-api/utils"
	_ "github.com/lib/pq" // db driver
)

const ConnStr = "user=postgres password=admin dbname=postgres sslmode=disable"

// IdentitiesDBStore is implementation of article store via Postgres
type IdentitiesDBStore struct {
	db *sqlx.DB
}

// GetIdentitiesList ...
func (i *IdentitiesDBStore) GetIdentitiesList() {
	panic("not implemented") // TODO: Implement
}

// GetIdentityByID ...
func (i *IdentitiesDBStore) GetIdentityByID(id int) {
	panic("not implemented") // TODO: Implement
}

// CreateIdentity ...
func (i *IdentitiesDBStore) CreateIdentity() {
	panic("not implemented") // TODO: Implement
}

// UpdateIdentity ...
func (i *IdentitiesDBStore) UpdateIdentity() {
	panic("not implemented") // TODO: Implement
}

// DeleteIdentity ...
func (i *IdentitiesDBStore) DeleteIdentity() {
	panic("not implemented") // TODO: Implement
}

