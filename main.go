package main

import (
	"github.com/korgottt/kratos-rest-api/datastore"
	"github.com/korgottt/kratos-rest-api/server"
	"github.com/rs/zerolog/log"
)

func main() {
	store := datastore.IdentitiesDBStore{}
	if err := store.Init(); err != nil {
		log.Info().Msgf("unable to access the database: %q", err)
	}
	defer store.Close()
	app := server.NewApplication(&store)

	if err := app.Server.Listen(3000); err != nil {
		log.Info().Msgf("could not listen on port 3000 %v", err)
	}

}
