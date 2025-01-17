package main

import (
	"log"

	"github.com/spf13/cobra"
)

func main() {

	rootCmd := &cobra.Command{Use: "backend"}

	rootCmd.AddCommand(NewMigrationsCommand())
	rootCmd.AddCommand(NewBotCommand())

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
