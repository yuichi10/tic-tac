package main

import (
	"github.com/yuichi10/tic-tac/computer"
)

func main() {
	// put info looks linke "1 2"
	// put := flag.String("p", "", "put info")
	// flag.Parse()

	// putMatrix.Show()
	neural := computer.NewNeural()
	neural.LoadTheta()
	neural.Theta1.Show()
	neural.Theta2.Show()
	// neural.Theta1.Size()
}
