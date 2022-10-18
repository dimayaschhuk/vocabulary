package main

import (
	"context"

	"vocabulary_t/pkg/di"
	"vocabulary_t/pkg/telegram"

	"github.com/spf13/cobra"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

func NewBotCommand() *cobra.Command {
	migrationsCmd := &cobra.Command{
		Use: "bot",
	}

	migrateCmd := &cobra.Command{
		Use:   "run",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {

			fxOptions := di.AppProviders()
			fxOptions = append(fxOptions,
				fx.Invoke(func(db *gorm.DB) {
					telegram.Run()
				}))

			app := fx.New(fxOptions...)

			return app.Start(context.Background())
		},
	}

	sendCmd := &cobra.Command{
		Use:   "send",
		Short: "Execute migrations",
		RunE: func(cmd *cobra.Command, args []string) error {

			fxOptions := di.AppProviders()
			fxOptions = append(fxOptions,
				fx.Invoke(func(db *gorm.DB) {
					telegram.Send()
				}))

			app := fx.New(fxOptions...)

			return app.Start(context.Background())
		},
	}

	migrationsCmd.AddCommand(migrateCmd)
	migrationsCmd.AddCommand(sendCmd)

	return migrationsCmd
}
