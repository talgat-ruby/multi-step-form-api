package router

import (
	"log/slog"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"

	"github.com/talgat-ruby/multi-step-form-api/cmd/api/handler/v1/forms"
	dbT "github.com/talgat-ruby/multi-step-form-api/cmd/db/types"
)

func v1Group(api *echo.Group, db dbT.DB, v *validator.Validate, l *slog.Logger) {
	g := api.Group("/v1")

	v1formsRouter(g, db, v, l)
}

func v1formsRouter(v1 *echo.Group, db dbT.DB, v *validator.Validate, l *slog.Logger) {
	h := forms.New(db, v, l)

	v1.POST("/forms", h.Add)
}
