package mycommand

import (
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// RunServe the main event loop for the service
func RunMyCommand() func(cmd *cobra.Command, args []string) {
	return func(cmd *cobra.Command, args []string) {

		flag := viper.GetString("flag")

		log.Info().Msgf("Hello World! Flag: %s", flag)

		os.Exit(0)
	}
}
