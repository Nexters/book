/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/rs/zerolog/log"

	"os"

	"github.com/nexters/book/http"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const version = "0.0.1"

func run(cmd *cobra.Command, args []string) {
	// export PORT to environ
	if Port != "" {
		err := os.Setenv("PORT", Port)
		if err != nil {
			log.Fatal().Err(err)
		}
	}
	fx.New(http.Modules).Run()
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "app",
	Short: "A brief description of your application",
	Run:   run,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		log.Fatal().Err(err)
	}
}

var Port string

func init() {
	// add port flag
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "", "--port 8080 or -p 8080")
}
