package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

const (
	//Set constants separately chromedriver.exe Address and local call port of
	seleniumPath = DRIVERPATH
	port         = 9515
)

func main() {

	go keepMouseMoving()

	//1. Enable selenium service
	//Set the option of the selium service to null. Set as needed.
	ops := []selenium.ServiceOption{}
	_, err := selenium.NewChromeDriverService(seleniumPath, port, ops...)
	if err != nil {
		fmt.Printf("Error starting the ChromeDriver server: %v", err)
	}
	//Delay service shutdown
	// defer service.Stop()

	//2. Call browser
	//Set browser compatibility. We set the browser name to chrome
	caps := selenium.Capabilities{
		"browserName": "chrome",
	}
	//Call browser urlPrefix: test reference: defaulturlprefix =“ http://127.0.0.1 :4444/wd/hub"
	driver, err := selenium.NewRemote(caps, "http://127.0.0.1:9515/wd/hub")
	if err != nil {
		panic(err)
	}

	defer driver.Quit()

	if err := driver.Get("https://appweb1.antt.gov.br/fisn/Site/TRO/Cadastrar.aspx"); err != nil {
		panic(err)
	}
	fmt.Println("abrindo pagina do Sifama")

	usuario, err := driver.FindElement(selenium.ByID, "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_TextBoxUsuario")
	errorHandling(err)
	senha, err := driver.FindElement(selenium.ByID, "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_TextBoxSenha")
	errorHandling(err)
	entrar, err := driver.FindElement(selenium.ByID, "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ButtonOk")
	errorHandling(err)

	usuario.SendKeys(USER)
	senha.SendKeys(PASSWORD)
	entrar.Click()

	waitForJsAndJquery(driver)

	tros := getInfoFromExcel()

	waitForJsAndJquery(driver)

	i := 1

	mainWindow, _ := driver.CurrentWindowHandle()

	for ; i < len(tros); i++ {
		tro := tros[i]
		driver.SwitchWindow(mainWindow)
		inicioVerificacao(driver, tro, tros, i)
	}

	driver.SwitchWindow(mainWindow)
	fmt.Println("Done")

	driver.ExecuteScript("alert('Terminou')", nil)

	time.Sleep(time.Second * 180)

}

