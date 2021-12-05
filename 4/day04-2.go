package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"time"
)

type Cell struct {
	marked bool
	value  int
}

type Board struct {
	cells *[5][5]Cell
	done  bool
}

func checkRow(row [5]Cell) bool {
	for _, cell := range row {
		if cell.marked == false {
			return false
		}
	}
	return true
}

func (b Board) checkWin() bool {
	for _, row := range b.cells {
		if checkRow(row) == true {
			return true
		}
	}
	var col [5]Cell
	for i, row := range b.cells {

		for j := range row {
			col[j] = b.cells[j][i]
		}

		if checkRow(col) == true {
			return true
		}
	}

	return false
}

func (b Board) calcScore(num int) int {
	total := 0
	for i, row := range b.cells {
		for j := range row {
			if b.cells[i][j].marked == false {
				total += b.cells[i][j].value
			}
		}
	}
	return total * num
}

func (b Board) markBoard(val int) {
	for i, row := range b.cells {
		for j := range row {
			if b.cells[i][j].value == val {
				b.cells[i][j].marked = true
			}
		}
	}
}

func (b Board) printBoard() {
	fmt.Println("[+] Printing Board =====> ", b.done)
	for i, row := range b.cells {
		for j := range row {
			fmt.Print(b.cells[i][j], " ")
		}
		fmt.Println()
	}
	fmt.Println("[+] Board End =====>")
}

func createCells(lines []string) *[5][5]Cell {
	var newCells = [5][5]Cell{}
	var fields []string
	var cellVal int
	for i, raw := range lines {
		fields = strings.Fields(raw)
		for j, num := range fields {
			cellVal, _ = strconv.Atoi(num)
			newCells[i][j] = Cell{marked: false, value: cellVal}
		}
	}

	return &newCells
}

func checkAllBoardsDone(boards []Board) bool {
	for _, board := range boards {
		if board.done == false {
			return false
		}
	}
	return true
}

func run(s string) interface{} {
	lines := strings.Split(s, "\n")

	drawnNums_str := strings.Split(lines[0], ",")
	drawnNums := make([]int, len(drawnNums_str))
	// This seems potentially inefficient
	// I miss my Python list comprehensions :(
	for i, raw := range drawnNums_str {
		val, _ := strconv.Atoi(raw)
		drawnNums[i] = val
	}

	var boards []Board
	var newCells *[5][5]Cell
	var newBoard Board
	for i := 2; i < len(lines); i += 6 {
		newCells = createCells(lines[i : i+5])
		newBoard = Board{cells: newCells, done: false}
		boards = append(boards, newBoard)
	}

	for _, num := range drawnNums {
		for i, b := range boards {
			if boards[i].done == false {
				b.markBoard(num)
				if b.checkWin() == true {
					// if using b here, it's a copied version
					// have to index to actual object
					boards[i].done = true
					if checkAllBoardsDone(boards) == true {
						return b.calcScore(num)
					}
				}
			}
		}
	}

	return 0
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
