package zonemaker

type TerrainDescription struct {
	Description string
	Areas       []string
	MapSize     string
}

var AvailablesTerrains = []TerrainDescription{
	{"Basic Terrain", []string{"LZ", "Alpha", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot"}, "4"},
	{"Antarctic Research Base\nWe picked up a signal from a station doing ice core research in the Antarctic",
		[]string{
			"LZ\nThe Ice ShelfOoutside the Station",
			"Drilling Station",
			"Crevasse",
			"Sleeping Quarters",
			"Lab",
			"Motor Pool",
			"Frozen Wastes"}, "6"},
	{"Urban Sprawl\nAliens have struck in the heart of a major city!",
		[]string{
			"LZ\nA Parking Lot Outside a Mall",
			"Gas Station",
			"Highway Off-ramp",
			"Commercial\nHigh-Rise",
			"Park",
			"Waterfront",
			"Subway Station"}, "6"},
	{"Secluded Farm\nThere have been reports of cattle mutilations in the area even before they made their presence known",
		[]string{
			"LZ\nThe Dirt Road Leading up to a Farm",
			"A Field of Corn Stalks",
			"An Orchard",
			"A Barn",
			"The Farmhouse",
			"A Dilapidated Shed",
			"A Small UFO Landed in a Field"}, "6"},
	{"Suburban Hellscape\nThere was a major battle here, and the aliens won.",
		[]string{
			"LZ\nJunior Leage Baseball Field",
			"A School,\nBurnt to the Ground",
			"A cul-de-sac,\nwrecked tanks",
			"A Fast Food Resteraunt\nwith a Caved-in Roof",
			"Rows of Houses",
			"A crashed Military\nHelicopter",
			"A Huge Crater where\na Mall Once Stood"}, "6"},
	{"Alien Base\nIntelligence has found an underground alien base.\nWe must strike them where it hurts!",
		[]string{
			"LZ\nA Cargo Elevator Leading Down",
			"Labyrinthine Corridors",
			"Alien Biology Lab",
			"Garden of Alien Plants",
			"Power Core",
			"Alien Stasis Room",
			"The Command Center"}, "6"},
}
