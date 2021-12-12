package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

type Point struct {
	x, y       int
	valid      bool
	basin_size int
}

type MapPoint struct {
	x, y   int
	value  int
	marked bool
}

func findBasinSize(mapPoints [][]MapPoint, pt Point) int {

	basinSize := 1
	var to_explore []MapPoint
	if !pt.valid || mapPoints[pt.x][pt.y].value == 9 {
		return 0
	}
	flag := true

	to_explore = append(to_explore, mapPoints[pt.x][pt.y])
	mapPoints[pt.x][pt.y].marked = true
	var to_add []MapPoint
	for flag {
		to_add = []MapPoint{}
		for _, item := range to_explore {
			if item.x == 0 {
				if mapPoints[item.x+1][item.y].value != 9 && !mapPoints[item.x+1][item.y].marked {
					to_add = append(to_add, mapPoints[item.x+1][item.y])
					mapPoints[item.x+1][item.y].marked = true
					basinSize += 1
				}
			} else if item.x < len(mapPoints)-1 {
				if mapPoints[item.x+1][item.y].value != 9 && !mapPoints[item.x+1][item.y].marked {
					to_add = append(to_add, mapPoints[item.x+1][item.y])
					mapPoints[item.x+1][item.y].marked = true
					basinSize += 1
				}
				if mapPoints[item.x-1][item.y].value != 9 && !mapPoints[item.x-1][item.y].marked {
					to_add = append(to_add, mapPoints[item.x-1][item.y])
					mapPoints[item.x-1][item.y].marked = true
					basinSize += 1
				}
			} else if item.x == len(mapPoints)-1 {
				if mapPoints[item.x-1][item.y].value != 9 && !mapPoints[item.x-1][item.y].marked {
					to_add = append(to_add, mapPoints[item.x-1][item.y])
					mapPoints[item.x-1][item.y].marked = true
					basinSize += 1
				}
			}

			if item.y == 0 {
				if mapPoints[item.x][item.y+1].value != 9 && !mapPoints[item.x][item.y+1].marked {
					to_add = append(to_add, mapPoints[item.x][item.y+1])
					mapPoints[item.x][item.y+1].marked = true
					basinSize += 1
				}
			} else if item.y < len(mapPoints[0])-1 {
				if mapPoints[item.x][item.y+1].value != 9 && !mapPoints[item.x][item.y+1].marked {
					to_add = append(to_add, mapPoints[item.x][item.y+1])
					mapPoints[item.x][item.y+1].marked = true
					basinSize += 1
				}
				if mapPoints[item.x][item.y-1].value != 9 && !mapPoints[item.x][item.y-1].marked {
					to_add = append(to_add, mapPoints[item.x][item.y-1])
					mapPoints[item.x][item.y-1].marked = true
					basinSize += 1
				}
			} else if item.y == len(mapPoints[0])-1 {
				if mapPoints[item.x][item.y-1].value != 9 && !mapPoints[item.x][item.y-1].marked {
					to_add = append(to_add, mapPoints[item.x][item.y-1])
					mapPoints[item.x][item.y-1].marked = true
					basinSize += 1
				}
			}
		}

		to_explore = []MapPoint{}
		for _, item := range to_add {
			to_explore = append(to_explore, item)
		}

		if len(to_add) == 0 {
			flag = false
		}
	}
	return basinSize
}

func run(s string) interface{} {
	lines := strings.Split(s, "\n")
	lines = lines[:len(lines)-1]
	var lows []Point

	heightmap := make([][]int, len(lines))
	for i := range heightmap {
		heightmap[i] = make([]int, len(lines[0]))
		for j := range lines[i] {
			val, _ := strconv.Atoi(string(lines[i][j]))
			heightmap[i][j] = val
		}
	}

	for i, row := range heightmap {
		for j := range row {
			lower_than := 0
			if j == 0 {
				lower_than += 1
			} else if row[j] < row[j-1] {
				lower_than += 1
			}

			if j == len(row)-1 {
				lower_than += 1
			} else if row[j] < row[j+1] {
				lower_than += 1
			}

			if i == 0 {
				lower_than += 1
			} else if row[j] < heightmap[i-1][j] {
				lower_than += 1
			}

			if i == len(heightmap)-1 {
				lower_than += 1
			} else if row[j] < heightmap[i+1][j] {
				lower_than += 1
			}
			if lower_than == 4 {
				lows = append(lows, Point{i, j, true, 0})
			}

		}
	}

	mappoints := make([][]MapPoint, len(heightmap))
	for i := range mappoints {
		mappoints[i] = make([]MapPoint, len(heightmap[i]))
		for j := range heightmap[i] {
			mappoints[i][j] = MapPoint{i, j, heightmap[i][j], false}
		}
	}

	for i := range lows {
		lows[i].basin_size = findBasinSize(mappoints, lows[i])
	}

	sort.Slice(lows, func(a, b int) bool { return lows[a].basin_size > lows[b].basin_size })

	return lows[0].basin_size * lows[1].basin_size * lows[2].basin_size
}

func main() {

	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	start := time.Now()
	result := run(string(input))

	fmt.Printf("Duration: %f\n", time.Now().Sub(start).Seconds()*1000)
	fmt.Print("Result: ")
	fmt.Println(result)
}
