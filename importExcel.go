package main

import (
	"fmt"

	"github.com/360EntSecGroup-Skylar/excelize"
)

func getInfoFromExcel() [][]string {
	f, err := excelize.OpenFile("D:\\sifamadocs\\planilha\\verificacao.xlsx")
	if err != nil {
		fmt.Println(err)
		return nil
	}

	// Get all the rows in the Sheet1.
	rows := f.GetRows("Planilha1")

	return rows
}
