package main

import "fmt"

type cell struct {
	x, y, z, w int
}

type mapa struct {
	cells []cell
}

func main() {

	var origin mapa

	origin.cells = []cell{
		cell{x: 3, y: 0, z: 0, w: 0},
		cell{x: 5, y: 0, z: 0, w: 0},
		cell{x: 7, y: 0, z: 0, w: 0},
		cell{x: 2, y: 1, z: 0, w: 0},
		cell{x: 5, y: 1, z: 0, w: 0},

		cell{x: 0, y: 2, z: 0, w: 0},
		cell{x: 2, y: 2, z: 0, w: 0},
		cell{x: 4, y: 2, z: 0, w: 0},
		cell{x: 5, y: 2, z: 0, w: 0},
		cell{x: 7, y: 2, z: 0, w: 0},

		cell{x: 0, y: 3, z: 0, w: 0},
		cell{x: 1, y: 3, z: 0, w: 0},
		cell{x: 2, y: 3, z: 0, w: 0},
		cell{x: 4, y: 3, z: 0, w: 0},
		cell{x: 5, y: 3, z: 0, w: 0},

		cell{x: 0, y: 4, z: 0, w: 0},
		cell{x: 1, y: 4, z: 0, w: 0},
		cell{x: 2, y: 4, z: 0, w: 0},
		cell{x: 3, y: 4, z: 0, w: 0},
		cell{x: 4, y: 4, z: 0, w: 0},
		cell{x: 6, y: 4, z: 0, w: 0},
		cell{x: 7, y: 4, z: 0, w: 0},

		cell{x: 0, y: 5, z: 0, w: 0},

		cell{x: 0, y: 6, z: 0, w: 0},
		cell{x: 3, y: 6, z: 0, w: 0},
		cell{x: 6, y: 6, z: 0, w: 0},
		cell{x: 7, y: 6, z: 0, w: 0},

		cell{x: 3, y: 7, z: 0, w: 0},
		cell{x: 4, y: 7, z: 0, w: 0},
		cell{x: 6, y: 7, z: 0, w: 0},
		cell{x: 7, y: 7, z: 0, w: 0},
	}

	for i := 0; i < 6; i++ {
		origin.iterate()
	}

	fmt.Printf("size: %d\n", len(origin.cells))

}

func (m *mapa) size() int {

	var max, min int = 0, 0

	for _, cell := range m.cells {
		if cell.x < min {
			min = cell.x
		}
		if cell.x > max {
			max = cell.x
		}
	}

	return abs(max) + abs(min) + 1
}

func (m *mapa) mins() (int, int, int, int) {
	x, y, z, w := 0, 0, 0, 0

	for _, cell := range m.cells {
		if cell.x < x {
			x = cell.x
		}
		if cell.y < y {
			y = cell.y
		}
		if cell.z < z {
			z = cell.z
		}
		if cell.w < w {
			w = cell.w
		}
	}
	return x, y, z, w
}

func (m *mapa) maxs() (int, int, int, int) {
	x, y, z, w := 0, 0, 0, 0

	for _, cell := range m.cells {
		if cell.x > x {
			x = cell.x
		}
		if cell.y > y {
			y = cell.y
		}
		if cell.z > z {
			z = cell.z
		}
		if cell.w > w {
			w = cell.w
		}
	}
	return x, y, z, w
}

func (m *mapa) iterate() {
	maxX, maxY, maxZ, maxW := m.maxs()
	minX, minY, minZ, minW := m.mins()

	cells := []cell{}

	for w := minW - 1; w <= maxW+1; w++ {
		for z := minZ - 1; z <= maxZ+1; z++ {
			for y := minY - 1; y <= maxY+1; y++ {
				for x := minX - 1; x <= maxX+1; x++ {
					//fmt.Printf("x: %d y:%d z:%d - val: %d\n", x, y, z, m.get(x, y, z))
					active := m.get(x, y, z, w)
					neigh := m.activeN(x, y, z, w)
					if active == 1 {
						if neigh == 2 || neigh == 3 {
							cells = append(cells, cell{x: x, y: y, z: z, w: w})
						}
					} else {
						if neigh == 3 {
							cells = append(cells, cell{x: x, y: y, z: z, w: w})
						}
					}

				}
			}
		}
	}
	m.cells = cells
}

func (m *mapa) activeN(xo, yo, zo, wo int) int {
	count := 0

	for w := wo - 1; w <= wo+1; w++ {
		for z := zo - 1; z <= zo+1; z++ {
			for y := yo - 1; y <= yo+1; y++ {
				for x := xo - 1; x <= xo+1; x++ {
					if x == xo && y == yo && z == zo && w == wo {
						continue
					}
					if m.get(x, y, z, w) == 1 {
						count++
					}
				}
			}
		}
	}
	return count
}

func (m *mapa) get(x, y, z, w int) int {

	for _, cell := range m.cells {
		if cell.x == x && cell.y == y && cell.z == z && cell.w == w {
			return 1
		}
	}
	return 0
}

func (m *mapa) draw(z int) {
	for _, cell := range m.cells {
		if cell.z == z {
			fmt.Printf("x: %d y:%d\n", cell.x, cell.y)
		}
	}
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
