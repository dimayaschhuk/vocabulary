package main

import (
	"context"

	"vocabulary_t/migrations"
	"vocabulary_t/pkg/di"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewMigrationsCommand() *cobra.Command {
	migrationsCmd := &cobra.Command{
		Use: "migrations",
	}

	migrateCmd := &cobra.Command{
		Use:   "migrate",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {

			fxOptions := di.AppProviders()
			fxOptions = append(fxOptions,
				fx.Invoke(func(db *gorm.DB) {
					if err := migrations.Create(db); err != nil {
						println(err.Error())
					}
				}))

			app := fx.New(fxOptions...)

			return app.Start(context.Background())
		},
	}

	createMigrationCmd := &cobra.Command{
		Use:   "create",
		Short: "Create new migration",
		Args:  cobra.MinimumNArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			return nil
		},
	}

	migrationsCmd.AddCommand(migrateCmd)
	migrationsCmd.AddCommand(createMigrationCmd)

	return migrationsCmd
}
