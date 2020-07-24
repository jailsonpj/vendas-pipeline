package main

import (
	"fmt"
	"os"
	"github.com/gocarina/gocsv"
	"process-data/elasticservice"
	"log"
	"bufio"
	"github.com/szyhf/go-excel"
	"regexp"
	"strings"
	"strconv"
)

func readCSV(filename string){

	in, err := os.Open(filename)
    if err != nil {
        panic(err)
    }
    defer in.Close()

    vendas := []*elasticservice.Vendas{}

    if err := gocsv.UnmarshalFile(in, &vendas); err != nil {
        panic(err)
    }
	
	
	for _, venda := range vendas {
		elasticservice.InsertionCSV(venda)
		fmt.Println("[READCSV][INSERT][DATA] ", venda)
	}

}


func readXLSX(filename string){
	
	conn := excel.NewConnecter()
	err := conn.Open(filename)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	rd, err := conn.NewReader("Sheet1")
	if err != nil {
		panic(err)
	}
	defer rd.Close()

	for rd.Next() {
		// Read a row into a struct.
		var s elasticservice.XlsxData
		var data elasticservice.Data
		err:=rd.Read(&s)
		if err!=nil{
			panic(err)
		}

		if i, err := strconv.Atoi(s.Escrv); err == nil {
			data.Escrv = i
		}

		if i, err := strconv.ParseInt(s.Material, 10, 64); err == nil {
			data.GrpMerc = i
		}

		if i, err := strconv.ParseInt(s.GrpMerc, 10, 64); err == nil {
			data.Material = i
		}
		
		if i, err := strconv.Atoi(s.QtdFaturd); err == nil {
			data.QtdFaturd = i
		}

		data.Data = s.Data

		elasticservice.InsertionXLSX(data)
		fmt.Println("[READXLSX][INSERT][DATA] ", data)
		
	}
	
}

// func readHTML(filename string){
// 	r := strings.NewReader(filename)
// 	str := NewSkipTillReader(r, []byte("<body>"))
// 	rtr := NewReadTillReader(str, []byte("</body>"))
// 	bs, err := ioutil.ReadAll(rtr)
// 	fmt.Println(string(bs), err)
// }

func readTXT(filename string){
	file, err := os.Open(filename)
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string
 
	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}
 
	file.Close()

	m := regexp.MustCompile(`^\|[0-9]{2}\.[0-9]{2}\.[0-9]{4}\|[0-9]{3}\s \|[0-9]{3,6}\s \|[0-9]{1,10}\|\s       [1-9]\s\|$`)
	
	for _, eachline := range txtlines {
		strmatch := m.MatchString(eachline)
		//fmt.Println("Match: ", strmatch)

		if strmatch == true{
			strcompile := m.FindString(eachline)
			strnospace := strings.Replace(strcompile, " ", "", -1)
			strsplit := strings.Split(strnospace,"|")
			
			strdata := strsplit[1:len(strsplit)-1]
			
			var data elasticservice.Data
			data.Data = strdata[0]

			if i, err := strconv.Atoi(strdata[1]); err == nil {
				data.Escrv = i
			}

			if i, err := strconv.ParseInt(strdata[2], 10, 64); err == nil {
				data.GrpMerc = i
			}

			if i, err := strconv.ParseInt(strdata[3], 10, 64); err == nil {
				data.Material = i
			}
			
			if i, err := strconv.Atoi(strdata[4]); err == nil {
				data.QtdFaturd = i
			}

			elasticservice.InsertionTXT(data)
			fmt.Println("[READTXT][INSERT][DATA] ", data)
		}
		
	}

}

func main() {
	elasticservice.Init()
	readCSV("./raw_data/VENDAS_20190519.csv")
	readTXT("./raw_data/VENDAS_20190523.txt")
	readXLSX("./raw_data/VENDAS_20190520_20190522.xlsx")
	//readHTML("./raw_data/VENDAS_20200524_20200525.html")

}