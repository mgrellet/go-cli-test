/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"github.com/mgrellet/cli/internal/constants"
	"github.com/mgrellet/cli/internal/structs"
	"github.com/spf13/cobra"
	"net/http"
	"os"
	"strconv"
)

// envCmd represents the env command
var envCmd = &cobra.Command{
	Use:   "env",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		if flags, err := cmd.Flags().GetBool("execute"); err == nil {
			fmt.Println("execute flag:", flags)
			compareRequests()

		}
		fmt.Println("env called")
	},
}

func init() {
	envCmd.Flags().BoolP("execute", "x", false, "Execute flag")
	compareCmd.AddCommand(envCmd)
}

func compareRequests() {
	fmt.Println("compare called")

	csvFile, err := os.Open("data.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	records, err := reader.ReadAll()
	if err != nil {
		fmt.Println("Error reading CSV file:", err)
		return
	}

	for i, record := range records {
		if i == 0 {
			continue
		}
		requestPayload := structs.Request{
			Name:       "data " + record[0],
			Percentage: parseStrToFloat(record[0]),
		}

		requestBody, err := json.Marshal(requestPayload)
		if err != nil {
			fmt.Println("Error marshalling request payload:", err)
			return
		}

		response1, err := http.Post(constants.URL_API_ONE, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			return
		}
		defer response1.Body.Close()

		var responseOne structs.ResponseOne
		err = json.NewDecoder(response1.Body).Decode(&responseOne)
		if err != nil {
			fmt.Println("Error decoding response body:", err)
			return
		}

		fmt.Println("Response 1 Rate: ", responseOne.Rate)

		response2, err := http.Post(constants.URL_API_TWO, "application/json", bytes.NewBuffer(requestBody))
		if err != nil {
			fmt.Println("Error making HTTP request:", err)
			return
		}
		defer response2.Body.Close()
		var responseTwo structs.ResponseOne
		err = json.NewDecoder(response2.Body).Decode(&responseTwo)
		if err != nil {
			fmt.Println("Error decoding response body:", err)
			return
		}

		fmt.Println("Response 2 Rate: ", responseTwo.Rate)
	}
}

func parseStrToFloat(str string) float64 {
	f, err := strconv.ParseFloat(str, 64)
	if err != nil {
		fmt.Println(err)
	}
	return f
}
