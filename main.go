package main

import (
	"bufio"
	"context"
	"crawler/crawler"
	"fmt"
	"os"
	"strings"

	"github.com/chromedp/chromedp"
)

func main() {
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// url := "https://cmlabs.co"
	reader := bufio.NewReader(os.Stdin)
	url := ""

	for {
		fmt.Print("> Input url: ")
		text, _ := reader.ReadString('\n')

		url = strings.Replace(text, "\n", "", -1)
		break
	}

	fmt.Printf("> Crawling %s. This might take a while...\n", url)

	crawler := crawler.New(ctx, url)
	if err := crawler.Crawl(); err != nil {
		fmt.Println(err)
	}
}
