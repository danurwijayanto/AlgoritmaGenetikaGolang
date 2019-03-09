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
type initialCromosom [][]float64

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

func GenerateCromosom() initialCromosom {
	var row []float64
	pop_size := 7
	gen_size := 6
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))

	individu := initialCromosom{}

	for i := 0; i < pop_size; i++ {
		row = nil
		for j := 0; j < gen_size; j++ {
			row = append(row, rand.Float64())
		}
		individu = append(individu, row)
	}

	return individu
}

func (this initialCromosom) PrintCsvData() {
	for _, val := range this {
		fmt.Println(val[0])
	}
}
func main() {

	// Mendefinisikan Variabel
	data := GetCsvData()
	data.PrintCsvData()
	cromosom := GenerateCromosom()
	cromosom.PrintCsvData()
}
