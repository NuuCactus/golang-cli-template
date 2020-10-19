package cmd

import (
	"github.com/nuucactus/golang-cli-template/pkg/mycommand"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var mycommandCmd = &cobra.Command{
	Use:   "mycommand",
	Short: "mycommand",
	Long:  `mycommand...`,
	Run:   mycommand.RunMyCommand(),
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mycommand.PersistentFlags().String("foo", "", "A help for foo")

	mycommandCmd.Flags().StringP("flag", "f", "default", "Some flag needed by mycommand")
	viper.BindPFlag("flag", mycommandCmd.Flags().Lookup("flag"))

	rootCmd.AddCommand(mycommandCmd)
}
