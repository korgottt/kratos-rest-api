package server

import (
	"strconv"

	"github.com/korgottt/kratos-rest-api/model"

	"github.com/gofiber/fiber"
)

// IdentitiesStore ...
type IdentitiesStore interface {
	GetIdentitiesList(ctx *fiber.Ctx)
	GetIdentityByID(ctx *fiber.Ctx)
	CreateIdentity(ctx *fiber.Ctx)
	UpdateIdentity(ctx *fiber.Ctx)
	DeleteIdentity(ctx *fiber.Ctx)
}

var identities = []model.Identity1{}

//NewApplication create a new fiber application
func NewApplication() *fiber.App {
	app := fiber.New()

	setupRoutes(app)

	return app
}

func setupRoutes(app *fiber.App) {
	app.Get("/identities", GetIdentitiesList)
	app.Get("/identities/:id", GetIdentityByID)
	app.Post("/identities", CreateIdentity)
	app.Delete("/identities/:id", DeleteIdentity)
}

// GetIdentitiesList get list all identities in the system
func GetIdentitiesList(ctx *fiber.Ctx) {
	ctx.Accepts("application/json")
	err := ctx.Status(fiber.StatusOK).JSON(identities)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusInternalServerError,
				Debug:   "Cannot get identities",
				Message: err.Error(),
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
	}
}

// GetIdentityByID get single identity by id
func GetIdentityByID(ctx *fiber.Ctx) {
	//v1 if param is int
	paramsID := ctx.Params("id")
	_, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Element was not found",
				Message: err.Error(),
				Reason:  "cannot parse id",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
		return
	}

	//v2 if param is uuid

	for _, identity := range identities {
		if identity.ID == paramsID {
			ctx.Status(fiber.StatusOK).JSON(identity)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

// CreateIdentity ...
func CreateIdentity(ctx *fiber.Ctx) {
	var body model.Identity1
	err := ctx.BodyParser(&body)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Element was not created",
				Message: err.Error(),
				Reason:  "cannot parse json",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
		return
	}

	newidentity := model.Identity1{
		ID:                  strconv.Itoa(len(identities) + 1),
		RecoveryAddresses:   body.RecoveryAddresses,
		SchemaID:            body.SchemaID,
		SchemaURL:           body.SchemaURL,
		Traits:              body.Traits,
		VerifiableAddresses: body.VerifiableAddresses,
	}

	identities = append(identities, newidentity)

	err = ctx.Status(fiber.StatusOK).JSON(newidentity)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError)
	}
}

// DeleteIdentity delete identity by id
func DeleteIdentity(ctx *fiber.Ctx) {
	//v1 if param is int
	paramsID := ctx.Params("id")
	_, err := strconv.Atoi(paramsID)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Element was not deleted",
				Message: err.Error(),
				Reason:  "cannot parse id",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
		return
	}

	//v2 if param is uuid

	for i, identity := range identities {
		if identity.ID == paramsID {
			identities = append(identities[:i], identities[i+1:]...)
			ctx.Status(fiber.StatusNoContent)
			return
		}
	}

	ctx.Status(fiber.StatusNotFound)
}

//v2 if param is uuid
// isValidUUID:= utils.IsValidUUID(paramsID)
// if !isValidUUID {
// 	ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
// 		"error": "cannot parse id",
// 	})
// 	return
// }
