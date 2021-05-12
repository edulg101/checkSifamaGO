package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/go-vgo/robotgo"
)

func reg(id string) string {

	r, _ := regexp.Compile(`TRO0*(\d+)2021`)

	subString := r.FindStringSubmatch(id)

	return subString[1]

}

func errorHandling(e error) {
	if e != nil {
		fmt.Println(e)
		fmt.Scanf("%s")
		os.Exit(1)
	}
}

func checkForMissingTros(tros [][]string, listaTrosEmOrdem []int) {
	flag := false
	var troFaltantes []string
	for i, listaTroExcel := range tros {
		if i < 1 {
			continue
		}
		for _, listaDisponivelNoSistema := range listaTrosEmOrdem {
			v := strconv.Itoa(listaDisponivelNoSistema)
			if strings.Contains(v, listaTroExcel[0]) {

				flag = true
				break
			} else {
				flag = false
			}
		}
		if !flag {
			troFaltantes = append(troFaltantes, listaTroExcel[0])
		}
	}

	if len(troFaltantes) > 0 {
		fmt.Printf("faltando os Seguintes Tros - (%d) : \n", len(troFaltantes))
		for _, v := range troFaltantes {
			fmt.Println(v)
		}
		os.Exit(1)
	}
}

func keepMouseMoving() {
	for {
		robotgo.MoveMouse(100, 300)
		time.Sleep(time.Second)
		robotgo.MoveMouse(300, 500)
		time.Sleep(time.Minute * 2)
	}
}
