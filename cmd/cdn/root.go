package cdn

import (
	"fmt"
	"os"
	"strings"

	"github.com/itross/sgul"

	"go.uber.org/zap"

	"github.com/spf13/cobra"
)

var logger *zap.SugaredLogger

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "cdn",
	Short: "Sgul CDN Service",
	Long: `---------------------------------------------
Sgul CDN Service v. 0.1.0
---------------------------------------------`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initialize)
}

func initialize() {
	logger = sgul.GetLogger()
	env := strings.ToLower(os.Getenv("ENV"))
	if env == "" {
		env = "dev"
		logger.Warn("no envirnoment specified, fall back to default")
	}
	logger.Infow("initialization", "environment", env)
}
