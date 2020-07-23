package main

import (
	"fmt"
	"os"

	"github.com/gocarina/gocsv"
	"process-data/elasticservice"
)

// //Vendas struct insert elasticsearch
// type Vendas struct {
// 	Data string `csv:"data"`
// 	Escrv int `csv:"escrv"`
// 	Material int64 `csv:"material"`
// 	GrpMerc int64 `csv:"grp.merc."`
// 	QtdFaturd int `csv:"qnt.faturd"`
// }

func readCSV(filename string)([]*elasticservice.Vendas, error){
	in, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer in.Close()

    vendas := []*elasticservice.Vendas{}

    if err := gocsv.UnmarshalFile(in, &vendas); err != nil {
        panic(err)
    }
	
	return vendas, err
}

func main() {

	// init elasticsearch
	elasticservice.Init()

	// Open the file
	vendas, err := readCSV("./raw_data/VENDAS_20190519.csv")
	if err != nil{
		panic(err)
	}
	
	for _, venda := range vendas {
		fmt.Println(venda.Data)
		//insert elastic
		
		elasticservice.Insertion(venda)
	}

	fmt.Println("Finalize")
}