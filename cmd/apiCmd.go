package cmd

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cobra"
	"newarea/handler"
	"newarea/pkg/db"
)

func apiCmd() *cobra.Command {
	return &cobra.Command{
		Use: "api",
		RunE: func(cmd *cobra.Command, args []string) error {
			router := gin.Default()

			client, err := db.Connect()
			if err != nil {
				return err
			}
			fmt.Println(client)

			healthHandler := handler.NewHealthHandler()

			router.GET("/health", healthHandler.Health)

			return router.Run(":8080")
		},
	}
}
