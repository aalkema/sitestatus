package main

import (
	"flag"
	"fmt"
	"github.com/aalkema/sitestatus"
	"os"
)

var (
	chromePath string
	tempDir    string
	chromePort string
	debug      bool
	host       string
	element    string
)

func init() {
	flag.StringVar(&chromePath, "chrome", "C:\\Program Files (x86)\\Google\\Chrome\\Application\\chrome.exe", "path to chrome")
	flag.StringVar(&tempDir, "dir", "C:\\temp\\", "temp directory")
	flag.StringVar(&chromePort, "port", "9222", "Debugger port")
	flag.StringVar(&host, "host", "", "The website to check for element")
	flag.StringVar(&element, "element", "", "The element to look for on host")
}

func main() {
	flag.Parse()

	cp := sitestatus.ChromeProperties{
		chromePath, tempDir, chromePort,
	}

	found, output := sitestatus.CheckSite(cp, host, element)

	if !found {
		fmt.Print(output)
		os.Exit(2)
	} else {
		fmt.Print(output)
		os.Exit(0)
	}
}
