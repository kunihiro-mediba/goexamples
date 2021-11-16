package main

import (
	"encoding/csv"
	"golang.org/x/text/encoding/japanese"
	"log"
	"os"
)

func main() {
	log.Println("start")
	defer func(){
		err := recover()
		if err != nil {
			log.Fatal(err)
		}
		log.Println("success")
	}()

	f, err := os.Create("./data/input.sjis.tsv")
	if err != nil {
		panic(err)
	}
	defer func(){
		_ = f.Close()
	}()
	w := csv.NewWriter(japanese.ShiftJIS.NewEncoder().Writer(f))
	w.Comma = '\t'

	data := [][]string{
		{"A1","B1","C1"},
		{"あいうえお","カキクケコ","表示\n改行"},
		{"Alpha","Bravo","Charlie"},
	}
	w.WriteAll(data)
	w.Flush()
}
