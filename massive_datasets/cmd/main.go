package main

import (
	"errors"
	"fmt"
	"os"
	"time"

	massivedatasets "github.com/algorithms/massive_datasets"
)

func main() {

	filename := "one-billion.bin"
	n := 1_000_000_000
	m := 1_000_000

	var file *os.File

	start := time.Now()
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Println("start create file")
		file, err = os.Create(filename)
		if err != nil {
			fmt.Println("error create file")
			return
		}
		massivedatasets.WriteIntTo(file, n)
		fmt.Println("end create file: ", time.Until(start))
	} else {
		file, err = os.Open(filename)
		if err != nil {
			fmt.Println("error open file")
			return
		}
	}
	defer file.Close()

	fmt.Println("file exists")

	start = time.Now()
	sum, err := massivedatasets.SumFirstOneMillion(file, m)
	if err != nil {
		fmt.Println("sum first one million: ", err)
		return
	}
	fmt.Println("end sum of the first 1 million: ", -time.Until(start))
	sumExpected := massivedatasets.SumOneMillion(m)
	fmt.Println("result ", sum, "result expected", sumExpected)

	start = time.Now()
	sum, sumRandom, err := massivedatasets.SumRandomlyChosenOneMillion(file, m)
	if err != nil {
		fmt.Println("sum randomly one million: ", err)
		return
	}

	fmt.Println("end sum of the randomly 1 million: ", -time.Until(start))
	fmt.Println("result ", sum, "result randomly", sumRandom)
	// start = time.Now()
	// sum, sumRandom, err = massivedatasets.SumRandomlyChosenOneMillion2(filename, m)
	// if err != nil {
	// 	fmt.Println("sum randomly one million: ", err)
	// 	return
	// }

	// fmt.Println("end sum of the randomly 1 million: ", -time.Until(start))
	// fmt.Println("result ", sum, "result randomly", sumRandom)
}
