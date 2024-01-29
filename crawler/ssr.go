package crawler

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/chromedp/chromedp"
)

type ssrCrawler struct {
	ctx context.Context
	url string
}

func (c *ssrCrawler) Crawl() error {
	err := chromedp.Run(c.ctx,
		chromedp.Navigate(c.url),
		chromedp.WaitVisible("body"),
	)

	if err != nil {
		fmt.Println("error navigating page:", err)
		return err
	}

	var content string
	err = chromedp.Run(c.ctx,
		chromedp.InnerHTML("html", &content, chromedp.ByQuery),
	)

	if err != nil {
		fmt.Println("error fetching html body:", err)
		return err
	}

	now := time.Now().Unix()
	filename := fmt.Sprintf("./out/ssr-%d.html", now)

	f, err := os.Create(filename)
	if err != nil {
		fmt.Println("error creating file:", err)
		return err
	}

	defer f.Close()

	_, err = f.WriteString(content)
	if err != nil {
		fmt.Println("error writing file:", err)
		return err
	}

	fmt.Printf("> Crawling finished. Result written on %s\n", filename)
	return nil
}
