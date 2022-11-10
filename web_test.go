package main

import (
	// "encoding/json"
	"fmt"
	// "io/ioutil"
	// "net/http"
	// "os"
	// "strings"
	"testing"

	"github.com/tebeka/selenium"
)

func TestSingle(test *testing.T) {
	test.Parallel()
	// asserter := assert.New(test)
	var buildName = "Demo-GoLang"
	caps := selenium.Capabilities{
		"bstack:options": map[string]interface{}{
			"osVersion":     "12.0",
			"deviceName":    "Samsung Galaxy S22",
			"local":         "false",
			"projectName":   "BrowserStack GoLang",
			"buildName":     buildName,
			"sessionName":   "GoLang Firefox Test Single",
			"debug":         "true",
			"networkLogs":   "true",
			"consoleLogs":   "verbose",
			"appiumVersion": "1.17.0",
		},
		// "browserName":    "Firefox",
		// "browserVersion": "latest", // enable if you want to test using some desktop browser
	}
	// wd, err := selenium.NewRemote(caps, fmt.Sprintf("https://%s:%s@hub-cloud.browserstack.com/wd/hub", os.Getenv("BROWSERSTACK_USERNAME"), os.Getenv("BROWSERSTACK_ACCESSKEY")))
	wd, err := selenium.NewRemote(caps, "https://<BrowserStack_username>_<Browserstack_access_key>@hub-cloud.browserstack.com/wd/hub")
	if err != nil {
		fmt.Print("error")
		panic(err)
	}
	test.Cleanup(func() { wd.Quit() })
	wd.Get("https://www.node-b.stage.webshop.globus.ch/?login")
	// wd.FindElement(selenium.ByName, "//div[@data-testid='general-page-error-content']")
	wd.FindElement(selenium.ByXPATH, "//*[@data-testid='login-loginform-id-input']//input")
	// title, titleErr := wd.Title()
	// if titleErr != nil {
	// 	test.Fatal(titleErr)
	// }
}

// func TestParallel(test *testing.T) {
// 	test.Parallel()
// 	var capabilities []map[string]interface{}
// 	fileData, _ := ioutil.ReadFile("./config/browsers.json")
// 	json.Unmarshal(fileData, &capabilities)
// 	var remoteServerURL = fmt.Sprintf("https://%s:%s@hub-cloud.browserstack.com/wd/hub", os.Getenv("BROWSERSTACK_USERNAME"), os.Getenv("BROWSERSTACK_ACCESSKEY"))
// 	for _, capability := range capabilities {
// 		test.Run(fmt.Sprintf("Running on %s", capability["browserName"]), func(nestedTest *testing.T) {
// 			// nestedTest.Parallel() // when enabled this it runs all tests in parallel but always run for the last capability
// 			if os.Getenv("JENKINS_ENV") != "" {
// 				var options map[string]interface{}
// 				tempOptions, _ := capability["bstack:options"]
// 				options = tempOptions.(map[string]interface{})
// 				options["buildName"] = os.Getenv("BROWSERSTACK_BUILD_NAME")
// 				capability["bstack:options"] = options
// 			}
// 			wd, err := selenium.NewRemote(capability, remoteServerURL)
// 			if err != nil {
// 				panic(err)
// 			}
// 			nestedTest.Cleanup(func() {
// 				sessionID := wd.SessionID()
// 				wd.Quit()
// 				nestedTest.Log(sessionID)
// 				var req *http.Request
// 				if nestedTest.Failed() {
// 					req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("https://api.browserstack.com/automate/sessions/%s.json", sessionID), strings.NewReader(`{"status":"failed", "reason":"failed all tests"}`))
// 					if err != nil {
// 						nestedTest.Fatal(err)
// 					}
// 				} else {
// 					req, err = http.NewRequest(http.MethodPut, fmt.Sprintf("https://api.browserstack.com/automate/sessions/%s.json", sessionID), strings.NewReader(`{"status":"passed", "reason":"passed all tests"}`))
// 					if err != nil {
// 						nestedTest.Fatal(err)
// 					}
// 				}
// 				req.SetBasicAuth(os.Getenv("BROWSERSTACK_USERNAME"), os.Getenv("BROWSERSTACK_ACCESSKEY"))
// 				req.Header.Set("Content-Type", "application/json")
// 				client := &http.Client{}
// 				_, err := client.Do(req)
// 				if err != nil {
// 					nestedTest.Fatal(err)
// 				}
// 			})
// 			nestedTest.Parallel() // adding here to run tests in parallel,
// 			asserter := assert.New(nestedTest)
// 			wd.Get("https://google.com")
// 			title, titleErr := wd.Title()
// 			if titleErr != nil {
// 				nestedTest.Fatal(titleErr)
// 			}
// 			nestedTest.Logf("Title Recieved: %s", title)
// 			asserter.Contains(title, "Google", "Title should contain Google")
// 		})
// 	}
// }

// func TestFail(test *testing.T) {
// 	test.Parallel()
// 	asserter := assert.New(test)
// 	var buildName = "Demo-GoLang"
// 	if os.Getenv("JENKINS_ENV") != "" {
// 		buildName = os.Getenv("BROWSERSTACK_BUILD_NAME")
// 	}
// 	caps := selenium.Capabilities{
// 		"bstack:options": map[string]interface{}{
// 			"os":              "Windows",
// 			"osVersion":       "10",
// 			"local":           "false",
// 			"seleniumVersion": "4.0.0-alpha-6",
// 			"projectName":     "BrowserStack GoLang",
// 			"buildName":       buildName,
// 			"sessionName":     "GoLang Firefox Test Fail",
// 			"debug":           "true",
// 			"networkLogs":     "true",
// 			"consoleLogs":     "verbose",
// 		},
// 		"browserName":    "Firefox",
// 		"browserVersion": "latest",
// 	}
// 	// wd, err := selenium.NewRemote(caps, fmt.Sprintf("https://%s:%s@hub-cloud.browserstack.com/wd/hub", os.Getenv("BROWSERSTACK_USERNAME"), os.Getenv("BROWSERSTACK_ACCESSKEY")))
// 	wd, err := selenium.NewRemote(caps, "https://adityagholkar_HwrVx3:7ctaSnUrzPkHMDjDd7ox@hub-cloud.browserstack.com/wd/hub")
// 	if err != nil {
// 		// panic(err)
// 	}
// 	test.Cleanup(func() {
// 		if test.Failed() {
// 			wd.ExecuteScript("browserstack_executor: {\"action\": \"setSessionStatus\", \"arguments\": {\"status\":\"failed\"}}", nil)
// 		} else {
// 			wd.ExecuteScript("browserstack_executor: {\"action\": \"setSessionStatus\", \"arguments\": {\"status\":\"passed\"}}", nil)
// 		}
// 		wd.Quit()
// 	})
// 	wd.Get("https://google.com")
// 	title, titleErr := wd.Title()
// 	if titleErr != nil {
// 		test.Fatal(titleErr)
// 	}
// 	asserter.Equal("Microsoft", title, "Title should have been Google")
// }
