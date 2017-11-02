package main

import (
	"flag"
	"log"
	"os"
	"strconv"
	"strings"

	"git.corp.yahoo.co.jp/ysawada/tic-tac/matrix"
	"git.corp.yahoo.co.jp/ysawada/tic-tac/ml"
)

func getBoardInfo(boardStr string) []float64 {
	var err error
	boardInfo := strings.Split(boardStr, " ")
	board := make([]float64, len(boardInfo))
	for i, val := range boardInfo {
		board[i], err = strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if board[i] != 0 && board[i] != 1 && board[i] != 2 {
			log.Fatal("the board value is strange")
			os.Exit(1)
		}
	}
	if len(board) != 9 {
		log.Fatal("the board Length is strange")
		os.Exit(1)
	}
	return board
}

func getPutInfo(putStr string) int {
	var err error
	putInfo := strings.Split(putStr, " ")
	put := make([]int, len(putInfo))
	for i, val := range putInfo {
		put[i], err = strconv.Atoi(val)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		if put[i] > 10 || put[i] < 0 {
			log.Fatal("the board value is strange")
			os.Exit(1)
		}
	}
	if len(put) != 2 {
		log.Fatal("the put number is strange")
		os.Exit(1)
	}
	result := (put[0]-1)*put[1] + put[1]
	return result
}

func main() {
	board := flag.String("b", "", "board info")
	put := flag.String("p", "", "put info")
	flag.Parse()
	if *board == "" || *put == "" {
		log.Fatal("there is argument error")
		os.Exit(1)
	}
	boardNums := getBoardInfo(*board)
	boardMatrix := matrix.NewMatrix(0, 0)
	boardMatrix.AddRow(boardNums)
	boardMatrix.Show()
	putNums := getPutInfo(*put)
	putMatrix := matrix.NewMatrix(1, 9)
	// 先行の1をいれる
	putMatrix.Set(1, putNums, 1)

	putMatrix.Show()
	neural := ml.NewNeural()
	neural.LoadTheta()
	// neural.Theta1.Show()
	// neural.Theta1.Size()
}
