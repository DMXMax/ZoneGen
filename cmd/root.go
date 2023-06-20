package cmd

import (
	"bytes"
	"fmt"
	zonegen "map/zonemaker"
	"math/rand"
	"os"
	"os/exec"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	terrain          int16
	makeDot, makePNG bool
	fileName         string

	rootCmd = &cobra.Command{
		Use:   "zonegen",
		Short: "zonegen makes maps for an alien invasion game",
		Long:  `Make maps quickly for your alien invasion`,
		Run: func(cmd *cobra.Command, args []string) {
			if viper.GetBool("dot") == false && viper.GetBool("png") == false {
				fmt.Println("You must specify either --dot, --png or both to produce any maps. Run with --help for more information")
				os.Exit(-1)
			}
			gm := zonegen.NewGameMap(int(terrain))
			b := gm.DrawZones()

			if makePNG {
				//look up the dot executable
				dot, err := exec.LookPath("dot")
				if err != nil {
					err = fmt.Errorf("error: %w -- couldn't find the 'dot' executable\nInstall graphviz and make sure it's in your path", err)
					fmt.Println(err)
					os.Exit(-1)

				}

				//send b to the dot executable
				cmd := exec.Command(dot, "-Tpng", "-o", fileName)
				cmd.Stdin = bytes.NewReader(b)
				err = cmd.Run()
				if err != nil {
					log.Fatal().Err(err).Msg("Failed to run dot executable")
				} else {
					fmt.Printf("Wrote %s\n", fileName)
				}

			}
			if makeDot {
				b := gm.DrawZones()
				fmt.Print(string(b))
			}

		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize()
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.SetConfigType("yaml")

	rootCmd.Flags().Int16VarP(&terrain, "terrain", "t", int16(rand.Intn(len(zonegen.AvailablesTerrains))), "Terrain to use. If not specified, a random terrain will be used")
	rootCmd.Flags().BoolVarP(&makeDot, "dot", "d", false, "Generate DOT output to stdout")
	rootCmd.Flags().BoolVarP(&makePNG, "png", "p", true, "Generate PNG file")
	rootCmd.Flags().StringVarP(&fileName, "output", "o", "output.png", "Generate PNG file")

	viper.SetDefault("png", true)
	viper.SetDefault("dot", false)
	viper.SetDefault("output", "map.png")
	viper.SetDefault("fontname", "Arial")
	viper.SetDefault("fontsize", 12)
	viper.BindPFlag("dot", rootCmd.Flags().Lookup("dot"))
	viper.BindPFlag("png", rootCmd.Flags().Lookup("png"))

	zerolog.SetGlobalLevel(zerolog.FatalLevel)

	err := viper.ReadInConfig()

	if err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("No config found. Using defaults.")
		} else {
			fmt.Printf("Something went wrong reading the config file: %v\n", err)
			os.Exit(-1)
		}
	}

}
