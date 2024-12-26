// /html/body/div[1]/div[1]/div/div[2]/div/div[2]/div[3]

package main

import (
	"fmt"
	"log"
	"os"
	"tgbot/samples/Sel/cfg"
	"time"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type TestBookFlight struct {
	driver selenium.WebDriver
	tla    *cfg.TestLoginAndOut
}

func TestBoolFlightBuild() *TestLogin {
	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"--headless", // comment out this line for testing
	}})
	driver, err := selenium.NewRemote(caps, "http://localhost:4444/wd/hub")
	if err != nil {
		fmt.Println(err)
	}
	err = driver.MaximizeWindow("")
	if err != nil {
		fmt.Println(err)
	}
	return &TestLogin{
		driver: driver,
		tla:    cfg.MustLoadElements(),
	}
}

func (t TestLogin) ProcessBuyTicketFunc() bool {
	err := t.driver.Get("https://tickets.kz/gd/search/results/forward/astana/kustanai/31.12.2024")
	if err != nil {
		log.Fatalf("Ошибка загрузки страницы: %v", err)
	}
	screenshot, err := t.driver.Screenshot()
	if err = os.WriteFile("ticket1.png", screenshot, 0677); err != nil {
		log.Fatalf("Ошибка получения текста: %v", err)
		return false
	}
	firstDivXPath := "/html/body/div[1]/div[1]/div/div[2]/div/div[2]/div[3]"
	firstDiv, err := t.driver.FindElement(selenium.ByXPATH, firstDivXPath)
	if err != nil {
		log.Fatalf("Ошибка поиска первого дива: %v", err)
	}
	nextDiv, err := firstDiv.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div[2]/div/div[2]/div[3]/div[1]")
	if err != nil {
		log.Fatalf("Ошибка поиска следующего дива: %v", err)
	}
	screenshot, err = t.driver.Screenshot()
	if err = os.WriteFile("ticket2.png", screenshot, 0677); err != nil {
		log.Fatalf("Ошибка получения текста: %v", err)
		return false
	}
	button, err := nextDiv.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[1]/div/div[2]/div/div[2]/div[3]/div[1]/div[2]/div/div/div[3]/div/div[1]/div/div[3]/div/button")
	button.Click()
	time.Sleep(5 * time.Second) //div.grid-cell:nth-child(30) > button:nth-child(1) > div:nth-child(1)  div.grid-cell:nth-child(13) > button:nth-child(1)
	seatElement, err := t.driver.FindElement(selenium.ByCSSSelector, "div.grid-cell:nth-child(35) > button:nth-child(1)")
	if err != nil {
		log.Fatalf("Ошибка получения сидения: %v", err)
	}
	screenshot, err = t.driver.Screenshot()
	if err = os.WriteFile("ticket3.png", screenshot, 0677); err != nil {
		log.Fatalf("Ошибка получения текста: %v", err)
		return false
	}
	seatElement.Click()
	time.Sleep(5 * time.Second)
	continueDiv, err := t.driver.FindElement(selenium.ByXPATH, "/html/body/div[1]/div[2]/div[2]/div/div[1]/div[5]")
	if err != nil {
		log.Fatalf("Ошибка получения дива `продолжить`: %v", err)
	}
	continueButton, err := continueDiv.FindElement(selenium.ByCSSSelector, "button.ml-auto")
	if err != nil {
		log.Fatalf("Ошибка получения кнопки `продолжить`: %v", err)
	}
	err = continueButton.Click()
	if err != nil {
		log.Fatalf("Ошибка получения текста: %v", err)
	}
	time.Sleep(3 * time.Second)
	screenshot, err = t.driver.Screenshot()
	if err = os.WriteFile("ticket4.png", screenshot, 0677); err != nil {
		log.Fatalf("Ошибка получения текста: %v", err)
		return false
	}

	return true
}
