package main

import (
	"github.com/korgottt/kratos-rest-api/server"
	"github.com/rs/zerolog/log"
)

func main() {
	app := server.NewApplication()

	if err := app.Listen(3000); err != nil {
		log.Info().Msgf("could not listen on port 3000 %v", err)
	}

}
