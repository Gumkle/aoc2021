package day

/*
MapSize indicates size of a side of a map, which is a square
*/
const (
	MapSize = 1000
	up      = -1
	down    = 1
	left    = -1
	right   = 1
)

/*
WorldMap Type for describing world map
*/
type WorldMap [MapSize][MapSize]uint

/*
AreCoordsPerpendicularLine takes coords of two points and examines if they are horizontal or vertical
*/
func AreCoordsPerpendicularLine(x1, x2, y1, y2 uint) bool {
	return x1 == x2 || y1 == y2
}

/*
AreCoordsDiagonalLine takes coords of two points and examines if they are diagonal, 45 degree
*/
func AreCoordsDiagonalLine(x1, x2, y1, y2 uint) bool {
	diffX := int(x1) - int(x2)
	diffY := int(y1) - int(y2)
	return abs(diffX) == abs(diffY)
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

/*
NewMap returns new map
*/
func NewMap() *WorldMap {
	return new(WorldMap)
}

/*
MarkPerpendicularLine marks perpendicular line on world map from given coords
*/
func (m *WorldMap) MarkPerpendicularLine(x1, x2, y1, y2 uint) {
	var smaller, bigger uint
	if x1 == x2 {
		if y1 > y2 {
			smaller, bigger = y2, y1
		} else {
			smaller, bigger = y1, y2
		}
		for i := smaller; i <= bigger; i++ {
			m[x1][i]++
		}
	}
	if y1 == y2 {
		if x1 > x2 {
			smaller, bigger = x2, x1
		} else {
			smaller, bigger = x1, x2
		}
		for i := smaller; i <= bigger; i++ {
			m[i][y1]++
		}
	}
}

/*
MarkDiagonalLine marks diagonal line on world map from given coords
*/
func (m *WorldMap) MarkDiagonalLine(x1, x2, y1, y2 uint) {
	var vert, horz int
	if y1 > y2 {
		vert = up
	} else {
		vert = down
	}
	if x1 > x2 {
		horz = left
	} else {
		horz = right
	}
	distance := abs(int(x1) - int(x2))
	for i := 0; i <= distance; i++ {
		m[int(x1)+i*horz][int(y1)+i*vert]++
	}
}
