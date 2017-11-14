package tictac

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/yuichi10/matrix"
	"github.com/yuichi10/tic-tac/computer"
)

var board *matrix.Matrix

func checkWinner() {

}

func fullBoard() {

}

func endOfGame() {

}

func initGame() {
	var err error
	board, err = matrix.New(3, 3, nil)
	if err != nil {
		log.Fatal("Failed to create board")
		os.Exit(1)
	}
}

func getPutInfo(putStr string) (x, y int, err error) {
	putInfo := strings.Split(putStr, " ")
	x = -1
	y = -1
	if len(putInfo) != 2 {
		err = errors.New("The put argument is invalid")
		return
	}
	x, err = strconv.Atoi(putInfo[0])
	if err != nil {
		return
	}
	y, err = strconv.Atoi(putInfo[1])
	if err != nil {
		return
	}
	return
}

func putXY(x, y int) error {
	if x > 3 || x < 1 || y > 3 || y < 1 {
		return errors.New("Inut number is out of board")
	}
	if val, err := board.At(y, x); err != nil || val != 0 {
		return errors.New("You can not put here")
	}
	err := board.Set(y, x, 1)
	return err
}

func Start() {
	initGame()
	cp := computer.New()
	stdin := bufio.NewScanner(os.Stdin)
	board.Show()
	for stdin.Scan() {
		text := stdin.Text()
		x, y, err := getPutInfo(text)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		err = putXY(x, y)
		if err != nil {
			continue
		}
		board.Show()
		time.Sleep(2)
		r, c := cp.Consider(board)
		val, err := board.At(r, c)
		if err != nil {
			fmt.Println("computer bug")
			os.Exit(1)
		}
		if val != 0 {
			fmt.Println("computer trainig does not work well...")
			os.Exit(1)
		}
		board.Set(r, c, 2)
		board.Show()
	}
}
