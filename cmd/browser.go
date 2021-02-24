/*
Copyright © 2021 ks6088ts <ks6088ts@gmail.com>

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
*/

// Package cmd ...
package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/sclevine/agouti"
	"github.com/spf13/cobra"
)

type browserCmdOption struct {
	url   string
	xpath string
	mode  string
}

var browserCmdOptionVar = &browserCmdOption{}

func getDriver(mode string) *agouti.WebDriver {
	if mode == "chrome" {
		return agouti.ChromeDriver()
	}
	return agouti.ChromeDriver(
		agouti.ChromeOptions("args", []string{
			"--headless",             // headlessモードの指定
			"--window-size=1280,800", // ウィンドウサイズの指定
		}),
	)
}

// browserCmd represents the browser command
var browserCmd = &cobra.Command{
	Use:   "browser",
	Short: "scrape with XPATH",
	Long:  `scrape elements matched with the specified XPATH with browser`,
	Run: func(cmd *cobra.Command, args []string) {
		driver := getDriver(browserCmdOptionVar.mode)
		if err := driver.Start(); err != nil {
			log.Print("Failed to start driver. Please set PATH to chromedriver.")
			os.Exit(1)
		}
		defer driver.Stop()

		page, err := driver.NewPage()
		if err != nil {
			log.Print("Failed to start page")
			driver.Stop()
			os.Exit(1)
		}

		if err := page.Navigate(browserCmdOptionVar.url); err != nil {
			log.Printf("Failed to navigate page %v", browserCmdOptionVar.url)
			driver.Stop()
			os.Exit(1)
		}

		title, err := page.Title()
		if err != nil {
			log.Printf("Failed to get title %v", browserCmdOptionVar.url)
			driver.Stop()
			os.Exit(1)
		}

		fmt.Printf("[%v](%v)\n", title, browserCmdOptionVar.url)
		items := page.AllByXPath(browserCmdOptionVar.xpath)
		itemsCount, err := items.Count()
		if err != nil {
			log.Printf("Failed to get items")
			driver.Stop()
			os.Exit(1)
		}
		for i := 0; i < itemsCount; i++ {
			if text, err := items.At(i).Text(); err == nil {
				fmt.Printf("<item> %v\n", text)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(browserCmd)

	browserCmd.Flags().StringVarP(&browserCmdOptionVar.url, "url", "l", "https://qiita.com/", "URL (required)")
	browserCmd.MarkFlagRequired("url")

	browserCmd.Flags().StringVarP(&browserCmdOptionVar.xpath, "xpath", "x", "//a[@class='tr-Item_title']", "XPath (required)")
	browserCmd.MarkFlagRequired("xpath")

	browserCmd.Flags().StringVarP(&browserCmdOptionVar.mode, "mode", "m", "headless", "Mode (headless|chrome)")
}
