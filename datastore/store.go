package datastore

import (
	"database/sql"
	"fmt"

	"github.com/korgottt/kratos-rest-api/model"
	"github.com/rs/zerolog/log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // db driver
)

const schema = `
	CREATE TABLE IF NOT EXISTS identity (
		id uuid not null,
		schema_id text not null,
		schema_url text
	);

	CREATE TABLE IF NOT EXISTS recoveryaddress (
		id uuid not null,
		value text,
		via text,
		identity_id text not null
	);

	CREATE TABLE IF NOT EXISTS verifiableaddresses (
		id uuid not null,
		value text,
		via text,
		expires_at DATE,
		verified bool,
		verified_at DATE,
		identity_id text not null
	);
`

// ConnStr database connection string
const ConnStr = "user=postgres password=admin dbname=postgres sslmode=disable"

// IdentitiesDBStore is implementation of article store via Postgres
type IdentitiesDBStore struct {
	db *sqlx.DB
}

// Init initializes connetion
func (s *IdentitiesDBStore) Init() (err error) {
	s.db, err = sqlx.Connect("postgres", ConnStr)
	s.db.MustExec(schema)
	return
}

// Close closes connetion
func (s *IdentitiesDBStore) Close() (err error) {
	var isConnected bool
	if isConnected, err = s.ensureConnection(); isConnected {
		err = s.db.Close()
	}
	return
}

func (s *IdentitiesDBStore) ensureConnection() (isConnected bool, err error) {
	isConnected = s.db != nil
	if !isConnected {
		err = fmt.Errorf("db connection is not initialized")
	}
	return
}

// GetIdentitiesList get list of all identities
func (s *IdentitiesDBStore) GetIdentitiesList() ([]model.Identity, error) {
	identityList := []model.Identity{}
	recoveryAddresses := []model.RecoveryAddresses{}
	verifiableAddresses := []model.VerifiableAddresses{}
	err := s.db.Select(&identityList, "SELECT * FROM identity")
	if err != nil {
		return nil, err
	}
	err = s.db.Select(&recoveryAddresses, "SELECT * FROM recoveryaddress")
	if err != nil {
		return nil, err
	}

	err = s.db.Select(&verifiableAddresses, "SELECT * FROM verifiableaddresses")
	if err != nil {
		return nil, err
	}
	return s.mapIdentityEntities(identityList, recoveryAddresses, verifiableAddresses), err
}

// GetIdentityByID get single identity by id
func (s *IdentitiesDBStore) GetIdentityByID(id string) (model.Identity, error) {
	identity := model.Identity{}
	recoveryAddresses := []model.RecoveryAddresses{}
	verifiableAddresses := []model.VerifiableAddresses{}
	err := s.db.Get(&identity, "SELECT * FROM identity WHERE id=$1", id)
	if err != nil {
		log.Info().Msgf("could not get identity %v", err)
	}
	err = s.db.Select(&recoveryAddresses, "SELECT * FROM recoveryaddress WHERE identity_id=$1", id)
	if err != nil {
		log.Info().Msgf("could not get recoveryaddress %v", err)
	}
	identity.RecoveryAddresses = recoveryAddresses
	err = s.db.Select(&verifiableAddresses, "SELECT * FROM verifiableaddresses WHERE identity_id=$1", id)
	if err != nil {
		log.Info().Msgf("could not get verifiableaddresses %v", err)
	}
	identity.VerifiableAddresses = verifiableAddresses
	return identity, err
}

// CreateIdentity create new identity
func (s *IdentitiesDBStore) CreateIdentity(identity model.Identity) (model.Identity, error) {
	t, e := s.createTx(&identity, nil)
	if e != nil && t != nil {
		t.Rollback()
		return identity, e
	}
	return identity, t.Commit()
}

func (s *IdentitiesDBStore) createTx(i *model.Identity, existingTx *sql.Tx) (*sql.Tx, error) {
	var err error
	if existingTx == nil {
		existingTx, err = s.db.Begin()
	}
	if err != nil {
		return existingTx, err
	}
	err = s.insertIdentity(existingTx, *i)
	if err != nil {
		return existingTx, err
	}
	if len(i.RecoveryAddresses) > 0 {
		err = s.insertRecoveryAddresses(existingTx, i.ID, i.RecoveryAddresses)
		if err != nil {
			return existingTx, err
		}
	}
	if len(i.VerifiableAddresses) > 0 {
		err = s.insertVerifiableAddresses(existingTx, i.ID, i.VerifiableAddresses)
		if err != nil {
			return existingTx, err
		}
	}
	return existingTx, nil
}

