package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
)

type csvData []float64

func GetCsvData() csvData {
	// Variable Initialization
	newData := csvData{}

	// Open CSV as file
	f, err := os.Open("USDJPY2014-2016.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	// Read CSV File
	data := csv.NewReader(f)
	c, err := data.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	for i := range c {
		// Break if slice size 6
		if i >= 6 {
			break
		}

		// Convert string to float 64
		convert, err := strconv.ParseFloat(c[i][5], 64)
		if err != nil {
			log.Fatal(err)
		}

		// Insert data to new Slice
		newData = append(newData, convert)
	}

	return newData
}

func (this csvData) PrintCsvData() {
	for _, val := range this {
		fmt.Println(val)
	}
}

func main() {

	// Mendefinisikan Variabel
	data := GetCsvData()
	pop_size := 7
	gen_size := 6
	individu := make([][]float64, pop_size)
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	// Mendefinisikan array 2d dengan slice
	for i := range individu {
		individu[i] = make([]float64, gen_size)
	}

	// Mengisi data array slice dengan nilai random antara 0 - 1 bertipe float
	for i := range individu {
		for j := range individu[i] {
			individu[i][j] = rand.Float64()
		}
		fmt.Println(individu[i])
	}

	data.PrintCsvData()
}
