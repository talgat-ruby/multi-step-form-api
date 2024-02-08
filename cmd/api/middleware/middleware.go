package middleware

import (
	"github.com/talgat-ruby/multi-step-form-api/cmd/api/types"
	dbT "github.com/talgat-ruby/multi-step-form-api/cmd/db/types"
)

type middlewareObject struct {
	api types.Api
	db  dbT.DB
}

func New(api types.Api, db dbT.DB) types.Middleware {
	return &middlewareObject{
		api: api,
		db:  db,
	}
}
