package scripts

type point struct {
	x, y int
}

// returns the six adjacent coordinates on a standard keyboard, where each row is slanted to the
// right from the last. adjacencies are clockwise, starting with key to the left, then two keys
// above, then right key, then two keys below. (that is, only near-diagonal keys are adjacent,
// so g's coordinate is adjacent to those of t,y,b,v, but not those of r,u,n,c.)
func getSlantedAdjacencyCoords(p point) []point {
	return []point{{p.x - 1, p.y}, {p.x, p.y - 1}, {p.x + 1, p.y - 1}, {p.x + 1, p.y}, {p.x, p.y + 1}, {p.x - 1, p.y + 1}}
}
