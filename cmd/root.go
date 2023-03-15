/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"robot/pkg/robot"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "chatgpt",
	Short: "chatgpt机器人",
	Long:  ``,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		configFile, _ := cmd.Flags().GetString("config")
		port := 11443
		if configFile == "" {
			log.Panic("config file is empty")
		} else {
			f, err := os.Open(configFile)
			if err != nil {
				log.Fatalln("failed to open config file:", configFile, err)
			}
			defer f.Close()
			content, err := ioutil.ReadAll(f)
			if err != nil {
				log.Fatalln("failed to read config file:", configFile, err)
			}
			var configList []robot.Config
			err = json.Unmarshal(content, &configList)
			if err != nil {
				log.Fatalln("failed to parse config file:", configFile, err)
			}
			robot.RunServer(port, configList)

		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringP("config", "c", "", "config file")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
