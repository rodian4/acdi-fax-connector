/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {

		var globalSettingsText []byte
		var jobSettingsText []byte
		var err error

		globalSettingsFileName, _ := cmd.Flags().GetString("global-settings")
		jobSettingsFileName, _ := cmd.Flags().GetString("job-settings")

		if globalSettingsText, err = os.ReadFile(globalSettingsFileName); err != nil {
			log.Fatalf("Could not read global settings file %v", globalSettingsFileName)
		}

		var globalSettings globalSettingsT

		if nil != json.Unmarshal([]byte(globalSettingsText), &globalSettings) {
			log.Fatalf("Could not unmarshal global settings json %v", globalSettingsText)
		}

		log.Printf("%v %v", globalSettings.Setting1, globalSettings.Setting2)

		if jobSettingsText, err = os.ReadFile(jobSettingsFileName); err != nil {
			log.Fatalf("Could not read job settings file %v", jobSettingsFileName)
		}

		var jobSettings jobSettingsT

		if nil != json.Unmarshal([]byte(jobSettingsText), &jobSettings) {
			log.Fatalf("Could not unmarshal jobs settings json %v", jobSettingsText)
		}

		log.Println("fax-connector fax create called")

		var faxResponses []faxResponseT

		for _, destination := range jobSettings.Destinations {
			faxResponses = append(faxResponses, processFax(destination, jobSettings.Documents))
		}

		j, _ := json.MarshalIndent(faxResponses, "", "  ")
		fmt.Print(string(j))
	},
}

func init() {
	faxCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports local flags which will only run when this command
	// is called directly,

	createCmd.Flags().String("job-settings", "", "Name of Job Settings file")

}

func processFax(dest destinationsT, docs []faxDocumentT) (response faxResponseT) {

	for _, doc := range docs {
		response = cpFax(doc.Path, dest.To+doc.Name)
		if response.Success {
			return response
		}
	}
	return response
}

func cpFax(src, dest string) faxResponseT {
	// Open original file
	destDir := "/tmp/"
	original, err := os.Open(src)
	if err != nil {
		return faxResponseT{false, "NOT OK", fmt.Sprintf("Fax %v NOT sent. Could not open source fax", src)}
	}
	defer original.Close()

	// Create new file
	new, err := os.Create(destDir + dest)
	if err != nil {
		return faxResponseT{false, "NOT OK", fmt.Sprintf("Fax %v NOT sent. Could not open dest %v", src, dest)}
	}
	defer new.Close()

	//This will copy
	bytesWritten, err := io.Copy(new, original)
	if err != nil {
		return faxResponseT{false, "NOT OK", fmt.Sprintf("Fax %v NOT sent to %v", src, dest)}
	}
	return faxResponseT{true, "OK", fmt.Sprintf("Fax of length %d sent", bytesWritten)}
}
