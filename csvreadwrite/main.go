package main

import (
	"encoding/csv"
	"fmt"
	"golang.org/x/text/encoding/japanese"
	"io"
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

	// 入力ファイル (Tab区切り/ShiftJIS/CRLF)
	src, err := os.Open("./data/input.sjis.tsv")
	if err != nil {
		panic(fmt.Errorf("input file open failed: %v", err))
	}
	defer func(){
		log.Println("close input file")
		_ = src.Close()
	}()
	r := csv.NewReader(japanese.ShiftJIS.NewDecoder().Reader(src))
	r.Comma = '\t'

	// 出力ファイル (カンマ区切り/UTF-8/LF)
	dst, err := os.Create("./data/output.csv")
	if err != nil {
		panic(fmt.Errorf("output file open failed: %v", err))
	}
	defer func(){
		log.Println("close output file")
		_ = dst.Close()
	}()
	w := csv.NewWriter(dst)
	w.Comma = ','
	w.UseCRLF = false

	// ヘッダ行
	header, err := r.Read()
	if err != nil {
		panic(fmt.Errorf("read header failed: %v", err))
	}
	err = w.Write(header)
	if err != nil {
		panic(fmt.Errorf("write header failed: %v", err))
	}

	// データ行
	for {
		rec, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(fmt.Errorf("read record failed: %v", err))
		}
		err = w.Write(rec)
		if err != nil {
			panic(fmt.Errorf("write record failed: %v", err))
		}
	}
	w.Flush() // バッファ書き出し
}
