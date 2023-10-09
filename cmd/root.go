package cmd

import (
	"github.com/spf13/cobra"
	"log"
	"template/cmd/migration"
	"template/config"
)

func Start() {
	cfg := config.ReadConfig()
	// root command
	root := &cobra.Command{}

	// command allowed
	cmds := []*cobra.Command{
		{
			Use:   "db:migrate",
			Short: "database migration",
			Run: func(cmd *cobra.Command, args []string) {
				migration.RunMigration(cfg)
			},
		},
		{
			Use:   "api",
			Short: "run api server",
			Run: func(cmd *cobra.Command, args []string) {
				migration.StartServer(cfg)
			},
		},
	}
	root.AddCommand(cmds...)
	if err := root.Execute(); err != nil {
		log.Fatal(err)
	}
}
