![test](https://github.com/ks6088ts/scraper-go/workflows/test/badge.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/ks6088ts/scraper-go)](https://goreportcard.com/report/github.com/ks6088ts/scraper-go)
[![GoDoc](https://godoc.org/github.com/ks6088ts/scraper-go?status.svg)](https://godoc.org/github.com/ks6088ts/scraper-go)

# scraper-go
Template repository for Go

# commands

## browser

scrape elements matched with the specified XPATH with browser

**Prerequisites**
- install Google Chrome
- set PATH to [chromedriver](https://chromedriver.chromium.org/downloads)

```bash
# help for browser
scraper-go browser --help

# scrape webpage
scraper-go browser \
    --mode chrome
    --xpath //a/article/h1 \
    --url https://search.yahoo.co.jp/realtime
```