package cmd

import (
	"fmt"
	"math/rand"
	"os"
	zonegen "zonegen/zonemaker"

	"github.com/goccy/go-graphviz"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
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
			gm := zonegen.NewGameMap(rand.Intn(7))
			b := gm.DrawZones()

			if makePNG {
				gv := graphviz.New()
				defer gv.Close()
				g, err := graphviz.ParseBytes(b)
				if err != nil {
					log.Fatal().Err(err).Msg("Failed to parse DOT file")
				}
				if err := gv.RenderFilename(g, graphviz.PNG, fileName); err != nil {
					log.Fatal().Err(err).Msg("Failed to render graph")
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
	rootCmd.Flags().Int16VarP(&terrain, "terrain", "t", int16(rand.Intn(7)), "Terrain to use, random if not specified")
	rootCmd.Flags().BoolVarP(&makeDot, "dot", "d", false, "Generate DOT output to stdout")
	rootCmd.Flags().BoolVarP(&makePNG, "png", "p", false, "Generate PNG file")
	rootCmd.Flags().StringVarP(&fileName, "output", "o", "output.png", "Generate PNG file")

}
