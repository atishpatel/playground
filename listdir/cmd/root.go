/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "listdir",
	Short: "listdir displays a human readable ls with dir depth",
	Long:  `listdir displays a readable ls with dir depth and sizes`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		err := cmd.ParseFlags(args)
		if err != nil {
			fmt.Printf("Invalid arguments: %v", err)
			return
		}
		// get size and depth
		size, err := cmd.Flags().GetBool("size")
		if err != nil {
			fmt.Printf("Failed to get size: %v", err)
			return
		}
		depth, err := cmd.Flags().GetInt("depth")
		if err != nil {
			fmt.Printf("Failed to get depth: %v", err)
			return
		}
		fmt.Printf("Running command with size(%t) and depth(%d)\n", size, depth)
		// walking filepaths

		type filestruct struct {
			Name  string
			Dirs  []filestruct
			Files []string
			Depth int
		}

		abs, err := filepath.Abs(".")
		if err != nil {
			fmt.Printf("Failed to get filepath.Abs: %v", err)
			return
		}
		fmt.Printf("%s\n", filepath.Base(abs))

		paths := []string{}
		err = filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.Name() == "." {
				abs, err := filepath.Abs(".")
				if err != nil {
					fmt.Printf("Failed to get filepath.Abs: %v", err)
					return nil
				}
				fmt.Printf("%s\n", filepath.Base(abs))
				return nil
			}
			if size {
				fmt.Printf("├── %s (%d)\n", path, info.Size())
			} else {
				fmt.Printf("├── %s\n", path)
			}
			paths = append(paths, path)
			return nil
		})
		if err != nil {
			fmt.Printf("Failed to filepath.Walk: %v", err)
			return
		}

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.listdir.yaml)")

	rootCmd.Flags().Int("depth", 2, "Specify how deep to treverse dirs in displaying")
	rootCmd.Flags().BoolP("size", "s", false, "Show size for files")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".listdir" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".listdir")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
