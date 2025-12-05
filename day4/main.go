package main

import (
	"fmt"
	"os"
	"strings"
)

var (
	roll  = rune('@')
	empty = rune('.')
)

func main() {
	f, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	answer := 0
	runAgain := true

	rows := strings.Split(string(f), "\n")
	// convert the rows to [][]rune
	grid := make([][]rune, len(rows))
	for i, row := range rows {
		grid[i] = []rune(row)
	}

	type point struct {
		i int
		j int
	}

	toRemove := make([]point, 0)

	for runAgain {
		// set to false, will be set to true if any rolls are removed this iteration
		runAgain = false

		// iterate over the grid
		for i, row := range grid {
			switch i {
			case 0:
				// first row
				for j, char := range row {
					// look at each character
					switch char {
					case roll:
						switch j {
						case 0:
							countNeighbours := 0
							// first char, first row, only neighbours are right and down
							neighbours := []rune{rune(grid[i][j+1]), rune(grid[i+1][j]), rune(grid[i+1][j+1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						case len(row) - 1:
							countNeighbours := 0
							// last char, first row, only neighbours are left and down
							neighbours := []rune{rune(grid[i][j-1]), rune(grid[i+1][j]), rune(grid[i+1][j-1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						default:
							countNeighbours := 0
							// middle char, first row, only neighbours are left, right and down
							neighbours := []rune{rune(grid[i][j-1]), rune(grid[i][j+1]), rune(grid[i+1][j-1]), rune(grid[i+1][j]), rune(grid[i+1][j+1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						}
					case empty:
						// do nothing
					}
				}
			case len(rows) - 1:
				// last row
				for j, char := range row {
					switch char {
					case roll:
						switch j {
						case 0:
							countNeighbours := 0
							// first char, last row, only neighbours are right and up
							neighbours := []rune{rune(grid[i][j+1]), rune(grid[i-1][j]), rune(grid[i-1][j+1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						case len(row) - 1:
							countNeighbours := 0
							// last char, last row, only neighbours are left and up
							neighbours := []rune{rune(grid[i][j-1]), rune(grid[i-1][j]), rune(grid[i-1][j-1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						default:
							countNeighbours := 0
							// middle char, last row, only neighbours are left, right and up
							neighbours := []rune{rune(grid[i][j-1]), rune(grid[i][j+1]), rune(grid[i-1][j-1]), rune(grid[i-1][j]), rune(grid[i-1][j+1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						}
					case empty:
						// do nothing
					}
				}
			default:
				// all other rows
				for j, char := range row {
					switch char {
					case roll:
						switch j {
						case 0:
							countNeighbours := 0
							// first char, middle row, only neighbours are right, up and down
							neighbours := []rune{rune(grid[i][j+1]), rune(grid[i-1][j]), rune(grid[i-1][j+1]), rune(grid[i+1][j]), rune(grid[i+1][j+1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						case len(row) - 1:
							countNeighbours := 0
							// last char, middle row, only neighbours are left, up and down
							neighbours := []rune{rune(grid[i][j-1]), rune(grid[i-1][j]), rune(grid[i-1][j-1]), rune(grid[i+1][j]), rune(grid[i+1][j-1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						default:
							countNeighbours := 0
							// middle char, middle row, all neighbours
							neighbours := []rune{rune(grid[i][j-1]), rune(grid[i][j+1]), rune(grid[i-1][j-1]), rune(grid[i-1][j]), rune(grid[i-1][j+1]), rune(grid[i+1][j-1]), rune(grid[i+1][j]), rune(grid[i+1][j+1])}
							for _, n := range neighbours {
								if n == roll {
									countNeighbours++
								}
							}
							if countNeighbours < 4 {
								toRemove = append(toRemove, point{i, j})
							}
						}
					case empty:
						// do nothing
					}
				}
			}
		}
		fmt.Println("number to be removed", len(toRemove))

		for _, p := range toRemove {
			runAgain = true
			answer++

			grid[p.i][p.j] = empty
		}

		toRemove = make([]point, 0)

	}

	fmt.Println("answer", answer)
}
