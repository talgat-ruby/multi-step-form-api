package db

import (
	"log/slog"

	"github.com/talgat-ruby/multi-step-form-api/cmd/db/model"
	"github.com/talgat-ruby/multi-step-form-api/cmd/db/types"
	"github.com/talgat-ruby/multi-step-form-api/configs"
)

func New(log *slog.Logger, conf *configs.DBConfig) (types.DB, error) {
	return model.New(log, conf)
}
