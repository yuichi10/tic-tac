package computer

import (
	"fmt"
	"os"

	"github.com/yuichi10/matrix"
)

var neural *Neural

type Computer struct {
	neural *Neural
}

func New() *Computer {
	computer := new(Computer)
	computer.initComputer()
	return computer
}

func (c *Computer) initComputer() {
	neural := NewNeural()
	neural.LoadTheta()
	c.neural = neural
}

func (c *Computer) Consider(borad *matrix.Matrix) (int, int) {
	result := c.neural.forwardPropagation(borad)
	resBoard, err := result.Reshape(3, 3)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	resBoard.Show()
	max := -1.0
	r, col := 0, 0
	for i := 1; i <= resBoard.Row(); i++ {
		for j := 1; j <= resBoard.Column(); j++ {
			val, _ := resBoard.At(i, j)
			if val >= max {
				max = val
				r, col = i, j
			}
		}
	}
	return r, col
}
