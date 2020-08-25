package server

import (
	"testing"

	"github.com/gofiber/fiber"
)

func TestApp_GetIdentitiesList(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.GetIdentitiesList(tt.args.ctx)
		})
	}
}

func TestApp_GetIdentityByID(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.GetIdentityByID(tt.args.ctx)
		})
	}
}

func TestApp_CreateIdentity(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.CreateIdentity(tt.args.ctx)
		})
	}
}

func TestApp_DeleteIdentity(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.DeleteIdentity(tt.args.ctx)
		})
	}
}

func TestApp_UpdateIdentity(t *testing.T) {
	type args struct {
		ctx *fiber.Ctx
	}
	tests := []struct {
		name string
		a    *App
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.a.UpdateIdentity(tt.args.ctx)
		})
	}
}
