/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"log"
	"net/http"
	"time"

	"github.com/brianseitel/gleeman/internal/builder"
	"github.com/brianseitel/gleeman/internal/watcher"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

// serveCmd represents the serve command
var serveCmd = &cobra.Command{
	Use:   "serve",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fs := http.FileServer(http.Dir("./public"))
		http.Handle("/", fs)

		handler := func() error {
			logger, _ := zap.NewDevelopment()
			b := builder.Builder{
				Logger: logger,
			}

			return b.Start(map[string]string{
				"name":      viper.GetString("name"),
				"copyright": time.Now().Format("2006"),
				"now":       time.Now().Format(time.RFC3339),
			})
		}

		wch := watcher.New([]string{"./tales/entries", "./tales/layout"}, handler)
		go wch.Start()

		log.Println("Listening on :3000...")
		err := http.ListenAndServe(":3000", nil)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func init() {
	rootCmd.AddCommand(serveCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serveCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serveCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