// DeleteIdentity delete identity by id
func (s *IdentitiesDBStore) DeleteIdentity(id string) error {
	t, e := s.deleteTx(id, nil)
	if e != nil && t != nil {
		t.Rollback()
		return e
	}
	return t.Commit()

}

// UpdateIdentity updates identity by id
func (s *IdentitiesDBStore) UpdateIdentity(id string, i model.Identity) (model.Identity, error) {
	return i, s.execTxChain(
		func(t *sql.Tx) error {
			_, e := s.deleteTx(id, t)
			return e
		}, func(t *sql.Tx) error {
			_, e := s.createTx(&i, t)
			return e
		},
	)
}

func (s *IdentitiesDBStore) insertRecoveryAddresses(t *sql.Tx, identity string, a []model.RecoveryAddresses) error {
	cnt := len(a)
	q := "INSERT INTO recoveryaddress (id, value, via, identity_id) VALUES "
	p := []interface{}{}
	for index, a := range a {
		pos := index * 4
		q += fmt.Sprintf("($%d,$%d,$%d,$%d)", pos+1, pos+2, pos+3, pos+4)
		if index != cnt-1 {
			q += ","
		}
		p = append(p, a.ID, a.Value, a.Via, identity)
	}
	_, err := t.Exec(q, p...)
	return err
}

func (s *IdentitiesDBStore) insertVerifiableAddresses(t *sql.Tx, identity string, a []model.VerifiableAddresses) error {
	cnt := len(a)
	q := "INSERT INTO verifiableaddresses (id, value, via, expires_at, verified, verified_at, identity_id) VALUES "
	p := []interface{}{}
	for index, a := range a {
		pos := index * 7
		q += fmt.Sprintf("($%d,$%d,$%d,$%d,$%d,$%d,$%d)", pos+1, pos+2, pos+3, pos+4, pos+5, pos+6, pos+7)
		if index != cnt-1 {
			q += ","
		}
		p = append(p, a.ID, a.Value, a.Via, a.ExpiresAt, a.Verified, a.VerifiedAt, identity)
	}
	_, err := t.Exec(q, p...)
	return err
}

func (s *IdentitiesDBStore) mapIdentityEntities(identities []model.Identity, recoveryAddresses []model.RecoveryAddresses, verifiableAddresses []model.VerifiableAddresses) []model.Identity {
	var res []model.Identity
	for _, identity := range identities {
		var reqAdr []model.RecoveryAddresses
		var verAdr []model.VerifiableAddresses
		for _, ra := range recoveryAddresses {
			if ra.Identity == identity.ID {
				reqAdr = append(reqAdr, ra)
			}
		}
		identity.RecoveryAddresses = reqAdr
		for _, va := range verifiableAddresses {
			if va.Identity == identity.ID {
				verAdr = append(verAdr, va)
			}
		}
		identity.VerifiableAddresses = verAdr
		res = append(res, identity)
	}
	return res
}

func (s *IdentitiesDBStore) deleteTx(id string, existingTx *sql.Tx) (*sql.Tx, error) {
	var err error
	if existingTx == nil {
		existingTx, err = s.db.Begin()
	}
	if err != nil {
		return existingTx, err
	}
	_, err = existingTx.Exec("DELETE FROM recoveryaddress WHERE identity_id=$1", id)
	if err != nil {
		return existingTx, err
	}
	_, err = existingTx.Exec("DELETE FROM verifiableaddresses WHERE identity_id=$1", id)
	if err != nil {
		return existingTx, err
	}
	r, err := existingTx.Exec("DELETE FROM identity WHERE id=$1", id)
	if cnt, _ := r.RowsAffected(); err == nil && cnt == 0 {
		err = sql.ErrNoRows
	}
	return existingTx, err
}

func (s *IdentitiesDBStore) findIdentity(id string, identities []model.Identity) (model.Identity, error) {
	for _, i := range identities {
		if i.ID == id {
			return i, nil
		}
	}
	return model.Identity{}, fmt.Errorf("not found")
}

func (s *IdentitiesDBStore) execTxChain(operations ...func(*sql.Tx) error) error {
	t, e := s.db.Begin()
	if e != nil {
		return e
	}
	for _, op := range operations {
		if e = op(t); e != nil {
			t.Rollback()
			return e
		}
	}
	return t.Commit()
}

//NoRows returns whether error is no rows error
func (s *IdentitiesDBStore) NoRows(e error) bool {
	return e == sql.ErrNoRows
}

func (s *IdentitiesDBStore) insertIdentity(t *sql.Tx, i model.Identity) error {
	_, err := t.Exec("INSERT INTO identity (id, schema_id, schema_url) VALUES ($1, $2, $3)",
		i.ID, i.SchemaID, i.SchemaURL)
	return err
}
