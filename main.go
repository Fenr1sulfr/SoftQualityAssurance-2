package main

import (
	"flag"
	"fmt"

	"github.com/tebeka/selenium"
)

func main() {
	flag.Parse()
	service, err := selenium.NewChromeDriverService("./chromedriver-win64/chromedriver.exe", 4444)
	if err != nil {
		fmt.Println(err)
	}
	defer service.Stop()

	testLogin := TestLoginBuild()
	testLogin.ProcessLoginFunc()
	// testCase := TestCaseBuild()
	// testCase.ProccessGetRawUrl()
}
