package server

import "testing"

func TestSetupRoutes(t *testing.T) {
	type args struct {
		app *App
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SetupRoutes(tt.args.app)
		})
	}
}
