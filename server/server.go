package server

import (
	"github.com/korgottt/kratos-rest-api/model"

	"github.com/gofiber/fiber"
)

// IdentitiesStore ...
type IdentitiesStore interface {
	GetIdentitiesList() ([]model.Identity, error)
	GetIdentityByID(id string) (model.Identity, error)
	CreateIdentity(identity model.Identity) (model.Identity, error)
	UpdateIdentity(id string, i model.Identity) (model.Identity, error)
	DeleteIdentity(id string) error
}

//App ...
type App struct {
	store  IdentitiesStore
	Server *fiber.App
}



//NewApplication create a new fiber application
func NewApplication(s IdentitiesStore) *App {
	app := &App{store : s, Server: fiber.New()}

	SetupRoutes(app)

	return app
}
