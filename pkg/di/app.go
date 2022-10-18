package di

import (
	"vocabulary_t/pkg/db"

	"go.uber.org/fx"
)

func AppProviders() []fx.Option {
	modules := []fx.Option{
		db.Module,
	}

	return modules
}
