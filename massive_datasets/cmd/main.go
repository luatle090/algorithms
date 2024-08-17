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

	start := time.Now()
	if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
		fmt.Println("start create file")
		massivedatasets.Create(filename, n)
		fmt.Println("end create file: ", time.Until(start))
	}

	fmt.Println("file exists")

	start = time.Now()
	sum, err := massivedatasets.SumFirstOneMillion(filename, m)
	if err != nil {
		fmt.Println("sum first one million: ", err)
		return
	}
	fmt.Println("end sum of the first 1 million: ", -time.Until(start))
	fmt.Println("result ", sum)

	start = time.Now()
	sum, sumRandom, err := massivedatasets.SumRandomlyChosenOneMillion(filename, m)
	if err != nil {
		fmt.Println("sum randomly one million: ", err)
		return
	}

	fmt.Println("end sum of the randomly 1 million: ", -time.Until(start))
	fmt.Println("result ", sum, "result randomly", sumRandom)
}
