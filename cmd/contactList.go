/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// contactListCmd represents the contactList command
var contactListCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var globalSettingsText []byte
		var loggedInUserDetailsText []byte
		var err error

		globalSettingsFileName, _ := cmd.Flags().GetString("global-settings")
		loggedInUserDetailsFileName, _ := cmd.Flags().GetString("contact-settings")

		if globalSettingsText, err = os.ReadFile(globalSettingsFileName); err != nil {
			log.Fatalf("Could not read global settings file %v", globalSettingsFileName)
		}

		var globalSettings globalSettingsT

		if nil != json.Unmarshal([]byte(globalSettingsText), &globalSettings) {
			log.Fatalf("Could not unmarshal global settings json %v", globalSettingsText)
		}

		log.Printf("%v %v", globalSettings.Setting1, globalSettings.Setting2)

		if loggedInUserDetailsText, err = os.ReadFile(loggedInUserDetailsFileName); err != nil {
			log.Fatalf("Could not read job settings file %v", loggedInUserDetailsFileName)
		}

		var loggedInUserDetails loggedInUserDetailsT

		if nil != json.Unmarshal([]byte(loggedInUserDetailsText), &loggedInUserDetails) {
			log.Fatalf("Could not unmarshal logged-in user json %v", loggedInUserDetailsText)
		}

		log.Printf("fax-connector contact list called, logged in user: %v", loggedInUserDetails.UserEmailAddress)

	},
}

func init() {
	contactCmd.AddCommand(contactListCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// contactListCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// contactListCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	contactListCmd.Flags().String("contact-settings", "", "Name of the logged-in user details file")
}
