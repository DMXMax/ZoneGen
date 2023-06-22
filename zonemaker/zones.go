package zonemaker

import (
	"bytes"
	"math/rand"

	"fmt"

	"github.com/dominikbraun/graph"
	"github.com/dominikbraun/graph/draw"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type GameMap struct {
	Terrain   TerrainDescription
	gameZones *graph.Graph[int, int]
}

func NewGameMap(terrain int) *GameMap {
	var gm GameMap
	gm.Terrain = AvailablesTerrains[terrain]
	gm.gameZones = gm.makeZones(len(AvailablesTerrains[terrain].Areas))
	return &gm
}

func (gm *GameMap) makeZones(zones int) *graph.Graph[int, int] {

	g := graph.New(graph.IntHash)

	g.AddVertex(0, graph.VertexAttribute("label", gm.Terrain.Areas[0]),
		graph.VertexAttribute("shape", "rectangle"),
		graph.VertexAttribute("style", "bold"),
		graph.VertexAttribute("fontname", viper.GetString("fontname")),
		graph.VertexAttribute("fontsize", viper.GetString("fontsize")))

	for i := 1; i < zones; i++ {
		g.AddVertex(i,
			graph.VertexAttribute("shape", "rectangle"),
			graph.VertexAttribute("fontname", viper.GetString("fontname")),
			graph.VertexAttribute("fontsize", viper.GetString("fontsize")),
			graph.VertexAttribute("label", fmt.Sprintf("%d: %s", i, gm.Terrain.Areas[i])),
		)
		g.AddEdge(i, rand.Intn(i))

	}

	//lets add some edges
	odds := 2
	for cnt := 0; cnt < 4; cnt++ {
		if rand.Int()%odds == 0 {
			x, y := rand.Intn(zones), rand.Intn(zones)
			if x < y {
				x, y = y, x
			}
			for x == y || func() bool {
				_, err := g.Edge(x, y)
				return err == nil
			}() {
				y = rand.Intn(zones)
			}

			log.Info().Msgf("Adding edge %d -> %d", x, y)
			g.AddEdge(x, y)
			odds++

		} else {
			break
		}
	}

	return &g
}

func (gm *GameMap) DrawZones() []byte {
	buf := bytes.NewBuffer(make([]byte, 0, 1024))

	if err := draw.DOT(*gm.gameZones, buf,
		draw.GraphAttribute("fontname", viper.GetString("titlefontname")),
		draw.GraphAttribute("fontsize", viper.GetString("titlefontsize")),
		draw.GraphAttribute("label", gm.Terrain.Description),
		draw.GraphAttribute("labelloc", "t"),
	); err != nil {
		log.Fatal().Err(err).Msg("Failed to create DOT file")
	}

	return buf.Bytes()

}
