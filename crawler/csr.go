package crawler

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type csrCrawler struct {
	ctx context.Context
	url string
}

func (c *csrCrawler) Crawl() error {
	var content string

	err := chromedp.Run(c.ctx,
		chromedp.Navigate(c.url),
		chromedp.Evaluate(`document.documentElement.outerHTML;`, &content),
	)

	if err != nil {
		fmt.Println("Error navigating to url:", err)
		return err
	}

	now := time.Now().Unix()
	filename := fmt.Sprintf("./out/csr-%d.html", now)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("Error writing file:", err)
		return err
	}

	fmt.Printf("> Crawling finished. Result written on %s\n", filename)
	return nil
}
