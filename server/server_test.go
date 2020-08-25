package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/korgottt/kratos-rest-api/model"
	"github.com/korgottt/kratos-rest-api/utils"
)

const notFound = "not found"

type stubStore struct {
	identities []model.Identity
}

func (s *stubStore) GetIdentitiesList() ([]model.Identity, error) {
	return s.identities, nil
}

func (s *stubStore) GetIdentityByID(id string) (model.Identity, error) {
	var r model.Identity
	e := fmt.Errorf(notFound)
	for _, v := range s.identities {
		if v.ID == id {
			r = v
			e = nil
			break
		}
	}
	return r, e
}

func (s *stubStore) CreateIdentity(i model.Identity) (model.Identity, error) {
	s.identities = append(s.identities, i)
	return i, nil
}

func (s *stubStore) UpdateIdentity(id string, i model.Identity) (model.Identity, error) {
	e := s.DeleteIdentity(id)
	if e != nil {
		return model.Identity{}, e
	}
	return s.CreateIdentity(i)
}

func (s *stubStore) DeleteIdentity(id string) error {
	var pos int
	e := fmt.Errorf(notFound)
	for i, v := range s.identities {
		if v.ID == id {
			pos = i
			e = nil
			break
		}
	}
	if e != nil {
		return e
	}
	s.identities = append(s.identities[:pos], s.identities[pos+1:]...)
	return nil
}

func (s *stubStore) NoRows(e error) bool {
	return e.Error() == notFound
}

func TestGetIdentities(t *testing.T) {
	store := stubStore{identities: []model.Identity{model.Identity{ID: utils.GenerateUUID()}, model.Identity{ID: utils.GenerateUUID()}}}
	app := NewApplication(&store)
	req, _ := http.NewRequest(http.MethodGet, "/identities", nil)
	resp, _ := app.Server.Test(req)

	assertStatus(t, resp.StatusCode, http.StatusOK)
}

func TestGetIdentityById(t *testing.T) {
	store := stubStore{identities: []model.Identity{model.Identity{ID: utils.GenerateUUID()}, model.Identity{ID: utils.GenerateUUID()}}}
	app := NewApplication(&store)
	tests := []model.Identity{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}
	for _, tc := range tests {

		req := createGetIdentitiesRequest(tc.ID)
		resp, _ := app.Server.Test(req)
		assertStatus(t, resp.StatusCode, http.StatusOK)
	}
}

func TestCreateIdentity(t *testing.T) {
	store := stubStore{identities: []model.Identity{model.Identity{ID: utils.GenerateUUID()}, model.Identity{ID: utils.GenerateUUID()}}}
	app := NewApplication(&store)
	req, _ := http.NewRequest(http.MethodPost, "/identities", nil)
	resp, _ := app.Server.Test(req)
	assertStatus(t, resp.StatusCode, http.StatusOK)
}

func TestUpdateIdentity(t *testing.T) {

}

func assertStatus(t *testing.T, got, want int) {
	t.Helper()
	if got != want {
		t.Errorf("got %d, want %d", got, want)
	}
}

func createGetIdentitiesRequest(id string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("/identities/%s", id), nil)
	return req
}
