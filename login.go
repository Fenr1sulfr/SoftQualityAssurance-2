package main

import (
	"fmt"
	"os"
	"tgbot/samples/Sel/cfg"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type TestLogin struct {
	driver selenium.WebDriver
	tla    *cfg.TestLoginAndOut
}

func TestLoginBuild() *TestLogin {
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

func (t TestLogin) ProcessLoginFunc() bool {
	if t.driver == nil || t.tla == nil {
		fmt.Println("Driver or TestLoginAndOut is not initialized")
		return false
	}

	err := t.driver.Get(t.tla.URL)
	if err != nil {
		fmt.Println("Failed to navigate to URL:", err)
		return false
	}

	loginElement, err := t.driver.FindElement(selenium.ByID, "Email")
	if err != nil {
		fmt.Println("Failed to find username element:", err)
		return false
	}

	err = loginElement.SendKeys(t.tla.Login)
	if err != nil {
		fmt.Println("Failed to send keys to username element:", err)
		return false
	}

	passwordElement, err := t.driver.FindElement(selenium.ByID, "Password")
	if err != nil {
		fmt.Println("Failed to find password element:", err)
		return false
	}

	err = passwordElement.SendKeys(t.tla.Password)
	if err != nil {
		fmt.Println("Failed to send keys to password element:", err)
		return false
	}

	loginButtonElement, err := t.driver.FindElement(selenium.ByCSSSelector, "button[type='submit']")
	if err != nil {
		fmt.Println("Failed to find login button:", err)
		return false
	}

	err = loginButtonElement.Click()
	if err != nil {
		fmt.Println("Failed to click login button:", err)
		return false
	}
	bytes, err := t.driver.Screenshot()

	if err != nil {
		fmt.Println("Failed to make screenshot :", err)
		return false
	}
	if err := os.WriteFile("screenshotIn.png", bytes, 0644); err != nil {
		fmt.Println("Failed to save screenshot :", err)
		return false
	}
	exitButtonElement, err := t.driver.FindElement(selenium.ByXPATH, "/html/body/nav/div[2]/form/button")
	if err != nil {
		fmt.Println("Failed to find exit button :", err)
		return false
	}
	err = exitButtonElement.Click()
	if err != nil {
		fmt.Println("Failed to click exit :", err)
		return false
	}
	bytes, err = t.driver.Screenshot()

	if err != nil {
		fmt.Println("Failed to make screenshot :", err)
		return false
	}
	if err := os.WriteFile("screenshotExit.png", bytes, 0644); err != nil {
		fmt.Println("Failed to save screenshot :", err)
		return false
	}
	return true
}
