/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var infoListCmd = &cobra.Command{
	Use:   "list",
	Short: "Info List",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		connectorOnly, _ := cmd.Flags().GetBool("connector-only")

		if connectorOnly {
			log.Println("fax-connector info list --connector-only called")
			fmt.Println(`{
				"name": "fax Connector Sample Code",
				"description": "PaperCut sample for Fax Connectors",
				"version": "dev-local",
				"capabilities": [
				 {
					"key": "contacts",
					"value": "yes"
				 }
				]
			 }
			 `)
			os.Exit(0)
		}

		// All other info commands need a global settings file

		globalSettingsFileName, _ := cmd.Flags().GetString("global-settings")

		globalSettingsText, error := os.ReadFile(globalSettingsFileName)

		if error != nil {
			log.Println("Can't read global settings file")
			os.Exit(1)
		}

		var globalSettings globalSettingsT

		json.Unmarshal([]byte(globalSettingsText), &globalSettings)

		log.Printf("%v %v", globalSettings.Setting1, globalSettings.Setting2)

		ping, _ := cmd.Flags().GetBool("ping")

		if ping {
			log.Println("fax-connector info list --ping called")
			fmt.Println(`{"reachable": true }`)
			os.Exit(0)
		}

		log.Println("fax connector info list called")
		fmt.Println(`{ "connectorInfo": {
			"name": "fax Connector Sample Code",
			"description": "PaperCut sample for Fax Connectors",
			"version": "dev-local",
			"capabilities": [
			 {
				"key": "contacts",
				"value": "yes"
			 }
			]
		 },
		 "providerInfo": {
			"reachable": true,
			"name": "Sample Fax Provider",
			"version": "beta",
			"server": "",
			"user": ""		
		 }
		 }`)
		os.Exit(0)

	},
}

func init() {
	infoCmd.AddCommand(infoListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly

	infoListCmd.Flags().Bool("connector-only", false, "Only return the Fax Provider name")
	infoListCmd.Flags().Bool("ping", false, "Test connection to Fax Provider service")

}
