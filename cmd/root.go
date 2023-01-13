/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"os"

	"github.com/chaewonkong/go-template/app"
	"github.com/spf13/cobra"
	"go.uber.org/fx"
)

const version = "0.0.1"

func run(cmd *cobra.Command, args []string) {
	// export PORT to environ
	if Port != "" {
		os.Setenv("PORT", Port)
	}
	fx.New(app.Modules).Run()
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
		log.Fatal(err)
	}
}

var Port string

func init() {
	// add port flag
	rootCmd.PersistentFlags().StringVarP(&Port, "port", "p", "", "--port 8080 or -p 8080")
}
