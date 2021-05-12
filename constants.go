package main

import (
	"runtime"
	"strings"
)

var DRIVERPATH string = getDriverPath()

var SPREADSHEETPATH string = getSpreadPath()

const (
	PASSWORD                   = ""
	USER                       = ""
	HORACAMPO                  = "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_txtHoraVerificacao"
	DATACAMPONAME              = "ctl00$ctl00$ctl00$ContentPlaceHolderCorpo$ContentPlaceHolderCorpo$ContentPlaceHolderCorpo$txtDataVerificacao"
	ATENDIDOCAMPOSELECT string = "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ddlResultadoAnaliseExecucao"
	IFRAMEOBS                  = "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ucEditorHTMLObservacaoExecucao_txbComEditorHTML_ifr"
	OBSCAMPO                   = "tinymce"
	PASSWORDCAMPO              = "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_SenhaCertificadoDigital_txbSenhaCertificadoDigital"
	SALVARBUTTON               = "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_btnSalvar"
)

func getDriverPath() string {
	if strings.Contains(runtime.GOOS, "window") {
		return "D:\\chromedriver.exe"
	}
	return "/home/eduardo/automation/chromedriver"

}

func getSpreadPath() string {
	if strings.Contains(runtime.GOOS, "window") {
		return "D:\\Documentos\\Users\\Eduardo\\Documentos\\ANTT\\OneDrive - ANTT- Agencia Nacional de Transportes Terrestres\\sistema\\sifamadocs\\planilha\\verificacao.xlsx"
	}
	return "/home/eduardo/Documentos/projetos/sifamaSources/verificacao.xlsx"
}
