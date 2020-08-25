package server

// SetupRoutes setup api routes for fiber application
func SetupRoutes(app *App) {
	identities := app.Server.Group("/identities")

	identities.Get("/", app.GetIdentitiesList)
	identities.Get("/:id", app.GetIdentityByID)
	identities.Post("/", app.CreateIdentity)
	identities.Delete("/:id", app.DeleteIdentity)
	identities.Put("/:id", app.UpdateIdentity)
}
