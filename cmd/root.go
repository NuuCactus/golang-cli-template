package cmd

import (
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var cfgFile string
var verbose bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golang-cli-template",
	Short: "Template cli",
	Long:  `Template cli...`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initLogging)
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "./.golang-cli-template.yaml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "enable verbose output")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
		viper.AutomaticEnv() // read in environment variables that match
		// OBS without this prefix sys variables like USER and USERNAME will be used with "user" flag which probably which gives headaches
		viper.SetEnvPrefix("ENV")
		viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

		// If a config file is found, read it in.
		if err := viper.ReadInConfig(); err == nil {
			log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
		}
	}
}

func initLogging() {

	if verbose {
		log.Logger = log.With().Caller().Logger()
	}

	log.Logger = log.Output(zerolog.ConsoleWriter{
		Out:        os.Stderr,
		TimeFormat: "15:04:05.999",
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	zerolog.TimestampFieldName = "time"
	zerolog.LevelFieldName = "level"
	zerolog.MessageFieldName = "msg"

	zerolog.SetGlobalLevel(zerolog.InfoLevel)
	if verbose {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	}
}
