package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/ethancox127/size/utils"
)

// Variables for storing the configuration file and the flags if set by the user (Human and Recursive).
var cfgFile string
var Human bool 
var Recursive bool

// The root command for cobra.
var rootCmd = &cobra.Command{
	Use:   "size",
	Short: "Size outputs the size of all listed directories.",
	Long:  "Size outputs the size of all listed directories.  Users can specify for subdirectories to be read from based on the given directories and can specify for the output to be more readable.",
	Run: func(cmd *cobra.Command, args []string) {

		// Check that at least one directory has been specified.
		if len(args) == 0 {
			fmt.Println("Please specify at least one directory.")
			return
		}

		// Variables for tracking the current directory (and subdirectories if specified) size and the cumulative size of all the listed directories.
		var cumulativeSize int64 = 0
		var dirSize int64 = 0
		var err error
		
		// Loop through the directories given by the user.
		for _, dir := range args {

			fmt.Println()

			// If the user specifies the recursive functionality, call the recursive dir size function, otherwise call the normal dir size function.
			if Recursive {
				dirSize = 0
				dirSize, err = utils.RecursiveDirSize(utils.AbsPath(dir), Human, &dirSize)
			} else {
				dirSize, err = utils.DirSize(utils.AbsPath(dir), Human)
			}

			// Check for errors and update the cumulative size.
			if err != nil {
				fmt.Println(err)
			}
			cumulativeSize += dirSize
		}

		// Print out the cumulative directory sizes.
		fmt.Print("Cumulative directory size: ")
		if Human {
			utils.PrettyPrint(cumulativeSize)
		} else {
			fmt.Print(cumulativeSize)
			fmt.Println(" bytes")
		}
		fmt.Println()
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
	cobra.OnInitialize(initConfig)

	// Set the flags to accept them from the user.
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.cmd_line.yaml)")
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	rootCmd.Flags().BoolVarP(&Human, "human", "a", false, "Human Readable output")
	rootCmd.Flags().BoolVarP(&Recursive, "recursive", "r", false, "Recursive directory search")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".cmd_line" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigType("yaml")
		viper.SetConfigName(".cmd_line")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
	}
}
