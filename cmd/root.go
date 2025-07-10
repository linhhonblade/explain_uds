package cmd

import (
	"context"
	"database/sql"
	uds "explain_uds/common"
	"github.com/spf13/cobra"
	"log"
)

var rootCmd = &cobra.Command{
	Use:   "explainuds",
	Short: "Explain UDS (Unix Domain Sockets) in Go",
	Run: func(cmd *cobra.Command, args []string) {
		// This is where you would implement the functionality to explain UDS.
		// For now, we can just print a message.
		// Init db connection
		// TODO: Abstract the database connection logic
		db, err := sql.Open("sqlite3", "./sqlite3/uds.db")
		if err != nil {
			log.Fatalf("cannot open DB: %v", err)
		}
		defer db.Close()

		// Store db in context
		ctx := context.Background()
		ctx = context.WithValue(ctx, uds.CtxKeyDB{}, db)

		println("Explain UDS (Unified Diagnostic Service) in Go.")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		println("Error executing command:", err.Error())
	}
}
