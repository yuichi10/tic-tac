package ml

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"git.corp.yahoo.co.jp/ysawada/tic-tac/matrix"
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
	neural.Theta1 = matrix.NewMatrix(0, 0)
	neural.Theta2 = matrix.NewMatrix(0, 0)
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
	for scanner.Scan() {
		text := scanner.Text()
		nums := n.lineToFloatArray(text)
		n.Theta1.AddRow(nums)
	}

	fp2, err := os.Open(n.configFile + "/" + n.theta2File)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	defer fp2.Close()
	scanner2 := bufio.NewScanner(fp)
	for scanner2.Scan() {
		nums := n.lineToFloatArray(scanner2.Text())
		n.Theta2.AddRow(nums)
	}
}

func calcMatrixMulti(m1 []int, m2 []int) {

}

func (n *Neural) CalculateValue(board, put matrix.Matrix) {
	// バイアスの追加
}