func inicioVerificacao(driver selenium.WebDriver, tro []string, tros [][]string, count int) {
	mainWindow, _ := driver.CurrentWindowHandle()
	troNumber := tro[0]
	troHora := tro[1]
	troText := tro[2]
	var codAtendimento string = ""

	if strings.ToLower(tro[3]) == "s" {
		codAtendimento = "2"
	}
	if strings.ToLower(tro[3]) == "n" {
		codAtendimento = "3"
	}

	waitForJsAndJquery(driver)

	we, err := waitForElementByXpath(driver, "/html/body/div[1]/div[1]/div[1]/div[1]")
	errorHandling(err)
	we.Click()
	fmt.Println("abrindo a lista de TROs ........")

	time.Sleep(time.Second * 2)

	waitForJsAndJquery(driver)

	listaTros, err := driver.FindElements(selenium.ByXPATH, "//div[@class='wingsDivNomeTarefa']")
	errorHandling(err)

	var listaTrosEmOrdem []int
	for _, we := range listaTros {
		text, e := we.Text()
		errorHandling(e)
		troStr := reg(text)
		troInt, _ := strconv.Atoi(troStr)
		listaTrosEmOrdem = append(listaTrosEmOrdem, troInt)
	}

	sort.Ints(listaTrosEmOrdem)

	fmt.Println("Tros disponiveis para análise:")

	for _, v := range listaTrosEmOrdem {
		fmt.Println(v)
	}

	if count == 1 {
		checkForMissingTros(tros, listaTrosEmOrdem)
	}

	for _, x := range listaTros {
		text, _ := x.Text()
		if strings.Contains(text, troNumber+"2021") {
			getTBody, e := x.FindElement(selenium.ByXPATH, "../../..")
			errorHandling(e)
			divToClick, _ := getTBody.FindElement(selenium.ByXPATH, "./tr[5]/td/div[2]")
			err = divToClick.Click()
			errorHandling(err)
			bytes := []byte(text)
			bytes = bytes[25:]
			fmt.Println("Entrando no TRO n. " + string(bytes))
		}
	}

	time.Sleep(time.Second * 10)

	waitForJsAndJquery(driver)

	// wait for tab to open and change focus to it

	flag := true
	for i := 0; i < 120 && flag; i++ {
		handles, err := driver.WindowHandles()
		errorHandling(err)
		for _, wh := range handles {
			if wh != mainWindow && flag {
				err = driver.SwitchWindow(wh)
				errorHandling(err)
				we, err := driver.FindElement("id", "ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_ContentPlaceHolderCorpo_div1")
				if we != nil && err == nil {
					flag = false
					break
				}
				fmt.Println(err)

			}
			time.Sleep(time.Second / 2)
		}
		time.Sleep(time.Second / 2)
		fmt.Println(i)
	}

	todayDate := time.Now()
	today := todayDate.Format("02/01/2006")

	waitForJsAndJquery(driver)

	dataCampo, err := waitForElementByXpath(driver, `//input[@name="ctl00$ctl00$ctl00$ContentPlaceHolderCorpo$ContentPlaceHolderCorpo$ContentPlaceHolderCorpo$txtDataVerificacao"]`)
	errorHandling(err)

	time.Sleep(time.Second * 3)
	fmt.Println("Insere data")
	dataCampo.Clear()

	time.Sleep(time.Second * 3)
	dataCampo.Click()

	time.Sleep(time.Second * 3)
	dataCampo.SendKeys(today)

	passwordElement, err := driver.FindElement("id", PASSWORDCAMPO)
	errorHandling(err)
	passwordElement.Click()

	waitForJsAndJquery(driver)

	_, err = waitForElementById(driver, HORACAMPO, time.Second*30)

	errorHandling(err)

	fmt.Println("insere hora")

	we, err = driver.FindElement("id", HORACAMPO)
	errorHandling(err)
	we.Clear()
	we.SendKeys(troHora)

	we.Clear()
	we.SendKeys(troHora)

	scriptToFillField(driver, HORACAMPO, troHora)

	waitForJsAndJquery(driver)

	fmt.Println("marca como atendido")

	jqueryScript(driver, ATENDIDOCAMPOSELECT, codAtendimento)

	time.Sleep(time.Second)

	scriptToClick(driver, PASSWORDCAMPO)

	jqueryScript(driver, ATENDIDOCAMPOSELECT, codAtendimento)

	waitForJsAndJquery(driver)

	fmt.Println("Insere senha")

	for i := 0; i < 3; i++ {
		we, err = waitForElementById(driver, PASSWORDCAMPO, time.Second*30)
		errorHandling(err)
		we.Clear()
		we.SendKeys(PASSWORD)
		time.Sleep(7000)

	}

	waitForElementById(driver, IFRAMEOBS, time.Second*30)

	driver.SwitchFrame(IFRAMEOBS)

	fmt.Println("insere texto na Observação")

	waitForElementById(driver, OBSCAMPO, time.Second*30)

	we, err = driver.FindElement("id", OBSCAMPO)
	errorHandling(err)

	we.SendKeys(troText)

	driver.SwitchFrame(nil)
	time.Sleep(time.Second / 2)
	waitForJsAndJquery(driver)

	fmt.Println("Envia formulario")

	scriptToClick(driver, SALVARBUTTON)

	time.Sleep(time.Second)

	checkForErrors(driver)

	scriptToClick(driver, "MessageBox_ButtonOk")

	time.Sleep(time.Second)

	fmt.Printf("Salva Tro n. %v/2021 \n", troNumber)

}
