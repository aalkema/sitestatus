package sitestatus

import (
	"fmt"
	"github.com/wirepair/autogcd"
	"io/ioutil"
	"log"
	"time"
)

type ChromeProperties struct {
	ChromePath string
	TempDir    string
	ChromePort string
}

var startupFlags = []string{"--disable-new-tab-first-run", "--no-first-run", "--disable-translate", "--headless"}
var waitForTimeout = time.Second * 5
var waitForRate = time.Millisecond * 25

var navigationTimeout = time.Second * 10

var stableAfter = time.Millisecond * 450
var stabilityTimeout = time.Second * 2

func CheckSite(c ChromeProperties, host string, element string) (found bool, output string) {
	settings := autogcd.NewSettings(c.ChromePath, tempDir(c))
	settings.RemoveUserDir(true)
	settings.AddStartupFlags(startupFlags)

	auto := autogcd.NewAutoGcd(settings)
	auto.Start()
	defer auto.Shutdown()

	tab, err := auto.GetTab()
	if err != nil {
		return false, "Failed to get tab"
	}

	if _, _, err := tab.Navigate(host); err != nil {
		return false, err.Error()
	}

	err = tab.WaitFor(waitForRate, waitForTimeout, autogcd.ElementByIdReady(tab, element))
	if err != nil {
		return false, "Failed waiting for element to load on page"
	}

	_, _, err = tab.GetElementById(element)
	if err != nil {
		return false, "Loaded page and couldn't find element"
	}

	return true, fmt.Sprintf("Loaded %s and found %s", host, element)
}

func configureTab(tab *autogcd.Tab) {
	tab.SetNavigationTimeout(navigationTimeout) // give up after 10 seconds for navigating, default is 30 seconds
	tab.SetStabilityTime(stableAfter)
}

func tempDir(c ChromeProperties) string {
	dir, err := ioutil.TempDir(c.TempDir, "autogcd")
	if err != nil {
		log.Fatalf("error getting temp dir: %s\n", err)
	}
	return dir
}
