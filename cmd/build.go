/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"time"

	"github.com/brianseitel/gleeman/internal/builder"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Builds the HTML output based on entries in the tales/entries directory.",
	Long:  `This command loops through each of the entires in your tales/entries directory and converts them to HTML, fills in dynamic fields, and places the output in the public/views directory.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()
		b := builder.Builder{
			Logger: logger,
		}

		err := b.Start(map[string]string{
			"name":      viper.GetString("name"),
			"copyright": time.Now().Format("2006"),
			"now":       time.Now().Format(time.RFC3339),
		})
		if err != nil {
			panic(err)
		}

		logger.Info("Done.")
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	viper.SetConfigName("settings") // name of config file (without extension)
	viper.SetConfigType("yaml")     // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath("./tales")  // path to look for the config file in
	viper.AddConfigPath(".")        // optionally look for config in the working directory
	err := viper.ReadInConfig()     // Find and read the config file
	if err != nil {                 // Handle errors reading the config file
		log.Println("Error loading settings file. Please run `gleeman init`.")
	}
}
