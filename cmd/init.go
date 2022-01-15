/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"os"

	"github.com/brianseitel/gleeman/internal/initializer"
	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		logger, _ := zap.NewDevelopment()

		dirs := []string{
			"tales",
			"tales/entries/",
			"tales/layout",
			"public/",
			"public/assets/",
		}

		for _, dir := range dirs {
			logger.Sugar().Infof("Creating directory %s", dir)
			os.Mkdir(dir, 0777)
		}

		files := []string{
			"tales/settings.yaml",
			"tales/layout/_layout.html",
			"tales/layout/_entry.html",
			"tales/layout/_main.html",
			"public/assets/main.css",
		}

		for _, file := range files {
			logger.Sugar().Infof("Creating file %s", file)
			data := initializer.Fetch(file)
			f, err := os.Create(file)
			if err != nil {
				panic(err)
			}
			f.Write(data)
			f.Close()
		}

		logger.Sugar().Info("Done")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
