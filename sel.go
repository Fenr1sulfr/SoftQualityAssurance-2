package main

import (
	"fmt"
	"regexp"
	"strings"
	"tgbot/samples/Sel/cfg"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

type TestCase struct {
	driver selenium.WebDriver
	config *cfg.Config
}

func TestCaseBuild() *TestCase {
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
	return &TestCase{
		driver: driver,
		config: cfg.MustLoad(),
	}
}

func (t TestCase) DissassembleQuery() []string {
	re := regexp.MustCompile("[^\\s]+")
	words := re.FindAllString(t.config.Query, -1)
	return words
}

func (t TestCase) BuildQuery() string {
	query := t.DissassembleQuery()
	var builder strings.Builder
	n := len(query)
	builder.Write([]byte(t.config.Host))
	for i, v := range query {
		if i == n-1 {
			builder.Write([]byte(v))
			break
		}
		builder.Write([]byte(v))
		builder.Write([]byte("+"))
	}
	return builder.String()
}

func main() {
	service, err := selenium.NewChromeDriverService("./samples/Sel/chromedriver-win64/chromedriver.exe", 4444)
	if err != nil {
		fmt.Println(err)
	}
	defer service.Stop()
	test := TestCaseBuild()

	query := test.BuildQuery()
	err = test.driver.Get(query)

	if err != nil {
		fmt.Println(err)
	}
	html, err := test.driver.PageSource()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(html)
}
