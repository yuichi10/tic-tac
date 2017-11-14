package computer

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/yuichi10/matrix"
)

const (
	inputSize  = 9
	hiddenSize = 1
	outputSize = 9
)

var (
	hiddenLayers     = []int{15}
	biasHiddenLayers = []int{16}
)

type Neural struct {
	configFile string
	theta1File string
	theta2File string
	Theta1     *matrix.Matrix
	Theta2     *matrix.Matrix
}

func getConfigPath() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s/config", dir), nil
}

func NewNeural() *Neural {
	var err error
	neural := new(Neural)
	neural.configFile, err = getConfigPath()
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	neural.theta1File = "theta1.txt"
	neural.theta2File = "theta2.txt"
	neural.Theta1, err = matrix.New(1, 10, nil)
	if err != nil {
		log.Fatal("failed to create matrix")
		os.Exit(1)
	}
	neural.Theta2, err = matrix.New(1, 16, nil)
	if err != nil {
		log.Fatal("failed to create matrix")
		os.Exit(1)
	}
	return neural
}

func (n *Neural) lineToFloatArray(line string) []float64 {
	var err error
	numStr := strings.Split(line, " ")
	res := make([]float64, len(numStr))
	for i, val := range numStr {
		res[i], err = strconv.ParseFloat(val, 64)
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
	}
	return res
}

func (n *Neural) LoadTheta() {
	fp, err := os.Open(n.configFile + "/" + n.theta1File)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fp.Close()
	scanner := bufio.NewScanner(fp)
	firstTime := true
	for scanner.Scan() {
		text := scanner.Text()
		nums := n.lineToFloatArray(text)
		if firstTime {
			n.Theta1, err = matrix.NewVector(nums)
			n.Theta1 = n.Theta1.Transpose()
			if err != nil {
				log.Fatal("failed to create matrix")
			}
			firstTime = false
		} else {
			n.Theta1, err = n.Theta1.AddRow(nums)
		}
	}

	fp2, err := os.Open(n.configFile + "/" + n.theta2File)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fp2.Close()
	scanner2 := bufio.NewScanner(fp2)
	firstTime = true
	for scanner2.Scan() {
		nums := n.lineToFloatArray(scanner2.Text())
		if firstTime {
			n.Theta2, err = matrix.NewVector(nums)
			n.Theta2 = n.Theta2.Transpose()
			if err != nil {
				log.Fatal("failed to create matrix")
			}
			firstTime = false
		} else {
			n.Theta2, err = n.Theta2.AddRow(nums)
		}
	}
}

func (n *Neural) forwardPropagation(board *matrix.Matrix) *matrix.Matrix {
	a1 := board.Vector()
	a1, err := a1.AddRowHEAD(1)
	if err != nil {
		log.Fatal("failed to add baias: %v", err)
		os.Exit(1)
	}
	z2 := n.Theta1.Multi(a1)
	a2 := z2.Sigmoid()
	a2, err = a2.AddRowHEAD(1)
	z3 := n.Theta2.Multi(a2)
	a3 := z3.Sigmoid()
	if a3.CalcErr() != nil {
		fmt.Println(a3.CalcErr())
		os.Exit(1)
	}
	return a3
}
