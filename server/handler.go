package server

import (
	"database/sql"

	"github.com/gofiber/fiber"
	"github.com/korgottt/kratos-rest-api/model"
	"github.com/korgottt/kratos-rest-api/utils"
	"github.com/rs/zerolog/log"
)

// GetIdentitiesList get list all identities in the system
func (a *App) GetIdentitiesList(ctx *fiber.Ctx) {
	identities, err := a.store.GetIdentitiesList()
	if err != nil {
		ctx.Send(err)
	}
	err = ctx.Status(fiber.StatusOK).JSON(identities)
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
func (a *App) GetIdentityByID(ctx *fiber.Ctx) {
	//v1 if param is int
	paramsID := ctx.Params("id")
	if !utils.IsValidUUID(paramsID) {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Item was not found",
				Message: "id is not valid, must be UUID format",
				Reason:  "cannot parse id",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
		return
	}
	identity, err := a.store.GetIdentityByID(paramsID)
	if err != nil {
		log.Info().Msgf("something got wrong %v", err)
		return
	}

	if identity.ID == "" {
		ctx.Status(fiber.StatusNotFound)
	} else {
		ctx.Status(fiber.StatusOK).JSON(identity)
	}
}

// CreateIdentity create new identity
func (a *App) CreateIdentity(ctx *fiber.Ctx) {
	var identity model.Identity
	err := ctx.BodyParser(&identity)
	if err != nil {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Item was not created",
				Message: err.Error(),
				Reason:  "cannot parse json",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
		return
	}

	newidentity, err := a.store.CreateIdentity(identity)
	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusInternalServerError,
				Debug:   "Item was not created",
				Message: err.Error(),
				Reason:  "",
				Request: ctx.Body(),
				Status:  "Internal Server Error",
			},
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(newidentity)
}

// DeleteIdentity delete identity by id
func (a *App) DeleteIdentity(ctx *fiber.Ctx) {
	paramsID := ctx.Params("id")
	if !utils.IsValidUUID(paramsID) {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Element was not deleted",
				Message: "id is not valid, must be UUID format",
				Reason:  "cannot parse id",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
	}

	err := a.store.DeleteIdentity(paramsID)
	if err == sql.ErrNoRows {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusNotFound,
				Debug:   "Item was not delete",
				Message: err.Error(),
				Reason:  "Item is missing from the database",
				Request: ctx.Body(),
				Status:  "Not found",
			},
		})
		return
	}

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusInternalServerError,
				Debug:   "Cannot get identities",
				Message: err.Error(),
				Request: ctx.Body(),
				Status: "Internal Server Error",
			},
		})
		return
	}

	ctx.Status(fiber.StatusNoContent).JSON("")
}

// UpdateIdentity update identity by id
func (a *App) UpdateIdentity(ctx *fiber.Ctx) {
	var identity model.Identity
	paramsID := ctx.Params("id")
	err := ctx.BodyParser(&identity)
	if err != nil || !utils.IsValidUUID(paramsID) {
		ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusBadRequest,
				Debug:   "Item was not Update",
				Message: err.Error(),
				Reason:  "cannot parse id",
				Request: ctx.Body(),
				Status:  "Bad Request",
			},
		})
		return
	}

	i, err := a.store.UpdateIdentity(paramsID, identity)

	if err == sql.ErrNoRows {
		ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusNotFound,
				Debug:   "Item was not found",
				Message: err.Error(),
				Reason:  "Item is missing from the database",
				Request: ctx.Body(),
				Status:  "Not found",
			},
		})
		return
	}

	if err != nil {
		ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": model.Errors{
				Code:    fiber.StatusInternalServerError,
				Debug:   "Cannot get identities",
				Message: err.Error(),
				Request: ctx.Body(),
				Status: "Internal Server Error	",
			},
		})
		return
	}

	ctx.Status(fiber.StatusOK).JSON(i)
}
