package scripts

type point struct {
	x, y int
}

// returns the six adjacent coordinates on a standard keyboard, where each row is slanted to the
// right from the last. adjacencies are clockwise, starting with key to the left, then two keys
// above, then right key, then two keys below. (that is, only near-diagonal keys are adjacent,
// so g's coordinate is adjacent to those of t,y,b,v, but not those of r,u,n,c.)
func getSlantedAdjacencyCoords(p point) []point {
	return []point{
		{p.x - 1, p.y}, {p.x, p.y - 1}, {p.x + 1, p.y - 1},
		{p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y + 1},
	}
}

// returns the nine clockwise adjacent coordinates on a keypad, where each row is vert aligned.
func getAlignedAdjacentCoords(p point) []point {
	return []point{
		{p.x - 1, p.y}, {p.x - 1, p.y - 1}, {p.x, p.y - 1}, {p.x + 1, p.y - 1},
		{p.x + 1, p.y}, {p.x + 1, p.y + 1}, {p.x, p.y + 1}, {p.x - 1, p.y + 1},
	}
}

// builds an adjacency graph as a map.
// adjacent characters occur in a clockwise order.
// for example:
// * on qwerty layout, 'g' maps to ['fF', 'tT', 'yY', 'hH', 'bB', 'vV']
// * on keypad layout, '7' maps to [None, None, None, '=', '8', '5', '4', None]
func buildGraph(layout string, slanted bool) map[rune][]rune {
	table := map[rune][]rune{}
	adjFn := getAlignedAdjacentCoords
	if slanted {
		adjFn = getSlantedAdjacencyCoords
	}

}
