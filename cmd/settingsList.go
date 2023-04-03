/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var settingsListCmd = &cobra.Command{
	Use:   "list",
	Short: "Settings List",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		log.Println("fax-connector settings list called")

		locale, _ := cmd.Flags().GetString("locale")

		settingsConfigList := map[string]string{
			"en-US": `[
				{
					"id": "test1",
					"label": "An enum value (not used)",
					"type": "enum",
					"values": ["Option 1", "Option 2"],
					"defaultValue": "Option 2",
					"isOptional": true
				},
				{
					"id": "destPath",
					"label": "Destination fax path (must be local)",
					"type": "text",
					"values": null,
					"defaultValue": "",
					"isOptional": false
				}
			]`,
			"en-AU": `[
				{
				"id": "test1",
				"label": "Strewth mate! Enum value not used, no problem",
				"type": "enum",
				"values": ["Option 1", "Option 2"],
				"defaultValue": "Option 2",
				"isOptional": true
			},
			{
				"id": "destPath",
				"label": "where the fax goes (must be local)",
				"type": "text",
				"values": null,
				"defaultValue": "",
				"isOptional": false
			}
			]`}

		fmt.Println(settingsConfigList[locale])
	},
}

func init() {
	settingsCmd.AddCommand(settingsListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	settingsListCmd.Flags().String("locale", "en-US", "Locale for text strings (Unsupported)")
	// Default locale is assumed to be US English

}
