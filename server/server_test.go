package server

import (
	"fmt"
	"net/http"
	"testing"

	"github.com/korgottt/kratos-rest-api/model"
)

type StubBlogStore struct {
	identities []model.Identity1
}

func TestGetIdentities(t *testing.T) {

	app := NewApplication()
	req, _ := http.NewRequest(http.MethodGet, "/identities", nil)
	resp, _ := app.Test(req)

	assertStatus(t, resp.StatusCode, http.StatusOK)
}

func TestGetIdentityById(t *testing.T) {
	app := NewApplication()
	tests := []model.Identity1{
		{
			ID: "1",
		},
		{
			ID: "2",
		},
	}
	for _, tc := range tests {

		req := createGetIdentitiesRequest(tc.ID)
		resp, _ := app.Test(req)
		assertStatus(t, resp.StatusCode, http.StatusOK)
	}
}

func TestCreateIdentity(t *testing.T) {
	app := NewApplication()
	req, _ := http.NewRequest(http.MethodPost, "/identities", nil)
	resp, _ := app.Test(req)
	assertStatus(t, resp.StatusCode, http.StatusOK)
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
