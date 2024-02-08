package types

import (
	"context"

	"github.com/talgat-ruby/multi-step-form-api/cmd/db/model"
)

type DB interface {
	AddForm(ctx context.Context, input model.Form) error
}
