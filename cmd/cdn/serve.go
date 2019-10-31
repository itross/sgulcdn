package cdn

import (
	"github.com/itross/sgulcdn/internal"
	"github.com/spf13/cobra"
)

var serveCommand = &cobra.Command{
	Use:   "serve",
	Short: "starts the service",
	Long:  "This command starts the CDN microservice to serve requests",
	Run: func(cmd *cobra.Command, args []string) {
		serve(args)
	},
}

func init() {
	RootCmd.AddCommand(serveCommand)
}

func serve(args []string) {
	app := internal.New()

	logger.Info("starting CDN Service")
	app.RunAndWait()
}
